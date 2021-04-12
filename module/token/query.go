package token

import (
	"fmt"

	"github.com/okex/exchain-go-sdk/module/token/types"
	"github.com/okex/exchain-go-sdk/types/params"
	"github.com/okex/exchain/x/token"
)

// QueryTokenInfo gets token info with a specific symbol or the owner address
func (tc tokenClient) QueryTokenInfo(ownerAddr, symbol string) (tokens []types.TokenResp, err error) {
	if err = params.CheckQueryTokenInfoParams(ownerAddr, symbol); err != nil {
		return
	}

	if len(symbol) != 0 {
		path := fmt.Sprintf("custom/%s/info/%s", token.QuerierRoute, symbol)
		res, _, err := tc.Query(path, nil)
		if err != nil {
			return tokens, fmt.Errorf("failed. token %s doesn't exist", symbol)
		}

		var tokenResp types.TokenResp
		tc.GetCodec().MustUnmarshalJSON(res, &tokenResp)
		tokens = append(tokens, tokenResp)
		return tokens, err
	}

	path := fmt.Sprintf("custom/%s/tokens/%s", token.QuerierRoute, ownerAddr)
	res, _, err := tc.Query(path, nil)
	if err != nil {
		return tokens, fmt.Errorf("failed. %s doesn't own any tokens: %s", ownerAddr, err.Error())
	}

	tc.GetCodec().MustUnmarshalJSON(res, &tokens)
	return
}
