package token

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/module/token/types"
	"github.com/okex/okchain-go-sdk/types/params"
	"github.com/okex/okchain-go-sdk/utils"
)

// QueryAccountTokenInfo gets a specific available token info of an account
func (tc tokenClient) QueryAccountTokenInfo(addrStr, symbol string) (accTokensInfo types.AccountTokensInfo, err error) {
	if err = params.IsValidAccAddr(addrStr); err != nil {
		return
	}

	accountParams := params.NewQueryAccTokenParams(symbol, "partial")

	jsonBytes, err := tc.GetCodec().MarshalJSON(accountParams)
	if err != nil {
		return accTokensInfo, utils.ErrMarshalJSON(err.Error())
	}

	res, err := tc.Query(fmt.Sprintf("%s%s", types.AccountTokensInfoPath, addrStr), jsonBytes)
	if err != nil {
		return accTokensInfo, utils.ErrClientQuery(err.Error())
	}

	if err = tc.GetCodec().UnmarshalJSON(res, &accTokensInfo); err != nil {
		return accTokensInfo, utils.ErrUnmarshalJSON(err.Error())
	}
	return
}

// QueryAccountTokensInfo gets all the available tokens info of an account
func (tc tokenClient) QueryAccountTokensInfo(addrStr string) (accTokensInfo types.AccountTokensInfo, err error) {
	if err = params.IsValidAccAddr(addrStr); err != nil {
		return
	}

	accountParams := params.NewQueryAccTokenParams("", "all")

	jsonBytes, err := tc.GetCodec().MarshalJSON(accountParams)
	if err != nil {
		return accTokensInfo, utils.ErrMarshalJSON(err.Error())
	}

	res, err := tc.Query(fmt.Sprintf("%s%s", types.AccountTokensInfoPath, addrStr), jsonBytes)
	if err != nil {
		return accTokensInfo, utils.ErrClientQuery(err.Error())
	}

	if err = tc.GetCodec().UnmarshalJSON(res, &accTokensInfo); err != nil {
		return accTokensInfo, utils.ErrUnmarshalJSON(err.Error())
	}
	return
}

// QueryTokenInfo gets token info with a specific symbol or the owner address
func (tc tokenClient) QueryTokenInfo(ownerAddr, symbol string) (tokens []types.Token, err error) {
	if err = params.CheckQueryTokenInfoParams(ownerAddr, symbol); err != nil {
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
