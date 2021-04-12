package order

import (
	"fmt"

	"github.com/okex/exchain-go-sdk/module/order/types"
	"github.com/okex/exchain-go-sdk/types/params"
	"github.com/okex/exchain-go-sdk/utils"
	orderkeeper "github.com/okex/exchain/x/order/keeper"
	ordertypes "github.com/okex/exchain/x/order/types"
)

// QueryDepthBook gets the current depth book info of a specific product
func (oc orderClient) QueryDepthBook(product string) (depthBook types.BookRes, err error) {
	jsonBytes, err := oc.GetCodec().MarshalJSON(orderkeeper.NewQueryDepthBookParams(product, orderkeeper.DefaultBookSize))
	if err != nil {
		return depthBook, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", ordertypes.QuerierRoute, ordertypes.QueryDepthBook)
	res, _, err := oc.Query(path, jsonBytes)
	if err != nil {
		return depthBook, utils.ErrClientQuery(err.Error())
	}

	if err = oc.GetCodec().UnmarshalJSON(res, &depthBook); err != nil {
		return depthBook, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

// QueryOrderDetail gets the detail info of an order by its order ID
func (oc orderClient) QueryOrderDetail(orderID string) (orderDetail types.OrderDetail, err error) {
	if err = params.CheckQueryOrderDetailParams(orderID); err != nil {
		return
	}

	path := fmt.Sprintf("custom/%s/%s/%s", ordertypes.QuerierRoute, ordertypes.QueryOrderDetail, orderID)
	res, _, err := oc.Query(path, nil)
	if err != nil {
		return orderDetail, utils.ErrClientQuery(err.Error())
	}

	if err = oc.GetCodec().UnmarshalJSON(res, &orderDetail); err != nil {
		return orderDetail, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
