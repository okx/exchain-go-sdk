package types

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// const
const (
	ModuleName = "tendermint"

	EventFormat = "{eventType}.{eventAttribute}={value}"
)

type (
	Block              = tmtypes.Block
	ResultBlockResults = ctypes.ResultBlockResults
	ResultCommit       = ctypes.ResultCommit
	ResultValidators   = ctypes.ResultValidators
	ResultTx           = ctypes.ResultTx
)

// ABCIResponses - structure for the responses of the various ABCI calls during block processing
type ABCIResponses struct {
	DeliverTx  []ResponseDeliverTx
	BeginBlock ResponseBeginBlock
	EndBlock   ResponseEndBlock
}

// ResponseDeliverTx - structure for the deliver tx response info
type ResponseDeliverTx struct {
	Code      uint32
	Data      []byte
	Log       string
	Info      string
	GasWanted int64
	GasUsed   int64
	Events    []Event
	Codespace string
}

// Event - structure for the event info in ResponseDeliverTx
type Event struct {
	Type       string
	Attributes []KVPair
}

// KVPair - structure for the kv pair info in Event
type KVPair struct {
	Key   []byte
	Value []byte
}

// ResponseBeginBlock - structure for the begin block response info
type ResponseBeginBlock struct {
	Events []Event
}

// ResponseEndBlock - structure for the end block response info
type ResponseEndBlock struct {
	ValidatorUpdates      []ValidatorUpdate
	ConsensusParamUpdates ConsensusParams
	Events                []Event
}

// ValidatorUpdate - structure for the validators update info in the end block
type ValidatorUpdate struct {
	PubKey PubKey
	Power  int64
}

// PubKey - structure for pubkey info in the ValidatorUpdates
type PubKey struct {
	Type string
	Data []byte
}

// ConsensusParams - structure for all consensus-relevant parameters
type ConsensusParams struct {
	Block     BlockParams
	Evidence  EvidenceParams
	Validator ValidatorParams
}

// BlockParams - structure for the limits on the block size and timestamp
type BlockParams struct {
	// Note: must be greater than 0
	MaxBytes int64
	// Note: must be greater or equal to -1
	MaxGas int64
}

// EvidenceParams - structure for the limits on the evidence
type EvidenceParams struct {
	// Note: must be greater than 0
	MaxAge int64
}

// ValidatorParams - structure for limits on validators
type ValidatorParams struct {
	PubKeyTypes []string
}

// ResultTxs - structure of txs result by a specific searching string
type ResultTxs struct {
	Txs        []ResultTx
	TotalCount int
}
