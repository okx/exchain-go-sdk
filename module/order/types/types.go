package types

import (
	orderkeeper "github.com/okex/okexchain/x/order/keeper"
	ordertypes "github.com/okex/okexchain/x/order/types"
)

// const
const (
	ModuleName = ordertypes.ModuleName
)

type (
	BookRes     = orderkeeper.BookRes
	OrderDetail = ordertypes.Order
)
