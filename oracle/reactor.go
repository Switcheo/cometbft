package oracle

import (
	"encoding/hex"
	"fmt"
	"math"
	"time"

	"github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/proxy"

	"github.com/cometbft/cometbft/crypto"

	cs "github.com/cometbft/cometbft/consensus"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/oracle/service/runner"
	oracletypes "github.com/cometbft/cometbft/oracle/service/types"
	"github.com/cometbft/cometbft/p2p"
	oracleproto "github.com/cometbft/cometbft/proto/tendermint/oracle"
	"github.com/cometbft/cometbft/types"
	"github.com/sirupsen/logrus"
)

const (
	OracleChannel = byte(0x42)

	// PeerCatchupSleepIntervalMS defines how much time to sleep if a peer is behind
	PeerCatchupSleepIntervalMS = 100

	MaxActiveIDs = math.MaxUint16
)

// Reactor handles mempool tx broadcasting amongst peers.
// It maintains a map from peer ID to counter, to prevent gossiping txs to the
// peers you received it from.
type Reactor struct {
	p2p.BaseReactor
	OracleInfo     *oracletypes.OracleInfo
	ids            *oracleIDs
	ConsensusState *cs.State
}

// NewReactor returns a new Reactor with the given config and mempool.
func NewReactor(config *config.OracleConfig, pubKey crypto.PubKey, privValidator types.PrivValidator, proxyApp proxy.AppConnConsensus) *Reactor {
	gossipVoteBuffer := &oracletypes.GossipVoteBuffer{
		Buffer: make(map[string]*oracleproto.GossipedVotes),
	}
	unsignedVoteBuffer := &oracletypes.UnsignedVoteBuffer{
		Buffer: []*oracleproto.Vote{},
	}

	oracleInfo := &oracletypes.OracleInfo{
		Config:             config,
		UnsignedVoteBuffer: unsignedVoteBuffer,
		GossipVoteBuffer:   gossipVoteBuffer,
		SignVotesChan:      make(chan *oracleproto.Vote, 1024),
		PubKey:             pubKey,
		PrivValidator:      privValidator,
		ProxyApp:           proxyApp,
		BlockTimestamps:    []int64{},
	}

	oracleR := &Reactor{
		OracleInfo: oracleInfo,
		ids:        newOracleIDs(),
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
	go func() {
		runner.Run(oracleR.OracleInfo, oracleR.ConsensusState)
	}()
	return nil
}

// GetChannels implements Reactor by returning the list of channels for this
// reactor.
func (oracleR *Reactor) GetChannels() []*p2p.ChannelDescriptor {
	// only gossip votes with a max size of x, where x = Config.MaxGossipMsgSize
	messageCap := oracleR.OracleInfo.Config.MaxGossipMsgSize

	return []*p2p.ChannelDescriptor{
		{
			ID:                  OracleChannel,
			Priority:            5,
			RecvMessageCapacity: messageCap,
			MessageType:         &oracleproto.GossipedVotes{},
		},
	}
}

// AddPeer implements Reactor.
// It starts a broadcast routine ensuring all txs are forwarded to the given peer.
func (oracleR *Reactor) AddPeer(peer p2p.Peer) {
	go func() {
		oracleR.broadcastVoteRoutine(peer)
	}()
}

// RemovePeer implements Reactor.
func (oracleR *Reactor) RemovePeer(peer p2p.Peer, _ interface{}) {
	oracleR.ids.Reclaim(peer)
	// broadcast routine checks if peer is gone and returns
}

// // Receive implements Reactor.
func (oracleR *Reactor) Receive(e p2p.Envelope) {
	oracleR.Logger.Debug("Receive", "src", e.Src, "chId", e.ChannelID, "msg", e.Message)
	switch msg := e.Message.(type) {
	case *oracleproto.GossipedVotes:
		// verify sig of incoming gossip vote, throw if verification fails
		_, val := oracleR.ConsensusState.Validators.GetByAddress(msg.Validator)
		if val == nil {
			logrus.Debugf("validator: %v not found in validator set, skipping gossip", hex.EncodeToString(msg.Validator))
			return
		}
		pubKey := val.PubKey

		// skip if its own buffer
		if oracleR.OracleInfo.PubKey.Equals(pubKey) {
			return
		}

		if success := pubKey.VerifySignature(types.OracleVoteSignBytes(oracleR.ConsensusState.GetState().ChainID, msg), msg.Signature); !success {
			logrus.Errorf("failed signature verification for validator: %v, skipping gossip", val.Address)
			return
		}

		preLockTime := time.Now().UnixMilli()
		oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.Lock()
		currentGossipVote, ok := oracleR.OracleInfo.GossipVoteBuffer.Buffer[pubKey.Address().String()]

		if !ok {
			// first gossipVote entry from this validator
			oracleR.OracleInfo.GossipVoteBuffer.Buffer[pubKey.Address().String()] = msg
		} else {
			// existing gossipVote entry from this validator
			previousTimestamp := currentGossipVote.SignedTimestamp
			newTimestamp := msg.SignedTimestamp
			// only replace if the gossipVote received has a later timestamp than our current one
			if newTimestamp > previousTimestamp {
				oracleR.OracleInfo.GossipVoteBuffer.Buffer[pubKey.Address().String()] = msg
			}
		}
		oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.Unlock()
		postLockTime := time.Now().UnixMilli()
		diff := postLockTime - preLockTime
		if diff > 100 {
			logrus.Warnf("WARNING!!! Receiving gossip lock took %v milliseconds", diff)
		}
	default:
		logrus.Warn("unknown message type", "src", e.Src, "chId", e.ChannelID, "msg", e.Message)
		oracleR.Switch.StopPeerForError(e.Src, fmt.Errorf("oracle cannot handle message of type: %T", e.Message))
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
	// gossip votes every x milliseconds, where x = Config.GossipInterval
	interval := oracleR.OracleInfo.Config.GossipInterval

	for {
		// In case of both next.NextWaitChan() and peer.Quit() are variable at the same time
		if !oracleR.IsRunning() || !peer.IsRunning() {
			return
		}
		select {
		case <-peer.Quit():
			return
		case <-oracleR.Quit():
			return
		default:
		}

		// Make sure the peer is up to date.
		_, ok := peer.Get(types.PeerStateKey).(PeerState)
		if !ok {
			// Peer does not have a state yet. We set it in the consensus reactor, but
			// when we add peer in Switch, the order we call reactors#AddPeer is
			// different every time due to us using a map. Sometimes other reactors
			// will be initialized before the consensus reactor. We should wait a few
			// milliseconds and retry.
			time.Sleep(PeerCatchupSleepIntervalMS * time.Millisecond)
			continue
		}

		// only gossip votes that are younger than the latestAllowableTimestamp, which is the max(earliest block timestamp collected, current time - maxOracleGossipAge)
		latestAllowableTimestamp := time.Now().Unix() - int64(oracleR.OracleInfo.Config.MaxOracleGossipAge)
		if len(oracleR.OracleInfo.BlockTimestamps) == oracleR.OracleInfo.Config.MaxOracleGossipBlocksDelayed && oracleR.OracleInfo.BlockTimestamps[0] > latestAllowableTimestamp {
			latestAllowableTimestamp = oracleR.OracleInfo.BlockTimestamps[0]
		}

		preLockTime := time.Now().UnixMilli()
		oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.RLock()
		votes := []*oracleproto.GossipedVotes{}
		for _, gossipVote := range oracleR.OracleInfo.GossipVoteBuffer.Buffer {
			// stop sending gossip votes that have passed the maxGossipVoteAge
			if gossipVote.SignedTimestamp < latestAllowableTimestamp {
				continue
			}

			votes = append(votes, gossipVote)
		}
		oracleR.OracleInfo.GossipVoteBuffer.UpdateMtx.RUnlock()
		postLockTime := time.Now().UnixMilli()
		diff := postLockTime - preLockTime
		if diff > 100 {
			logrus.Warnf("WARNING!!! Sending gossip lock took %v milliseconds", diff)
		}

		for _, vote := range votes {
			success := peer.Send(p2p.Envelope{
				ChannelID: OracleChannel,
				Message:   vote,
			})
			if !success {
				break
			}
		}
		time.Sleep(interval)
	}
}
