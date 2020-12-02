package backend

import (
	"fmt"
	"github.com/okex/okexchain-go-sdk/module/backend/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/okex/okexchain-go-sdk/utils"
	backendtypes "github.com/okex/okexchain/x/backend/types"
)

// QueryCandles gets the candles data of a specific product
func (bc backendClient) QueryCandles(product string, granularity, size int) (candles [][]string, err error) {
	jsonBytes, err := bc.GetCodec().MarshalJSON(backendtypes.NewQueryKlinesParams(product, granularity, size))
	if err != nil {
		return candles, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", backendtypes.QuerierRoute, backendtypes.QueryCandleList)
	res, _, err := bc.Query(path, jsonBytes)
	if err != nil {
		return candles, utils.ErrClientQuery(err.Error())
	}

	if err = utils.GetDataFromBaseResponse(res, &candles); err != nil {
		return candles, utils.ErrFilterDataFromBaseResponse("candles", err.Error())
	}

	return
}

// QueryTickers gets all tickers' data
// NOTE: all products are involved with setting "" to product
func (bc backendClient) QueryTickers(product string, count ...int) (tickers []types.Ticker, err error) {
	countNum, err := params.CheckQueryTickersParams(count)
	if err != nil {
		return
	}

	queryParams := backendtypes.QueryTickerParams{
		Product: product,
		Count:   countNum,
		Sort:    true,
	}

	jsonBytes, err := bc.GetCodec().MarshalJSON(queryParams)
	if err != nil {
		return tickers, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", backendtypes.QuerierRoute, backendtypes.QueryTickerList)
	res, _, err := bc.Query(path, jsonBytes)
	if err != nil {
		return tickers, utils.ErrClientQuery(err.Error())
	}

	if err = utils.GetDataFromBaseResponse(res, &tickers); err != nil {
		return tickers, utils.ErrFilterDataFromBaseResponse("tickers", err.Error())
	}

	return
}

// QueryRecentTxRecord gets the specific product's record of recent transactions
func (bc backendClient) QueryRecentTxRecord(product string, start, end, page, perPage int) (record []types.MatchResult,
	err error) {
	perPageNum, err := params.CheckQueryRecentTxRecordParams(product, start, end, page, perPage)
	if err != nil {
		return
	}

	jsonBytes, err := bc.GetCodec().MarshalJSON(backendtypes.NewQueryMatchParams(product, int64(start), int64(end), page, perPageNum))
	if err != nil {
		return record, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", backendtypes.QuerierRoute, backendtypes.QueryMatchResults)
	res, _, err := bc.Query(path, jsonBytes)
	if err != nil {
		return record, utils.ErrClientQuery(err.Error())
	}

	if err = utils.UnmarshalListResponse(res, &record); err != nil {
		return record, utils.ErrFilterDataFromListResponse("recent tx record", err.Error())
	}

	return
}

// QueryOpenOrders gets the open orders of a specific account
func (bc backendClient) QueryOpenOrders(addrStr, product, side string, start, end, page, perPage int) (orders []types.Order,
	err error) {
	perPageNum, err := params.CheckQueryOrdersParams(addrStr, product, side, start, end, page, perPage)
	if err != nil {
		return
	}

	// field hideNoFill fixed by false
	ordersParams := params.NewQueryOrderListParams(addrStr, product, side, page, perPageNum, int64(start), int64(end), false)
	jsonBytes, err := bc.GetCodec().MarshalJSON(ordersParams)
	if err != nil {
		return orders, utils.ErrMarshalJSON(err.Error())
	}

	res, _, err := bc.Query(types.OpenOrdersPath, jsonBytes)
	if err != nil {
		return orders, utils.ErrClientQuery(err.Error())
	}

	if err = utils.UnmarshalListResponse(res, &orders); err != nil {
		return orders, utils.ErrFilterDataFromListResponse("open orders", err.Error())
	}

	return
}

// QueryClosedOrders gets the closed orders of a specific account
func (bc backendClient) QueryClosedOrders(addrStr, product, side string, start, end, page, perPage int) (orders []types.Order,
	err error) {
	perPageNum, err := params.CheckQueryOrdersParams(addrStr, product, side, start, end, page, perPage)
	if err != nil {
		return
	}

	// field hideNoFill fixed by false
	ordersParams := params.NewQueryOrderListParams(addrStr, product, side, page, perPageNum, int64(start), int64(end), false)
	jsonBytes, err := bc.GetCodec().MarshalJSON(ordersParams)
	if err != nil {
		return orders, utils.ErrMarshalJSON(err.Error())
	}

	res, _, err := bc.Query(types.ClosedOrdersPath, jsonBytes)
	if err != nil {
		return orders, utils.ErrClientQuery(err.Error())
	}

	if err = utils.UnmarshalListResponse(res, &orders); err != nil {
		return orders, utils.ErrFilterDataFromListResponse("closed orders", err.Error())
	}

	return
}

// QueryDeals gets the deals info of a specific account
func (bc backendClient) QueryDeals(addrStr, product, side string, start, end, page, perPage int) (deals []types.Deal, err error) {
	perPageNum, err := params.CheckQueryOrdersParams(addrStr, product, side, start, end, page, perPage)
	if err != nil {
		return
	}

	dealsParams := params.NewQueryDealsParams(addrStr, product, int64(start), int64(end), page, perPageNum, side)
	jsonBytes, err := bc.GetCodec().MarshalJSON(dealsParams)
	if err != nil {
		return deals, utils.ErrMarshalJSON(err.Error())
	}

	res, _, err := bc.Query(types.DealsPath, jsonBytes)
	if err != nil {
		return deals, utils.ErrClientQuery(err.Error())
	}

	if err = utils.UnmarshalListResponse(res, &deals); err != nil {
		return deals, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

// QueryTransactions gets the transactions of a specific account
func (bc backendClient) QueryTransactions(addrStr string, typeCode, start, end, page, perPage int) (transactions []types.Transaction, err error) {
	perPageNum, err := params.CheckQueryTransactionsParams(addrStr, typeCode, start, end, page, perPage)
	if err != nil {
		return
	}

	transactionsParams := params.NewQueryTxListParams(addrStr, int64(typeCode), int64(start), int64(end), page, perPageNum)
	jsonBytes, err := bc.GetCodec().MarshalJSON(transactionsParams)
	if err != nil {
		return transactions, utils.ErrMarshalJSON(err.Error())
	}

	res, _, err := bc.Query(types.TransactionsPath, jsonBytes)
	if err != nil {
		return transactions, utils.ErrClientQuery(err.Error())
	}

	if err = utils.UnmarshalListResponse(res, &transactions); err != nil {
		return transactions, utils.ErrFilterDataFromListResponse("transactions", err.Error())
	}

	return
}
