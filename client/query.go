package client

import (
	"errors"
	"fmt"
	"github.com/ok-chain/gosdk/common"
	"github.com/ok-chain/gosdk/common/queryParams"
	"github.com/ok-chain/gosdk/crypto/encoding/codec"
	"github.com/ok-chain/gosdk/types"
	"github.com/ok-chain/gosdk/utils"
)

const (
	accountInfoPath       = "/store/acc/key"
	accountTokensInfoPath = "custom/token/accounts/"
	tokensInfoPath        = "custom/token/tokens"
	tokenInfoPath         = "custom/token/info/"
	productsInfoPath      = "custom/token/products"
	depthbookInfoPath     = "custom/order/depthbook"
	candlesInfoPath       = "custom/backend/candles"
	tickersInfoPath       = "custom/backend/tickers"
	recentTxRecordPath    = "custom/backend/matches"
	openOrdersPath        = "custom/backend/orders/open"
	closedOrdersPath      = "custom/backend/orders/closed"
	dealsInfoPath         = "custom/backend/deals"
	transactionsInfoPath  = "custom/backend/txs"
)

func (cli *OKChainClient) GetAccountInfoByAddr(addr string) (types.Account, error) {
	accAddr, err := types.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errors.New("err : AccAddress converted from Bech32 Failed")
	}

	res, err := cli.query(accountInfoPath, utils.AddressStoreKey(accAddr))
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var account types.Account
	if err = cli.cdc.UnmarshalBinaryBare(res, &account); err != nil {
		return nil, fmt.Errorf("err : %s", err.Error())
	}

	return account, nil
}

func (cli *OKChainClient) GetTokensInfoByAddr(addr string) (types.AccountTokensInfo, error) {
	if !common.IsValidAccaddr(addr) {
		return types.AccountTokensInfo{}, fmt.Errorf("err : invalid account address")
	}

	accountParams := queryParams.NewQueryAccTokenParams("", "all")

	jsonBytes, err := cli.cdc.MarshalJSON(accountParams)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("error : AccTokenParam failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(accountTokensInfoPath+addr, jsonBytes)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var accTokensInfo types.AccountTokensInfo
	if err = cli.cdc.UnmarshalJSON(res, &accTokensInfo); err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("err : %s", err.Error())
	}
	return accTokensInfo, nil
}

func (cli *OKChainClient) GetTokenInfoByAddr(addr, symbol string) (types.AccountTokensInfo, error) {
	accountParams := queryParams.NewQueryAccTokenParams(symbol, "partial")

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

func (cli *OKChainClient) GetTokensInfo() ([]types.Token, error) {
	res, err := cli.query(tokensInfoPath, nil)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var tokensList []types.Token
	if err = cli.cdc.UnmarshalJSON(res, &tokensList); err != nil {
		return nil, fmt.Errorf("err : %s", err.Error())
	}

	return tokensList, nil
}

func (cli *OKChainClient) GetTokenInfo(symbol string) (types.Token, error) {
	res, err := cli.query(tokenInfoPath+symbol, nil)
	if err != nil {
		return types.Token{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var token types.Token
	if err = cli.cdc.UnmarshalJSON(res, &token); err != nil {
		return types.Token{}, fmt.Errorf("err : %s", err.Error())
	}

	return token, nil
}

func (cli *OKChainClient) GetProductsInfo() ([]types.TokenPair, error) {
	res, err := cli.query(productsInfoPath, nil)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var productsList []types.TokenPair
	if err = cli.cdc.UnmarshalJSON(res, &productsList); err != nil {
		return nil, fmt.Errorf("err : %s", err.Error())
	}

	return productsList, nil
}

func (cli *OKChainClient) GetDepthbookInfo(product string) (types.BookRes, error) {
	params := queryParams.NewQueryDepthBookParams(product, 200)
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
	params := queryParams.NewQueryKlinesParams(product, granularity, size)
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

func (cli *OKChainClient) GetTickersInfo(count int) (types.Tickers, error) {
	params := queryParams.NewQueryTickerParams("", count, true)
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
	params := queryParams.NewQueryMatchParams(product, int64(start), int64(end), page, perPage)
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
	// field hideNoFill fixed by false
	params := queryParams.NewQueryOrderListParams(addr, product, side, page, perPage, int64(start), int64(end), false)
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
	// field hideNoFill fixed by false
	params := queryParams.NewQueryOrderListParams(addr, product, side, page, perPage, int64(start), int64(end), false)
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
	params := queryParams.NewQueryDealsParams(addr, product, int64(start), int64(end), page, perPage, side)
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
	params := queryParams.NewQueryTxListParams(addr, int64(type_), int64(start), int64(end), page, perPage)
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
