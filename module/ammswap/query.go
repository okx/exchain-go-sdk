package ammswap

import (
	"fmt"

	"github.com/okex/okexchain-go-sdk/module/ammswap/types"
)

// QuerySwapTokenPair used for querying the exchange with token name
func (pc ammswapClient) QuerySwapTokenPair(token string) (types.SwapTokenPair, error) {
	var exchange types.SwapTokenPair

	res, err := pc.QueryStore(types.GetTokenPairKey(token), ModuleName, "key")
	if err != nil {
		return exchange, err
	}
	if len(res) == 0 {
		return exchange, fmt.Errorf("failed. no swapTokenPair found based on token %s", token)
	}

	pc.GetCodec().UnmarshalBinaryLengthPrefixed(res, &exchange)
	return exchange, nil
}