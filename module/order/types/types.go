package types

import (
	orderkeeper "github.com/okex/exchain/x/order/keeper"
	ordertypes "github.com/okex/exchain/x/order/types"
)

// const
const (
	ModuleName = ordertypes.ModuleName
)

type (
	BookRes     = orderkeeper.BookRes
	OrderDetail = ordertypes.Order
)
