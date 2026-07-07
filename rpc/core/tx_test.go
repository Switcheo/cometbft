package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/cometbft/cometbft-db"

	abci "github.com/cometbft/cometbft/abci/types"
	cfg "github.com/cometbft/cometbft/config"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
	"github.com/cometbft/cometbft/state/txindex/kv"
	"github.com/cometbft/cometbft/types"
)

// TestTxSearchPagination exercises the full TxSearch RPC path (pagination and
// ordering now happen inside the indexer) against a real kv indexer. prove is
// false so the BlockStore is never accessed.
func TestTxSearchPagination(t *testing.T) {
	indexer := kv.NewTxIndex(dbm.NewMemDB())

	// 12 txs sharing a common indexed event, across 3 blocks of 4 txs each, so
	// both the height and the intra-block index ordering are exercised.
	type ref struct {
		height int64
		index  uint32
	}
	var order []ref
	for h := int64(1); h <= 3; h++ {
		for i := uint32(0); i < 4; i++ {
			txr := &abci.TxResult{
				Height: h,
				Index:  i,
				Tx:     types.Tx(fmt.Sprintf("tx-%d-%d", h, i)),
				Result: abci.ExecTxResult{
					Code: abci.CodeTypeOK,
					Events: []abci.Event{{
						Type:       "app",
						Attributes: []abci.EventAttribute{{Key: "group", Value: "g1", Index: true}},
					}},
				},
			}
			require.NoError(t, indexer.Index(txr))
			order = append(order, ref{h, i})
		}
	}
	require.Len(t, order, 12)

	env := &Environment{
		TxIndexer: indexer,
		Config:    *cfg.TestRPCConfig(),
	}
	rctx := &rpctypes.Context{}
	q := `app.group = 'g1'`

	refsOf := func(res *ctypes.ResultTxSearch) []ref {
		out := make([]ref, 0, len(res.Txs))
		for _, r := range res.Txs {
			out = append(out, ref{r.Height, r.Index})
		}
		return out
	}

	perPage := 5

	// Ascending pages of 5: 5 + 5 + 2, reconstructing the full (height, index) order.
	var gotAsc []ref
	for page := 1; page <= 3; page++ {
		p := page
		res, err := env.TxSearch(rctx, q, false, &p, &perPage, "asc")
		require.NoError(t, err)
		require.Equal(t, 12, res.TotalCount, "total count is independent of page")
		gotAsc = append(gotAsc, refsOf(res)...)
	}
	assert.Equal(t, order, gotAsc)

	// The last page must hold exactly the remaining 2 txs.
	lastPage := 3
	last, err := env.TxSearch(rctx, q, false, &lastPage, &perPage, "asc")
	require.NoError(t, err)
	assert.Len(t, last.Txs, 2)

	// Descending first page returns the highest (height, index) tuples.
	firstPage, three := 1, 3
	desc, err := env.TxSearch(rctx, q, false, &firstPage, &three, "desc")
	require.NoError(t, err)
	require.Len(t, desc.Txs, 3)
	assert.Equal(t, []ref{{3, 3}, {3, 2}, {3, 1}}, refsOf(desc))

	// Invalid order_by is rejected.
	_, err = env.TxSearch(rctx, q, false, &firstPage, &perPage, "sideways")
	assert.Error(t, err)

	// Out-of-range page is rejected (validatePage moved into the indexer).
	badPage := 99
	_, err = env.TxSearch(rctx, q, false, &badPage, &perPage, "asc")
	assert.Error(t, err)
}
