package types

import (
	backendtypes "github.com/okex/okexchain/x/backend/types"
)

// const
const (
	ModuleName = backendtypes.ModuleName

	RecentTxRecordPath = "custom/backend/matches"
	OpenOrdersPath     = "custom/backend/orders/open"
	ClosedOrdersPath   = "custom/backend/orders/closed"
	DealsPath          = "custom/backend/deals"
	TransactionsPath   = "custom/backend/txs"
)

type (
	Ticker = backendtypes.Ticker
)

// MatchResult - structure for recent tx record
type MatchResult struct {
	Timestamp   int64   `json:"timestamp"`
	BlockHeight int64   `json:"block_height"`
	Product     string  `json:"product"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"volume"`
}

// BaseResponse - structure for base response of data
type BaseResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	DetailMsg string      `json:"detail_msg"`
	Data      interface{} `json:"data"`
}

// ParamPage - structure of page params
type ParamPage struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

// ListDataRes - structure of list data in the list response
type ListDataRes struct {
	Data      interface{} `json:"data"`
	ParamPage ParamPage   `json:"param_page"`
}

// ListResponse - structure for list response of data
type ListResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	DetailMsg string      `json:"detail_msg"`
	Data      ListDataRes `json:"data"`
}

// Order - structure of order query result
type Order struct {
	TxHash         string `json:"txhash"`
	OrderID        string `json:"order_id"`
	Sender         string `json:"sender"`
	Product        string `json:"product"`
	Side           string `json:"side"`
	Price          string `json:"price"`
	Quantity       string `json:"quantity"`
	Status         int64  `json:"status"`
	FilledAvgPrice string `json:"filled_avg_price"`
	RemainQuantity string `json:"remain_quantity"`
	Timestamp      int64  `json:"timestamp"`
}

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
