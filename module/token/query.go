package token

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/module/token/types"
)

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
