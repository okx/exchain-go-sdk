package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/crypto/encoding/codec"
	"github.com/okex/okchain-go-sdk/types"
)

const (
	depthbookInfoPath     = "custom/order/depthbook"
	candlesInfoPath       = "custom/backend/candles"
	tickersInfoPath       = "custom/backend/tickers"
	recentTxRecordPath    = "custom/backend/matches"
	openOrdersPath        = "custom/backend/orders/open"
	closedOrdersPath      = "custom/backend/orders/closed"
	dealsInfoPath         = "custom/backend/deals"
	transactionsInfoPath  = "custom/backend/txs"

)


func (cli *OKChainClient) GetTokenInfoByAddr(addr, symbol string) (types.AccountTokensInfo, error) {
	accountParams := params.NewQueryAccTokenParams(symbol, "partial")

	jsonBytes, err := cli.cdc.MarshalJSON(accountParams)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("error : AccTokenParam failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(accountTokensInfoPath+addr, jsonBytes)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var accTokenInfo types.AccountTokensInfo
	if err = cli.cdc.UnmarshalJSON(res, &accTokenInfo); err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("err : %s", err.Error())
	}
	return accTokenInfo, nil
}

func (cli *OKChainClient) GetDepthbookInfo(product string) (types.BookRes, error) {
	params := params.NewQueryDepthBookParams(product, 200)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return types.BookRes{}, fmt.Errorf("error : QueryDepthBookParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(depthbookInfoPath, jsonBytes)
	if err != nil {
		return types.BookRes{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var depthbook types.BookRes
	if err = cli.cdc.UnmarshalJSON(res, &depthbook); err != nil {
		return types.BookRes{}, fmt.Errorf("err : %s", err.Error())
	}

	return depthbook, nil
}

func (cli *OKChainClient) GetCandlesInfo(product string, granularity, size int) ([][]string, error) {
	params := params.NewQueryKlinesParams(product, granularity, size)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryKlinesParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(candlesInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var candles [][]string
	if err = codec.GetDataFromBaseResponse(res, &candles); err != nil {
		return nil, fmt.Errorf("candles unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return candles, nil
}

func (cli *OKChainClient) GetTickersInfo(count ...int) (types.Tickers, error) {
	countTmp, err := checkParamsGetTickersInfo(count)
	if err != nil {
		return nil, err
	}

	params := params.NewQueryTickerParams("", countTmp, true)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryTickerParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(tickersInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var tickers types.Tickers
	if err = codec.GetDataFromBaseResponse(res, &tickers); err != nil {
		return nil, fmt.Errorf("tickers unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return tickers, nil
}

func (cli *OKChainClient) GetRecentTxRecord(product string, start, end, page, perPage int) ([]types.MatchResult, error) {
	perPageTmp, err := checkParamsGetRecentTxRecord(product, start, end, page, perPage)
	if err != nil {
		return nil, err
	}

	params := params.NewQueryMatchParams(product, int64(start), int64(end), page, perPageTmp)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryMatchParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(recentTxRecordPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var records []types.MatchResult
	if err = codec.UnmarshalListResponse(res, &records); err != nil {
		return nil, fmt.Errorf("tx records unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return records, nil
}

func (cli *OKChainClient) GetOpenOrders(addr, product, side string, start, end, page, perPage int) ([]types.Order, error) {
	perPageTmp, err := checkParamsGetOpenClosedOrders(addr, product, side, start, end, page, perPage)
	if err != nil {
		return nil, err
	}

	// field hideNoFill fixed by false
	params := params.NewQueryOrderListParams(addr, product, side, page, perPageTmp, int64(start), int64(end), false)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryOrderListParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(openOrdersPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var openOrdersList []types.Order
	if err = codec.UnmarshalListResponse(res, &openOrdersList); err != nil {
		return nil, fmt.Errorf("open orders list unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return openOrdersList, nil

}

func (cli *OKChainClient) GetClosedOrders(addr, product, side string, start, end, page, perPage int) ([]types.Order, error) {
	perPageTmp, err := checkParamsGetOpenClosedOrders(addr, product, side, start, end, page, perPage)
	if err != nil {
		return nil, err
	}

	// field hideNoFill fixed by false
	params := params.NewQueryOrderListParams(addr, product, side, page, perPageTmp, int64(start), int64(end), false)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryOrderListParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(closedOrdersPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}
	var closedOrdersList []types.Order
	if err = codec.UnmarshalListResponse(res, &closedOrdersList); err != nil {
		return nil, fmt.Errorf("closed orders list unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return closedOrdersList, nil

}

func (cli *OKChainClient) GetDealsInfo(addr, product, side string, start, end, page, perPage int) ([]types.Deal, error) {
	perPageTmp, err := checkParamsGetDealsInfo(addr, product, side, start, end, page, perPage)
	if err != nil {
		return nil, err
	}

	params := params.NewQueryDealsParams(addr, product, int64(start), int64(end), page, perPageTmp, side)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryDealsParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(dealsInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var dealsInfo []types.Deal

	if err = codec.UnmarshalListResponse(res, &dealsInfo); err != nil {
		return nil, fmt.Errorf("deals Info list unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return dealsInfo, nil
}

func (cli *OKChainClient) GetTransactionsInfo(addr string, type_, start, end, page, perPage int) ([]types.Transaction, error) {
	perPageTmp, err := checkParamsGetTransactionsInfo(addr, type_, start, end, page, perPage)
	if err != nil {
		return nil, err
	}

	params := params.NewQueryTxListParams(addr, int64(type_), int64(start), int64(end), page, perPageTmp)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryTxListParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(transactionsInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var transactionsInfo []types.Transaction
	if err = codec.UnmarshalListResponse(res, &transactionsInfo); err != nil {
		return nil, fmt.Errorf("transactions Info list unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return transactionsInfo, nil
}


// dex module

