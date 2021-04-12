package ammswap

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/exchain-go-sdk/module/ammswap/types"
	"github.com/okex/exchain-go-sdk/utils"
	ammswaptypes "github.com/okex/exchain/x/ammswap/types"
)

// QuerySwapTokenPair used for querying one swap token pair
func (ac ammswapClient) QuerySwapTokenPair(token string) (exchange types.SwapTokenPair, err error) {
	res, _, err := ac.QueryStore(ammswaptypes.GetTokenPairKey(token), ammswaptypes.StoreKey, "key")
	if err != nil {
		return
	}

	if len(res) == 0 {
		return exchange, fmt.Errorf("failed. no swapTokenPair found based on token %s", token)
	}

	err = ac.GetCodec().UnmarshalBinaryLengthPrefixed(res, &exchange)
	if err != nil {
		return exchange, err
	}

	return
}

// QuerySwapTokenPairs used for querying the all the swap token pairs
func (ac ammswapClient) QuerySwapTokenPairs() (exchanges []types.SwapTokenPair, err error) {
	path := fmt.Sprintf("custom/%s/%s", ammswaptypes.QuerierRoute, ammswaptypes.QuerySwapTokenPairs)
	res, _, err := ac.Query(path, nil)
	if err != nil {
		return exchanges, utils.ErrClientQuery(err.Error())
	}

	if err = ac.GetCodec().UnmarshalJSON(res, &exchanges); err != nil {
		return exchanges, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

// QueryBuyAmount used for querying how much token would get from a pool
func (ac ammswapClient) QueryBuyAmount(tokenToSellStr, tokenDenomToBuy string) (amount sdk.Dec, err error) {
	tokenToSell, err := sdk.ParseDecCoin(tokenToSellStr)
	if err != nil {
		return
	}

	queryParams := ammswaptypes.QueryBuyAmountParams{
		SoldToken:  tokenToSell,
		TokenToBuy: tokenDenomToBuy,
	}
	jsonBytes, err := ac.GetCodec().MarshalJSON(queryParams)
	if err != nil {
		return amount, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", ammswaptypes.QuerierRoute, ammswaptypes.QueryBuyAmount)
	res, _, err := ac.Query(path, jsonBytes)
	if err != nil {
		return amount, utils.ErrClientQuery(err.Error())
	}

	if err = ac.GetCodec().UnmarshalJSON(res, &amount); err != nil {
		return amount, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
