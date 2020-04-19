package utils

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/module/tendermint/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// ParseBlock converts raw tendermint block type to the one gosdk requires
func ParseBlock(cdc sdk.SDKCodec, pTmBlock *tmtypes.Block) (block types.Block, err error) {
	var stdTxs []sdk.StdTx
	for _, txBytes := range pTmBlock.Txs {
		var stdTx sdk.StdTx
		if err = cdc.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx); err != nil {
			return block, fmt.Errorf("failed. unmarshal tx info from tendermint block query error: %s", err)
		}
		stdTxs = append(stdTxs, stdTx)
	}

	return types.NewBlock(pTmBlock.Header, types.NewData(stdTxs), pTmBlock.Evidence, *pTmBlock.LastCommit), err
}

// ParseBlockResults converts raw tendermint block result type to the one gosdk requires
func ParseBlockResults(pTmBlockResults *ctypes.ResultBlockResults) types.BlockResults {
	// build ResponseDeliverTx
	respDeliverTxsLen := len(pTmBlockResults.Results.DeliverTx)
	respDeliverTxs := make([]types.ResponseDeliverTx, respDeliverTxsLen)
	for i := 0; i < respDeliverTxsLen; i++ {
		respDeliverTxs[i] = parseResponseDeliverTx(pTmBlockResults.Results.DeliverTx[i])
	}

	return types.BlockResults{
		Height: pTmBlockResults.Height,
		Results: types.ABCIResponses{
			DeliverTx:  respDeliverTxs,
			BeginBlock: parseResponseBeginBlock(pTmBlockResults.Results.BeginBlock),
			EndBlock:   parseResponseEndBlock(pTmBlockResults.Results.EndBlock),
		},
	}
}

func parseResponseDeliverTx(pTmRespDeliverTx *abci.ResponseDeliverTx) types.ResponseDeliverTx {
	return types.ResponseDeliverTx{
		Code:      pTmRespDeliverTx.Code,
		Data:      pTmRespDeliverTx.Data,
		Log:       pTmRespDeliverTx.Log,
		Info:      pTmRespDeliverTx.Info,
		GasWanted: pTmRespDeliverTx.GasWanted,
		GasUsed:   pTmRespDeliverTx.GasUsed,
		Events:    parseEvents(pTmRespDeliverTx.Events),
		Codespace: pTmRespDeliverTx.Codespace,
	}
}

func parseResponseBeginBlock(pTmRespBeginBlock *abci.ResponseBeginBlock) types.ResponseBeginBlock {
	return types.ResponseBeginBlock{
		Events: parseEvents(pTmRespBeginBlock.Events),
	}
}

func parseResponseEndBlock(pTmRespEndBlock *abci.ResponseEndBlock) types.ResponseEndBlock {
	return types.ResponseEndBlock{
		ValidatorUpdates:      parseValidatorUpdates(pTmRespEndBlock.ValidatorUpdates),
		ConsensusParamUpdates: parseConsensusParams(pTmRespEndBlock.ConsensusParamUpdates),
		Events:                parseEvents(pTmRespEndBlock.Events),
	}
}

func parseEvents(tmEvents []abci.Event) []types.Event {
	eventsLens := len(tmEvents)
	events := make([]types.Event, eventsLens)
	for i := 0; i < eventsLens; i++ {
		events[i] = types.Event{
			Type:       tmEvents[i].Type,
			Attributes: parseKVPairs(tmEvents[i].Attributes),
		}
	}

	return events
}

func parseKVPairs(tmKVPair []common.KVPair) []types.KVPair {
	kvPairsLen := len(tmKVPair)
	kvPairs := make([]types.KVPair, kvPairsLen)
	for i := 0; i < kvPairsLen; i++ {
		kvPairs[i] = types.KVPair{
			Key:   tmKVPair[i].Key,
			Value: tmKVPair[i].Value,
		}
	}

	return kvPairs
}

func parseValidatorUpdates(tmValUpdates []abci.ValidatorUpdate) []types.ValidatorUpdate {
	valUpdatesLen := len(tmValUpdates)
	valUpdates := make([]types.ValidatorUpdate, valUpdatesLen)
	for i := 0; i < valUpdatesLen; i++ {
		valUpdates[i] = types.ValidatorUpdate{
			PubKey: types.PubKey{
				Type: tmValUpdates[i].PubKey.Type,
				Data: tmValUpdates[i].PubKey.Data,
			},
			Power: tmValUpdates[i].Power,
		}
	}

	return valUpdates
}

func parseConsensusParams(tmConsParams *abci.ConsensusParams) types.ConsensusParams {
	if tmConsParams == nil {
		return types.ConsensusParams{}
	}

	return types.ConsensusParams{
		Block: types.BlockParams{
			MaxBytes: tmConsParams.Block.MaxBytes,
			MaxGas:   tmConsParams.Block.MaxGas,
		},
		Evidence: types.EvidenceParams{
			MaxAge: tmConsParams.Evidence.MaxAge,
		},
		Validator: types.ValidatorParams{
			PubKeyTypes: tmConsParams.Validator.PubKeyTypes,
		},
	}
}
