package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	ModuleName = "order"
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
