package backend

import (
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/module/backend/types"
	"github.com/okex/okchain-go-sdk/utils"
)

// QueryCandles gets the candles data of a specific product
func (bc backendClient) QueryCandles(product string, granularity, size int) (candles [][]string, err error) {
	klinesParams := params.NewQueryKlinesParams(product, granularity, size)
	jsonBytes, err := bc.GetCodec().MarshalJSON(klinesParams)
	if err != nil {
		return candles, utils.ErrMarshalJSON(err.Error())
	}

	res, err := bc.Query(types.CandlesPath, jsonBytes)
	if err != nil {
		return candles, utils.ErrClientQuery(err.Error())
	}

	if err = utils.GetDataFromBaseResponse(res, &candles); err != nil {
		return candles, utils.ErrFilterDataFromBaseResponse("candles", err.Error())
	}

	return
}

// QueryTickers gets all tickers' data
func (bc backendClient) QueryTickers(count ...int) (tickers []types.Ticker, err error) {
	countNum, err := params.CheckQueryTickersParams(count)
	if err != nil {
		return
	}

	tickersParams := params.NewQueryTickerParams("", countNum, true)
	jsonBytes, err := bc.GetCodec().MarshalJSON(tickersParams)
	if err != nil {
		return tickers, utils.ErrMarshalJSON(err.Error())
	}

	res, err := bc.Query(types.TickersPath, jsonBytes)
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

	mathcParams := params.NewQueryMatchParams(product, int64(start), int64(end), page, perPageNum)
	jsonBytes, err := bc.GetCodec().MarshalJSON(mathcParams)
	if err != nil {
		return record, utils.ErrMarshalJSON(err.Error())
	}

	res, err := bc.Query(types.RecentTxRecordPath, jsonBytes)
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
	orderParams := params.NewQueryOrderListParams(addrStr, product, side, page, perPageNum, int64(start), int64(end), false)
	jsonBytes, err := bc.GetCodec().MarshalJSON(orderParams)
	if err != nil {
		return orders, utils.ErrMarshalJSON(err.Error())
	}

	res, err := bc.Query(types.OpenOrdersPath, jsonBytes)
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
	orderParams := params.NewQueryOrderListParams(addrStr, product, side, page, perPageNum, int64(start), int64(end), false)
	jsonBytes, err := bc.GetCodec().MarshalJSON(orderParams)
	if err != nil {
		return orders, utils.ErrMarshalJSON(err.Error())
	}

	res, err := bc.Query(types.ClosedOrdersPath, jsonBytes)
	if err != nil {
		return orders, utils.ErrClientQuery(err.Error())
	}

	if err = utils.UnmarshalListResponse(res, &orders); err != nil {
		return orders, utils.ErrFilterDataFromListResponse("closed orders", err.Error())
	}

	return
}
