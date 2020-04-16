package exposed

import (
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
}
