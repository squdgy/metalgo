package benchlist

import (
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/constants"
)

// Manager provides an interface for a benchlist to register whether
// queries have been successful or unsuccessful and place validators with
// consistently failing queries on a benchlist to prevent waiting up to
// the full network timeout for their responses.
type Manager interface {
	// RegisterQuery registers a sent query and returns whether the query is subject to benchlist
	RegisterQuery(ids.ID, ids.ShortID, uint32, constants.MsgType) bool
	// RegisterResponse registers the response to a query message
	RegisterResponse(ids.ID, ids.ShortID, uint32)
	// QueryFailed registers that a query did not receive a response within our synchrony bound
	QueryFailed(ids.ID, ids.ShortID, uint32)
}

// Config defines the configuration for a benchlist
type Config struct {
	Validators         validators.Manager
	Threshold          int
	Duration           time.Duration
	MaxPortion         float64
	PeerSummaryEnabled bool
}

type benchlistManager struct {
	config    *Config
	ctxLookup snow.ContextLookup
	// Chain ID --> benchlist for that chain
	chainBenchlists map[[32]byte]QueryBenchlist
}

// NewManager returns a manager for chain-specific query benchlisting
func NewManager(config *Config, ctxLookup snow.ContextLookup) Manager {
	return &benchlistManager{
		config:          config,
		ctxLookup:       ctxLookup,
		chainBenchlists: make(map[[32]byte]QueryBenchlist),
	}
}

// RegisterQuery implements the Manager interface
func (bm *benchlistManager) RegisterQuery(chainID ids.ID, validatorID ids.ShortID, requestID uint32, msgType constants.MsgType) bool {
	key := chainID.Key()
	chain, exists := bm.chainBenchlists[key]
	if !exists {
		vdrs, ok := bm.config.Validators.GetValidatorsByChain(chainID)
		if !ok {
			return false
		}
		ctx, exists := bm.ctxLookup.GetContext(chainID)
		if !exists {
			return false
		}
		chain = NewQueryBenchlist(vdrs, ctx, bm.config.Threshold, bm.config.Duration, bm.config.MaxPortion, bm.config.PeerSummaryEnabled)
		bm.chainBenchlists[key] = chain
	}

	return chain.RegisterQuery(validatorID, requestID, msgType)
}

// RegisterResponse implements the Manager interface
func (bm *benchlistManager) RegisterResponse(chainID ids.ID, validatorID ids.ShortID, requestID uint32) {
	chain, exists := bm.chainBenchlists[chainID.Key()]
	if !exists {
		return
	}

	chain.RegisterResponse(validatorID, requestID)
}

// QueryFailed implements the Manager interface
func (bm *benchlistManager) QueryFailed(chainID ids.ID, validatorID ids.ShortID, requestID uint32) {
	chain, exists := bm.chainBenchlists[chainID.Key()]
	if !exists {
		return
	}

	chain.QueryFailed(validatorID, requestID)
}

type noBenchlist struct{}

// NewNoBenchlist returns an empty benchlist that will never stop any queries
func NewNoBenchlist() Manager { return &noBenchlist{} }

// RegisterQuery ...
func (b *noBenchlist) RegisterQuery(chainID ids.ID, validatorID ids.ShortID, requestID uint32, msgType constants.MsgType) bool {
	return true
}

// RegisterResponse ...
func (b *noBenchlist) RegisterResponse(chainID ids.ID, validatorID ids.ShortID, requestID uint32) {}

// QueryFailed ...
func (b *noBenchlist) QueryFailed(chainID ids.ID, validatorID ids.ShortID, requestID uint32) {}
