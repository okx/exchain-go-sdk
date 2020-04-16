package token

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/module/token/types"
)

// QueryAccountTokenInfo gets a specific available token info of an account
func (tc tokenClient) QueryAccountTokenInfo(addrStr, symbol string) (accTokensInfo types.AccountTokensInfo, err error) {
	if !params.IsValidAccAddr(addrStr) {
		return accTokensInfo, fmt.Errorf("failed. invalid account address")
	}

	accountParams := params.NewQueryAccTokenParams(symbol, "partial")

	jsonBytes, err := tc.GetCodec().MarshalJSON(accountParams)
	if err != nil {
		return accTokensInfo, fmt.Errorf("failed. AccTokenParam failed in JSON marshal : %s", err.Error())
	}

	res, err := tc.Query(fmt.Sprintf("%s%s", types.AccountTokensInfoPath, addrStr), jsonBytes)
	if err != nil {
		return accTokensInfo, fmt.Errorf("failed. ok client query error : %s", err.Error())
	}

	if err = tc.GetCodec().UnmarshalJSON(res, &accTokensInfo); err != nil {
		return accTokensInfo, fmt.Errorf("failed. unmarshal JSON error : %s", err.Error())
	}
	return
}

// QueryAccountTokensInfo gets all the available tokens info of an account
func (tc tokenClient) QueryAccountTokensInfo(addrStr string) (accTokensInfo types.AccountTokensInfo, err error) {
	if !params.IsValidAccAddr(addrStr) {
		return accTokensInfo, fmt.Errorf("failed. invalid account address")
	}

	accountParams := params.NewQueryAccTokenParams("", "all")

	jsonBytes, err := tc.GetCodec().MarshalJSON(accountParams)
	if err != nil {
		return accTokensInfo, fmt.Errorf("failded. AccTokenParam failed in json marshal : %s", err.Error())
	}

	res, err := tc.Query(fmt.Sprintf("%s%s", types.AccountTokensInfoPath, addrStr), jsonBytes)
	if err != nil {
		return accTokensInfo, fmt.Errorf("failed. ok client query error : %s", err.Error())
	}

	if err = tc.GetCodec().UnmarshalJSON(res, &accTokensInfo); err != nil {
		return accTokensInfo, fmt.Errorf("failed. unmarshal JSON error : %s", err.Error())
	}
	return
}

// QueryTokenInfo gets token info with a specific symbol or the owner address
func (tc tokenClient) QueryTokenInfo(ownerAddr, symbol string) (tokens []types.Token, err error) {
	if err = params.CheckQueryTokenInfo(ownerAddr, symbol); err != nil {
		return
	}

	if len(symbol) != 0 {
		res, err := tc.Query(fmt.Sprintf("custom/%s/info/%s", types.ModuleName, symbol), nil)
		if err != nil {
			return tokens, fmt.Errorf("failed. token %s doesn't exist", symbol)
		}

		var token types.Token
		tc.GetCodec().MustUnmarshalJSON(res, &token)
		tokens = append(tokens, token)
		return tokens, err
	}

	res, err := tc.Query(fmt.Sprintf("custom/%s/tokens/%s", types.ModuleName, ownerAddr), nil)
	if err != nil {
		return tokens, fmt.Errorf("failed. %s doesn't own any tokens: %s", ownerAddr, err.Error())
	}

	tc.GetCodec().MustUnmarshalJSON(res, &tokens)
	return

}
