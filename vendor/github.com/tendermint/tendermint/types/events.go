package types

import (
	abci "github.com/ok-chain/gosdk/types/abci"
	"github.com/tendermint/go-amino"
)

// TMEventData implements events.EventData.
type TMEventData interface {
	// empty interface
}

func RegisterEventDatas(cdc *amino.Codec) {
	cdc.RegisterInterface((*TMEventData)(nil), nil)
	cdc.RegisterConcrete(EventDataNewBlock{}, "tendermint/event/NewBlock", nil)
	cdc.RegisterConcrete(EventDataNewBlockHeader{}, "tendermint/event/NewBlockHeader", nil)
	cdc.RegisterConcrete(EventDataTx{}, "tendermint/event/Tx", nil)
	cdc.RegisterConcrete(EventDataRoundState{}, "tendermint/event/RoundState", nil)
	cdc.RegisterConcrete(EventDataNewRound{}, "tendermint/event/NewRound", nil)
	cdc.RegisterConcrete(EventDataCompleteProposal{}, "tendermint/event/CompleteProposal", nil)
	cdc.RegisterConcrete(EventDataVote{}, "tendermint/event/Vote", nil)
	cdc.RegisterConcrete(EventDataValidatorSetUpdates{}, "tendermint/event/ValidatorSetUpdates", nil)
	cdc.RegisterConcrete(EventDataString(""), "tendermint/event/ProposalString", nil)
}

// Most event messages are basic types (a block, a transaction)
// but some (an input to a call tx or a receive) are more exotic

type EventDataNewBlock struct {
	Block *Block `json:"block"`

	ResultBeginBlock abci.ResponseBeginBlock `json:"result_begin_block"`
	ResultEndBlock   abci.ResponseEndBlock   `json:"result_end_block"`
}

// light weight event for benchmarking
type EventDataNewBlockHeader struct {
	Header Header `json:"header"`

	ResultBeginBlock abci.ResponseBeginBlock `json:"result_begin_block"`
	ResultEndBlock   abci.ResponseEndBlock   `json:"result_end_block"`
}

// All txs fire EventDataTx
type EventDataTx struct {
	TxResult
}

// NOTE: This goes into the replay WAL
type EventDataRoundState struct {
	Height int64  `json:"height"`
	Round  int    `json:"round"`
	Step   string `json:"step"`
}

type ValidatorInfo struct {
	Address Address `json:"address"`
	Index   int     `json:"index"`
}

type EventDataNewRound struct {
	Height int64  `json:"height"`
	Round  int    `json:"round"`
	Step   string `json:"step"`

	Proposer ValidatorInfo `json:"proposer"`
}

type EventDataCompleteProposal struct {
	Height int64  `json:"height"`
	Round  int    `json:"round"`
	Step   string `json:"step"`

	BlockID BlockID `json:"block_id"`
}

type EventDataVote struct {
	Vote *Vote
}

type EventDataString string

type EventDataValidatorSetUpdates struct {
	ValidatorUpdates []*Validator `json:"validator_updates"`
}