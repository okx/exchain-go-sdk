package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/common"
	"github.com/okex/okchain-go-sdk/types"
)

const (
	countDefault = 100
)

func checkParamsGetTickersInfo(count []int) (countRet int, err error) {
	if len(count) > 1 {
		return 0, errors.New("invalid params input for 'GetTickersInfo'")
	}

	if len(count) == 0 {
		countRet = countDefault
	} else {
		if count[0] < 0 {
			return 0, errors.New("'count' cannot be negative")
		}
		countRet = count[0]
	}
	return
}

func checkParamsGetRecentTxRecord(product string, start, end, page, perPage int) (perPageRet int, err error) {
	if product == "" {
		return 0, errors.New("'product' is empty")
	}

	perPageRet, err = common.CheckParamsPaging(start, end, page, perPage)
	return
}

func checkParamsGetOpenClosedOrders(addr, product, side string, start, end, page, perPage int) (perPageRet int, err error) {
	if !common.IsValidAccaddr(addr) {
		return 0, errors.New("invalid address input")
	}

	if product == "" {
		return 0, errors.New("'product' is empty")
	}

	if !common.IsValidSide(side) {
		return 0, errors.New("'side' can only be 'BUY' or 'SELL'")

	}

	perPageRet, err = common.CheckParamsPaging(start, end, page, perPage)
	return

}

func checkParamsGetDealsInfo(addr, product, side string, start, end, page, perPage int) (perPageRet int, err error) {
	return checkParamsGetOpenClosedOrders(addr, product, side, start, end, page, perPage)
}

func checkParamsGetTransactionsInfo(addr string, type_, start, end, page, perPage int) (perPageRet int, err error) {
	if !common.IsValidAccaddr(addr) {
		return 0, errors.New("invalid address input")
	}

	if type_ < 0 {
		return 0, errors.New("'type_' cannot be negative")

	}

	perPageRet, err = common.CheckParamsPaging(start, end, page, perPage)
	return
}

func getOrderIdsFromResponse(txResp *types.TxResponse) (orderIds []string) {
	for _, event := range txResp.Events {
		if event.Type == "message" {
			for _, attribute := range event.Attributes {
				if attribute.Key == "orders" {
					var orderRes []types.OrderResult
					if err := json.Unmarshal([]byte(attribute.Value), &orderRes); err != nil {
						fmt.Println(err)
						return
					}

					for _, res := range orderRes {
						orderIds = append(orderIds, res.OrderID)
					}
				}
			}
		}
	}

	return
}
