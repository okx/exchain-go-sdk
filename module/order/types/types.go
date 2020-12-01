package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/order"
	ordertypes "github.com/okex/okexchain/x/order/types"
)

// const
const (
	ModuleName = ordertypes.ModuleName

	DepthbookPath   = "custom/order/depthbook"
	OrderDetailPath = "custom/order/detail"
)

var (
	msgCdc = gosdktypes.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for token module
func RegisterCodec(cdc *codec.Codec) {
	order.RegisterCodec(cdc)
}



// OrderResult - structure for the filter of orderID
type OrderResult struct {
	Code    uint32 `json:"code"`
	Message string `json:"msg"`
	OrderID string `json:"orderid"`
}

// BookRes - structure of depthbook
type BookRes struct {
	Asks []BookResItem `json:"asks"`
	Bids []BookResItem `json:"bids"`
}

// BookResItem - structure of an item in BookRes
type BookResItem struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

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
