package utils

import (
	"github.com/tendermint/tendermint/libs/kv"

	"github.com/okex/okexchain-go-sdk/module/tendermint/types"
	abci "github.com/tendermint/tendermint/abci/types"
	//"github.com/tendermint/tendermint/libs/common"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)


// ParseTxResult converts raw tendermint tx result type to the one gosdk requires
func ParseTxResult(pTmTxResult *ctypes.ResultTx) types.ResultTx {
	return types.ResultTx{
		Hash:     pTmTxResult.Hash,
		Height:   pTmTxResult.Height,
		Index:    pTmTxResult.Index,
		TxResult: parseResponseDeliverTx(&pTmTxResult.TxResult),
		Tx:       pTmTxResult.Tx,
		Proof:    pTmTxResult.Proof,
	}
}

// ParseTxsResult converts raw tendermint txs result type to the one gosdk requires
func ParseTxsResult(pTmTxsResult *ctypes.ResultTxSearch) types.ResultTxs {
	txsLen := len(pTmTxsResult.Txs)
	txsResult := make([]types.ResultTx, txsLen)
	for i := 0; i < txsLen; i++ {
		txsResult[i] = ParseTxResult(pTmTxsResult.Txs[i])
	}

	return types.ResultTxs{
		Txs:        txsResult,
		TotalCount: pTmTxsResult.TotalCount,
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

func parseKVPairs(tmKVPair []kv.Pair) []types.KVPair {
	//kvPairsLen := len(tmKVPair)
	//kvPairs := make([]types.KVPair, kvPairsLen)
	//for i := 0; i < kvPairsLen; i++ {
	//	kvPairs[i] = types.KVPair{
	//		Key:   tmKVPair[i].Key,
	//		Value: tmKVPair[i].Value,
	//	}
	//}
	//
	//return kvPairs
	return nil
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
	//if tmConsParams == nil {
	//	return types.ConsensusParams{}
	//}
	//
	//return types.ConsensusParams{
	//	Block: types.BlockParams{
	//		MaxBytes: tmConsParams.Block.MaxBytes,
	//		MaxGas:   tmConsParams.Block.MaxGas,
	//	},
	//	Evidence: types.EvidenceParams{
	//		MaxAge: tmConsParams.Evidence.MaxAge,
	//	},
	//	Validator: types.ValidatorParams{
	//		PubKeyTypes: tmConsParams.Validator.PubKeyTypes,
	//	},
	//}
	return types.ConsensusParams{}
}
