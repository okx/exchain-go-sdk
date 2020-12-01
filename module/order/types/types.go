package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	orderkeeper "github.com/okex/okexchain/x/order/keeper"
	ordertypes "github.com/okex/okexchain/x/order/types"
)

// const
const (
	ModuleName = ordertypes.ModuleName
	OrderDetailPath = "custom/order/detail"
)

type (
	BookRes = orderkeeper.BookRes
)

// OrderDetail - structure for the detail info of an order
type OrderDetail struct {
	TxHash            string         `json:"txhash"`
	OrderID           string         `json:"order_id"`
	Sender            sdk.AccAddress `json:"sender"`
	Product           string         `json:"product"`
	Side              string         `json:"side"`
	Price             sdk.Dec        `json:"price"`
	Quantity          sdk.Dec        `json:"quantity"`
	Status            int64          `json:"status"`
	FilledAvgPrice    sdk.Dec        `json:"filled_avg_price"`
	RemainQuantity    sdk.Dec        `json:"remain_quantity"`
	RemainLocked      sdk.Dec        `json:"remain_locked"`
	Timestamp         int64          `json:"timestamp"`
	OrderExpireBlocks int64          `json:"order_expire_blocks"`
	FeePerBlock       sdk.DecCoin    `json:"fee_per_block"`
	ExtraInfo         string         `json:"extra_info"`
}
