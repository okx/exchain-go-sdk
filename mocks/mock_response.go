package mocks

import (
	sdk "github.com/okex/okexchain-go-sdk/types"
)

// DefaultMockSuccessTxResponse returns the default mock success tx response for transaction testing
func DefaultMockSuccessTxResponse() sdk.TxResponse {
	return sdk.TxResponse{
		Height:    1024,
		TxHash:    "default tx hash",
		Code:      0,
		Data:      "default data",
		RawLog:    "default raw log",
		Logs:      nil,
		Info:      "default info",
		GasWanted: 2048,
		GasUsed:   2048,
		Codespace: "default code space",
		Tx:        nil,
		Timestamp: "default time stamp",
		Events:    nil,
	}
}
