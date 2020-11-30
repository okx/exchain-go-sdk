package order

import (
	"fmt"

	"github.com/okex/okexchain-go-sdk/module/order/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/okex/okexchain-go-sdk/utils"
)

// QueryDepthBook gets the current depth book info of a specific product
func (oc orderClient) QueryDepthBook(product string) (depthBook types.BookRes, err error) {
	depthBookParams := params.NewQueryDepthBookParams(product, 200)
	jsonBytes, err := oc.GetCodec().MarshalJSON(depthBookParams)
	if err != nil {
		return depthBook, utils.ErrMarshalJSON(err.Error())
	}

	res, _, err := oc.Query(types.DepthbookPath, jsonBytes)
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

	res, _, err := oc.Query(fmt.Sprintf("%s/%s", types.OrderDetailPath, orderID), nil)
	if err != nil {
		return orderDetail, utils.ErrClientQuery(err.Error())
	}

	if err = oc.GetCodec().UnmarshalJSON(res, &orderDetail); err != nil {
		return orderDetail, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
