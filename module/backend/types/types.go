package types

import (
	backendtypes "github.com/okex/okexchain/x/backend/types"
)

// const
const (
	ModuleName = backendtypes.ModuleName

	TransactionsPath = "custom/backend/txs"
)

type (
	Ticker      = backendtypes.Ticker
	MatchResult = backendtypes.MatchResult
	Order       = backendtypes.Order
	Deal        = backendtypes.Deal
)

// Transaction - structure of transaction query result
type Transaction struct {
	TxHash    string `json:"txhash"`
	Type      int64  `json:"type"`
	Address   string `json:"address"`
	Symbol    string `json:"symbol"`
	Side      int64  `json:"side"`
	Quantity  string `json:"quantity"`
	Fee       string `json:"fee"`
	Timestamp int64  `json:"timestamp"`
}
