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
