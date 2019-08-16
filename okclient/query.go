package okclient

import (
	"errors"
	"fmt"
	"github.com/ok-chain/ok-gosdk/common/queryParams"
	"github.com/ok-chain/ok-gosdk/crypto/encoding/codec"
	"github.com/ok-chain/ok-gosdk/types"
	"github.com/ok-chain/ok-gosdk/utils"
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
)

func (okCli *OKClient) GetAccountInfoByAddr(addr string) (types.Account, error) {
	accAddr, err := types.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errors.New("err : AccAddress converted from Bech32 Failed")
	}

	res, err := okCli.query(accountInfoPath, utils.AddressStoreKey(accAddr))
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var account types.Account
	if err = okCli.cdc.UnmarshalBinaryBare(res, &account); err != nil {
		return nil, fmt.Errorf("err : %s", err.Error())
	}

	return account, nil
}

func (okCli *OKClient) GetTokensInfoByAddr(addr string) (types.AccountTokensInfo, error) {
	accountParams := queryParams.NewQueryAccTokenParams("", "all")

	jsonBytes, err := okCli.cdc.MarshalJSON(accountParams)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("error : AccTokenParam failed in json marshal : %s", err.Error())
	}

	res, err := okCli.query(accountTokensInfoPath+addr, jsonBytes)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var accTokensInfo types.AccountTokensInfo
	if err = okCli.cdc.UnmarshalJSON(res, &accTokensInfo); err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("err : %s", err.Error())
	}
	return accTokensInfo, nil
}

func (okCli *OKClient) GetTokenInfoByAddr(addr, symbol string) (types.AccountTokensInfo, error) {
	accountParams := queryParams.NewQueryAccTokenParams(symbol, "partial")

	jsonBytes, err := okCli.cdc.MarshalJSON(accountParams)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("error : AccTokenParam failed in json marshal : %s", err.Error())
	}

	res, err := okCli.query(accountTokensInfoPath+addr, jsonBytes)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var accTokenInfo types.AccountTokensInfo
	if err = okCli.cdc.UnmarshalJSON(res, &accTokenInfo); err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("err : %s", err.Error())
	}
	return accTokenInfo, nil
}

func (okCli *OKClient) GetTokensInfo() ([]types.Token, error) {
	res, err := okCli.query(tokensInfoPath, nil)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var tokensList []types.Token
	if err = okCli.cdc.UnmarshalJSON(res, &tokensList); err != nil {
		return nil, fmt.Errorf("err : %s", err.Error())
	}

	return tokensList, nil
}

func (okCli *OKClient) GetTokenInfo(symbol string) (types.Token, error) {
	res, err := okCli.query(tokenInfoPath+symbol, nil)
	if err != nil {
		return types.Token{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var token types.Token
	if err = okCli.cdc.UnmarshalJSON(res, &token); err != nil {
		return types.Token{}, fmt.Errorf("err : %s", err.Error())
	}

	return token, nil
}

func (okCli *OKClient) GetProductsInfo() ([]types.TokenPair, error) {
	res, err := okCli.query(productsInfoPath, nil)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var productsList []types.TokenPair
	if err = okCli.cdc.UnmarshalJSON(res, &productsList); err != nil {
		return nil, fmt.Errorf("err : %s", err.Error())
	}

	return productsList, nil
}

func (okCli *OKClient) GetDepthbookInfo(product string) (types.BookRes, error) {
	params := queryParams.NewQueryDepthBookParams(product, 200)
	jsonBytes, err := okCli.cdc.MarshalJSON(params)
	if err != nil {
		return types.BookRes{}, fmt.Errorf("error : QueryDepthBookParams failed in json marshal : %s", err.Error())
	}

	res, err := okCli.query(depthbookInfoPath, jsonBytes)
	if err != nil {
		return types.BookRes{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var depthbook types.BookRes
	if err = okCli.cdc.UnmarshalJSON(res, &depthbook); err != nil {
		return types.BookRes{}, fmt.Errorf("err : %s", err.Error())
	}

	return depthbook, nil
}

func (okCli *OKClient) GetCandlesInfo(product string, granularity, size int) ([][]string, error) {
	params := queryParams.NewQueryKlinesParams(product, granularity, size)
	jsonBytes, err := okCli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryKlinesParams failed in json marshal : %s", err.Error())
	}

	res, err := okCli.query(candlesInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var candles [][]string
	if err = codec.UnmarshalBaseResponse(res, &candles); err != nil {
		return nil, fmt.Errorf("candles unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return candles, nil
}

func (okCli *OKClient) GetTickersInfo(count int) (types.Tickers, error) {
	params := queryParams.NewQueryTickerParams("", count, true)
	jsonBytes, err := okCli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryTickerParams failed in json marshal : %s", err.Error())
	}

	res, err := okCli.query(tickersInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var tickers types.Tickers
	if err = codec.UnmarshalBaseResponse(res, &tickers); err != nil {
		return nil, fmt.Errorf("tickers unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return tickers, nil
}
