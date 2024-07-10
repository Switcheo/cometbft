package types

import (
	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	oracleproto "github.com/cometbft/cometbft/proto/tendermint/oracle"
	"github.com/cometbft/cometbft/proxy"
	"github.com/cometbft/cometbft/types"
)

// App struct for app
type OracleInfo struct {
	Config             *config.OracleConfig
	UnsignedVoteBuffer *UnsignedVoteBuffer
	GossipVoteBuffer   *GossipVoteBuffer
	SignVotesChan      chan *oracleproto.Vote
	PubKey             crypto.PubKey
	PrivValidator      types.PrivValidator
	StopChannel        chan int
	ProxyApp           proxy.AppConnConsensus
	BlockTimestamps    []int64
}
type GossipVoteBuffer struct {
	Buffer    map[string]*oracleproto.GossipedVotes
	UpdateMtx cmtsync.RWMutex
}

type UnsignedVoteBuffer struct {
	Buffer    []*oracleproto.Vote
	UpdateMtx cmtsync.RWMutex
}

var MainAccountSigPrefix = []byte{0x00}
var SubAccountSigPrefix = []byte{0x01}
