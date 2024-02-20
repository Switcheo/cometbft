package oracle

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	// cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/oracle/service/adapters"
	"github.com/cometbft/cometbft/oracle/service/runner"
	oracletypes "github.com/cometbft/cometbft/oracle/service/types"
	"github.com/cometbft/cometbft/p2p"
	oracleproto "github.com/cometbft/cometbft/proto/tendermint/oracle"
	"github.com/cometbft/cometbft/redis"
	"github.com/cometbft/cometbft/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	OracleChannel = byte(0x42)

	// PeerCatchupSleepIntervalMS defines how much time to sleep if a peer is behind
	PeerCatchupSleepIntervalMS = 100

	// UnknownPeerID is the peer ID to use when running CheckTx when there is
	// no peer (e.g. RPC)
	UnknownPeerID uint16 = 0

	MaxActiveIDs = math.MaxUint16
)

// Reactor handles mempool tx broadcasting amongst peers.
// It maintains a map from peer ID to counter, to prevent gossiping txs to the
// peers you received it from.
type Reactor struct {
	p2p.BaseReactor
	OracleInfo  *oracletypes.OracleInfo
	grpcAddress string
	// config  *cfg.MempoolConfig
	// mempool *CListMempool
	ids *oracleIDs
}

// NewReactor returns a new Reactor with the given config and mempool.
func NewReactor(configPath string, grpcAddress string, pubKey crypto.PubKey, privValidator types.PrivValidator) *Reactor {
	// load oracle.json config if present
	jsonFile, openErr := os.Open(configPath)
	if openErr != nil {
		logrus.Warnf("[oracle] error opening oracle.json config file: %v", openErr)
	}

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		logrus.Warnf("[oracle] error reading oracle.json config file: %v", err)
	}

	var config oracletypes.Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		logrus.Warnf("[oracle] error parsing oracle.json config file: %v", err)
	}

	voteDataBuffer := &oracletypes.VoteDataBuffer{
		Buffer: make(map[uint64]map[string][]*oracleproto.Vote),
	}

	gossipVoteBuffer := &oracletypes.GossipVoteBuffer{
		Buffer: make(map[string]*oracleproto.GossipVote),
	}

	unsignedVoteBuffer := &oracletypes.UnsignedVoteBuffer{
		Buffer: []*oracletypes.UnsignedVotes{},
	}

	oracleInfo := &oracletypes.OracleInfo{
		Oracles:            nil,
		Config:             config,
		VoteDataBuffer:     voteDataBuffer,
		GossipVoteBuffer:   gossipVoteBuffer,
		UnsignedVoteBuffer: unsignedVoteBuffer,
		SignVotesChan:      make(chan *oracleproto.Vote),
		PubKey:             pubKey,
		PrivValidator:      privValidator,
	}

	jsonFile.Close()

	oracleR := &Reactor{
		OracleInfo:  oracleInfo,
		grpcAddress: grpcAddress,
		ids:         newOracleIDs(),
	}
	oracleR.BaseReactor = *p2p.NewBaseReactor("Oracle", oracleR)

	return oracleR
}

// InitPeer implements Reactor by creating a state for the peer.
func (oracleR *Reactor) InitPeer(peer p2p.Peer) p2p.Peer {
	oracleR.ids.ReserveForPeer(peer)
	return peer
}

// SetLogger sets the Logger on the reactor and the underlying mempool.
func (oracleR *Reactor) SetLogger(l log.Logger) {
	oracleR.Logger = l
	oracleR.BaseService.SetLogger(l)
}

// OnStart implements p2p.BaseReactor.
func (oracleR *Reactor) OnStart() error {
	oracleR.OracleInfo.Redis = redis.NewService(0)

	grpcMaxRetryCount := 12
	retryCount := 0
	sleepTime := time.Second
	var client *grpc.ClientConn

	for {
		logrus.Infof("[oracle] trying to connect to grpc with address %s : %d", oracleR.grpcAddress, retryCount)
		if retryCount == grpcMaxRetryCount {
			panic("failed to connect to grpc:grpcClient after 12 tries")
		}
		time.Sleep(sleepTime)

		// reinit otherwise connection will be idle, in idle we can't tell if it's really ready
		var err error
		client, err = grpc.Dial(
			oracleR.grpcAddress,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			panic(err)
		}
		// give it some time to connect after dailing, but not too long as connection can become idle
		time.Sleep(time.Duration(retryCount*int(time.Second) + 1))

		if client.GetState() == connectivity.Ready {
			oracleR.OracleInfo.GrpcClient = client
			break
		}
		client.Close()
		retryCount++
		sleepTime *= 2
	}

	oracleR.OracleInfo.AdapterMap = adapters.GetAdapterMap(oracleR.OracleInfo.GrpcClient, &oracleR.OracleInfo.Redis)
	logrus.Info("[oracle] running oracle service...")
	runner.Run(oracleR.OracleInfo)

	return nil
}

// GetChannels implements Reactor by returning the list of channels for this
// reactor.
func (oracleR *Reactor) GetChannels() []*p2p.ChannelDescriptor {
	// largestTx := make([]byte, oracleR.config.MaxTxBytes)
	// TODO, confirm these params
	return []*p2p.ChannelDescriptor{
		{
			ID:                  OracleChannel,
			Priority:            5,
			RecvMessageCapacity: 1024,
			RecvBufferCapacity:  50 * 4096,
			SendQueueCapacity:   1000,
			MessageType:         &oracleproto.Vote{},
		},
	}
}

