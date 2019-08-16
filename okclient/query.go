package okclient

import (
	"errors"
	"fmt"
	"github.com/ok-chain/ok-gosdk/common/queryParams"
	"github.com/ok-chain/ok-gosdk/types"
	"github.com/ok-chain/ok-gosdk/utils"
)

const (
	accountPath       = "/store/acc/key"
	accountTokensPath = "custom/token/accounts/"
	tokensPath        = "custom/token/tokens"
	tokenPath         = "custom/token/info/"
)

func (okCli *OKClient) GetAccountInfoByAddr(addr string) (types.Account, error) {
	accAddr, err := types.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errors.New("err : AccAddress converted from Bech32 Failed")
	}

	res, err := okCli.query(accountPath, utils.AddressStoreKey(accAddr))
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
	accountParam := queryParams.AccTokenParam{
		Symbol: "",
		Show:   "all",
	}

	jsonBytes, err := okCli.cdc.MarshalJSON(accountParam)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("error : AccTokenParam failed in json marshal : %s", err.Error())
	}

	res, err := okCli.query(accountTokensPath+addr, jsonBytes)
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
	accountParam := queryParams.AccTokenParam{
		Symbol: symbol,
		Show:   "partial",
	}

	jsonBytes, err := okCli.cdc.MarshalJSON(accountParam)
	if err != nil {
		return types.AccountTokensInfo{}, fmt.Errorf("error : AccTokenParam failed in json marshal : %s", err.Error())
	}

	res, err := okCli.query(accountTokensPath+addr, jsonBytes)
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
	res, err := okCli.query(tokensPath, nil)
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
	res, err := okCli.query(tokenPath+symbol, nil)
	if err != nil {
		return types.Token{}, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var token types.Token
	if err = okCli.cdc.UnmarshalJSON(res, &token); err != nil {
		return types.Token{}, fmt.Errorf("err : %s", err.Error())
	}

	return token, nil
}
