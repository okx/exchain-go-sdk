package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

// const
const (
	ModuleName = "tendermint"
)

// Block - structure for the result of block query
type Block struct {
	tmtypes.Header `json:"header"`
	Data           `json:"data"`
	Evidence       tmtypes.EvidenceData `json:"evidence"`
	LastCommit     tmtypes.Commit       `json:"last_commit"`
}

// NewBlock creates a new instance of Block
func NewBlock(header tmtypes.Header, data Data, evidence tmtypes.EvidenceData, lastCommit tmtypes.Commit) Block {
	return Block{
		Header:     header,
		Data:       data,
		Evidence:   evidence,
		LastCommit: lastCommit,
	}
}

// Data - structure of the stdTxs in a block
type Data struct {
	Txs []sdk.StdTx `json:"txs"`
}

// NewData creates a new instance of Data
func NewData(stdTxs []sdk.StdTx) Data {
	return Data{
		Txs: stdTxs,
	}
}

// BlockResults - structure for ABCI results from a block
type BlockResults struct {
	Height  int64         `json:"height"`
	Results ABCIResponses `json:"results"`
}

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

// ResultCommit - structure for the commit info
type ResultCommit struct {
	SignedHeader
	CanonicalCommit bool
}

// SignedHeader is a header along with the commits that prove it
// It is the basis of the lite client
type SignedHeader struct {
	tmtypes.Header
	Commit tmtypes.Commit
}

// ResultValidators - structure for the validators info on a specific height
type ResultValidators struct {
	BlockHeight int64
	Validators  []Validator
}

// Validator - structure of the volatile state for each Validator
type Validator struct {
	Address          tmtypes.Address
	PubKey           crypto.PubKey
	VotingPower      int64
	ProposerPriority int64
}
