package poolswap

import (
	"fmt"

	"github.com/okex/okchain-go-sdk/module/poolswap/types"
)

func (pc poolswapClient) QueryExchange(token string) (types.SwapTokenPair, error) {
	exchange := types.SwapTokenPair{}

	res, err := pc.QueryStore(types.GetTokenPairKey(token), ModuleName, "key")
	if err != nil {
		return exchange, err
	}
	if len(res) == 0 {
		return exchange, fmt.Errorf("failed. no swapTokenPair found based on token %s", token)
	}

	pc.GetCodec().MustUnmarshalJSON(res, exchange)
	return exchange, nil
}
