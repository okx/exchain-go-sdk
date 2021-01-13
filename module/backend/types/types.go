package types

import (
	backendtypes "github.com/okex/okexchain/x/backend/types"
)

// const
const (
	ModuleName = backendtypes.ModuleName
)

type (
	Ticker      = backendtypes.Ticker
	MatchResult = backendtypes.MatchResult
	Order       = backendtypes.Order
	Deal        = backendtypes.Deal
	Transaction = backendtypes.Transaction
)