// AddPeer implements Reactor.
// It starts a broadcast routine ensuring all txs are forwarded to the given peer.
func (oracleR *Reactor) AddPeer(peer p2p.Peer) {
	// if oracleR.config.Broadcast {
	go oracleR.broadcastVoteRoutine(peer)
	// }
}

// RemovePeer implements Reactor.
func (oracleR *Reactor) RemovePeer(peer p2p.Peer, _ interface{}) {
	oracleR.ids.Reclaim(peer)
	// broadcast routine checks if peer is gone and returns
}

// // Receive implements Reactor.
// // It adds any received transactions to the mempool.
func (oracleR *Reactor) Receive(e p2p.Envelope) {
	oracleR.Logger.Debug("Receive", "src", e.Src, "chId", e.ChannelID, "msg", e.Message)
	switch msg := e.Message.(type) {
	case *oracleproto.GossipVote:
		oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.RLock()
		currentGossipVote, ok := oracleR.OracleInfo.GossipVoteBuffer.Buffer[msg.Validator]
		oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.RUnlock()

		if !ok {
			// first gossipVote entry from this validator
			oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.Lock()
			oracleR.OracleInfo.GossipVoteBuffer.Buffer[msg.Validator] = msg
			oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.Unlock()

			// safe to assume that if gossipVote does not exist in gossipBuffer, it also does not exist in dataBuffer?
			// oracleR.OracleInfo.VoteDataBuffer.UpdateMtx.Lock()
			// for _, vote := range msg.Votes {
			// 	runner.AddVoteToDataBuffer(oracleR.OracleInfo, vote)
			// }
			// oracleR.OracleInfo.VoteDataBuffer.UpdateMtx.Unlock()
		} else {
			// existing gossipVote entry from this validator
			oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.Lock()
			previousTimestamp := currentGossipVote.SignedTimestamp
			newTimestamp := msg.SignedTimestamp
			// only replace if the gossipVote received has a later timestamp than our current one
			if newTimestamp > previousTimestamp {
				oracleR.OracleInfo.GossipVoteBuffer.Buffer[msg.Validator] = msg
			}
			oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.Unlock()
		}
	default:
		oracleR.Logger.Error("unknown message type", "src", e.Src, "chId", e.ChannelID, "msg", e.Message)
		oracleR.Switch.StopPeerForError(e.Src, fmt.Errorf("mempool cannot handle message of type: %T", e.Message))
		return
	}

	// broadcasting happens from go routines per peer
}

// PeerState describes the state of a peer.
type PeerState interface {
	GetHeight() int64
}

// // Send new oracle votes to peer.
func (oracleR *Reactor) broadcastVoteRoutine(peer p2p.Peer) {
	// peerID := oracleR.ids.GetForPeer(peer)

	for {
		// In case of both next.NextWaitChan() and peer.Quit() are variable at the same time
		if !oracleR.IsRunning() || !peer.IsRunning() {
			return
		}
		// This happens because the CElement we were looking at got garbage
		// collected (removed). That is, .NextWait() returned nil. Go ahead and
		// start from the beginning.
		select {
		// case <-oracleR.mempool.TxsWaitChan(): // Wait until a tx is available
		// 	if next = oracleR.mempool.TxsFront(); next == nil {
		// 		continue
		// 	}
		case <-peer.Quit():
			return
		case <-oracleR.Quit():
			return
		default:
		}

		// Make sure the peer is up to date.
		// peerState, ok := peer.Get(types.PeerStateKey).(PeerState)
		// if !ok {
		// 	// Peer does not have a state yet. We set it in the consensus reactor, but
		// 	// when we add peer in Switch, the order we call reactors#AddPeer is
		// 	// different every time due to us using a map. Sometimes other reactors
		// 	// will be initialized before the consensus reactor. We should wait a few
		// 	// milliseconds and retry.
		// 	time.Sleep(PeerCatchupSleepIntervalMS * time.Millisecond)
		// 	continue
		// }

		// // Allow for a lag of 1 block.
		// memTx := next.Value.(*mempoolTx)
		// if peerState.GetHeight() < memTx.Height()-1 {
		// 	time.Sleep(PeerCatchupSleepIntervalMS * time.Millisecond)
		// 	continue
		// }

		// NOTE: Transaction batching was disabled due to
		// https://github.com/tendermint/tendermint/issues/5796

		// if !memTx.isSender(peerID) {
		for _, gossipVote := range oracleR.OracleInfo.GossipVoteBuffer.Buffer {
			success := peer.Send(p2p.Envelope{
				ChannelID: OracleChannel,
				Message:   gossipVote,
			})
			if !success {
				time.Sleep(PeerCatchupSleepIntervalMS * time.Millisecond)
				continue
			}
		}
		// }

		// select {
		// case <-next.NextWaitChan():
		// 	// see the start of the for loop for nil check
		// 	next = next.Next()
		// case <-peer.Quit():
		// 	return
		// case <-oracleR.Quit():
		// 	return
		// }
	}
}

// TxsMessage is a Message containing transactions.
type TxsMessage struct {
	Txs []types.Tx
}

// String returns a string representation of the TxsMessage.
func (m *TxsMessage) String() string {
	return fmt.Sprintf("[TxsMessage %v]", m.Txs)
}
