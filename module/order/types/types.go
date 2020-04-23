package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	ModuleName = "order"

	DepthbookPath   = "custom/order/depthbook"
	OrderDetailPath = "custom/order/detail"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for token module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgNewOrders{}, "okchain/order/MsgNew")
	cdc.RegisterConcrete(MsgCancelOrders{}, "okchain/order/MsgCancel")
}

// OrderItem - structure for a item in MsgNewOrders
type OrderItem struct {
	Product  string  `json:"product"`
	Side     string  `json:"side"`
	Price    sdk.Dec `json:"price"`
	Quantity sdk.Dec `json:"quantity"`
}

// NewOrderItem creates a new instance of OrderItem
func NewOrderItem(product string, side string, price string, quantity string) OrderItem {
	return OrderItem{
		Product:  product,
		Side:     side,
		Price:    sdk.MustNewDecFromStr(price),
		Quantity: sdk.MustNewDecFromStr(quantity),
	}
}

// BuildOrderItems returns the set of OrderItem
// params must be checked by function CheckNewOrderParams
func BuildOrderItems(products, sides, prices, quantities []string) []OrderItem {
	productsLen := len(products)
	orderItems := make([]OrderItem, productsLen)
	for i := 0; i < productsLen; i++ {
		orderItems[i] = NewOrderItem(products[i], sides[i], prices[i], quantities[i])
	}

	return orderItems
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
