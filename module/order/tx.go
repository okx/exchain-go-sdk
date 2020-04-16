package order

import (
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/module/order/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/params"
	"strings"
)

// NewOrders places orders with some detail info
func (oc orderClient) NewOrders(fromInfo keys.Info, passWd, products, sides, prices, quantities, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	productStrs := strings.Split(products, ",")
	sideStrs := strings.Split(sides, ",")
	priceStrs := strings.Split(prices, ",")
	quantityStrs := strings.Split(quantities, ",")
	if err = params.CheckNewOrderParams(fromInfo, passWd, productStrs, sideStrs, priceStrs, quantityStrs); err != nil {
		return
	}

	orderItems := types.BuildOrderItems(productStrs, sideStrs, priceStrs, quantityStrs)
	msg := types.NewMsgNewOrders(fromInfo.GetAddress(), orderItems)

	return oc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// CancelOrders cancels orders by orderIDs
func (oc orderClient) CancelOrders(fromInfo keys.Info, passWd, orderIDs, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	orderIDStrs := strings.Split(orderIDs, ",")
	if err = params.CheckCancelOrderParams(fromInfo, passWd, orderIDStrs); err != nil {
		return
	}

	msg := types.NewMsgCancelOrders(fromInfo.GetAddress(), orderIDStrs)

	return oc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}
