package utils

import (
	"github.com/okex/okexchain-go-sdk/module/tendermint/types"
	//"github.com/tendermint/tendermint/libs/common"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// ParseTxsResult converts raw tendermint txs result type to the one gosdk requires
func ParseTxsResult(pTmTxsResult *ctypes.ResultTxSearch) types.ResultTxs {
	txsLen := len(pTmTxsResult.Txs)
	txsResult := make([]types.ResultTx, txsLen)
	for i := 0; i < txsLen; i++ {
		//txsResult[i] = ParseTxResult(pTmTxsResult.Txs[i])
	}

	return types.ResultTxs{
		Txs:        txsResult,
		TotalCount: pTmTxsResult.TotalCount,
	}
}