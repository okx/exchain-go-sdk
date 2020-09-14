package exposed

import (
	"github.com/okex/okexchain-go-sdk/module/backend/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
)

// Backend shows the expected behavior for inner backend client
type Backend interface {
	sdk.Module
	BackendQuery
}

// BackendQuery shows the expected query behavior for inner backend client
type BackendQuery interface {
	QueryCandles(product string, granularity, size int) ([][]string, error)
	QueryTickers(product string, count ...int) ([]types.Ticker, error)
	QueryRecentTxRecord(product string, start, end, page, perPage int) ([]types.MatchResult, error)
	QueryOpenOrders(addrStr, product, side string, start, end, page, perPage int) ([]types.Order, error)
	QueryClosedOrders(addrStr, product, side string, start, end, page, perPage int) ([]types.Order, error)
	QueryDeals(addrStr, product, side string, start, end, page, perPage int) ([]types.Deal, error)
	QueryTransactions(addrStr string, typeCode, start, end, page, perPage int) ([]types.Transaction, error)
}
