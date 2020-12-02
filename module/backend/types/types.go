package types

import (
	backendtypes "github.com/okex/okexchain/x/backend/types"
)

// const
const (
	ModuleName = backendtypes.ModuleName

	ClosedOrdersPath = "custom/backend/orders/closed"
	DealsPath        = "custom/backend/deals"
	TransactionsPath = "custom/backend/txs"
)

type (
	Ticker      = backendtypes.Ticker
	MatchResult = backendtypes.MatchResult
	Order       = backendtypes.Order
)

// Deal - structure of deal query result
type Deal struct {
	Timestamp   int64   `json:"timestamp"`
	BlockHeight int64   `json:"block_height"`
	OrderID     string  `json:"order_id"`
	Sender      string  `json:"sender"`
	Product     string  `json:"product"`
	Side        string  `json:"side"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"volume"`
	Fee         string  `json:"fee"`
}

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
