package exposed

import (
	"github.com/okex/okchain-go-sdk/module/backend/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// Backend shows the expected behavior for inner backend client
type Backend interface {
	sdk.Module
	BackendQuery
}

// BackendQuery shows the expected query behavior for inner backend client
type BackendQuery interface {
	QueryCandles(product string, granularity, size int) ([][]string, error)
	QueryTickers(count ...int) ([]types.Ticker, error)
	QueryRecentTxRecord(product string, start, end, page, perPage int) ([]types.MatchResult, error)
	QueryOpenOrders(addrStr, product, side string, start, end, page, perPage int) ([]types.Order, error)
}
