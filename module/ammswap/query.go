package ammswap

import (
	"fmt"
	"math/big"

	"github.com/okex/okexchain-go-sdk/module/ammswap/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
)

var (
	DefaultBondDenom = "okt"

	FeeRate = sdk.NewDecWithPrec(3,3)
)

// QuerySwapTokenPair used for querying one swap token pair
func (pc ammswapClient) QuerySwapTokenPair(token string) (types.SwapTokenPair, error) {
	var exchange types.SwapTokenPair

	res, err := pc.QueryStore(types.GetTokenPairKey(token), ModuleName, "key")
	if err != nil {
		return exchange, err
	}
	if len(res) == 0 {
		return exchange, fmt.Errorf("failed. no swapTokenPair found based on token %s", token)
	}

	err = pc.GetCodec().UnmarshalBinaryLengthPrefixed(res, &exchange)
	if err != nil {
		return exchange, err
	}
	return exchange, nil
}

// QuerySwapTokenPairs used for querying the all the swap token pairs
func (pc ammswapClient) QuerySwapTokenPairs() ([]types.SwapTokenPair, error) {
	var exchanges []types.SwapTokenPair

	resKVs, err := pc.QuerySubspace(types.TokenPairPrefixKey, ModuleName)
	if err != nil {
		return nil, err
	}
	for _, kv := range resKVs {
		var exchange types.SwapTokenPair
		pc.GetCodec().MustUnmarshalBinaryLengthPrefixed(kv.Value, &exchange)
		exchanges = append(exchanges, exchange)
	}
	return exchanges, nil
}

// QueryBuyAmount used for querying how much token would get from a pool
func (pc ammswapClient) QueryBuyAmount(soldToken sdk.DecCoin, tokenToBuy string) (sdk.Dec, error) {
	var buyAmount sdk.Dec

	swapTokenPair := types.GetSwapTokenPairName(soldToken.Denom, tokenToBuy)
	tokenPair, errTokenPair := pc.QuerySwapTokenPair(swapTokenPair)
	if errTokenPair == nil {
		buyAmount = calculateTokenToBuy(tokenPair, soldToken, tokenToBuy, FeeRate).Amount
	}else {
		tokenPairName1 := types.GetSwapTokenPairName(soldToken.Denom, DefaultBondDenom)
		tokenPair1, err := pc.QuerySwapTokenPair(tokenPairName1)
		if err != nil {
			return buyAmount, err
		}

		tokenPairName2 := types.GetSwapTokenPairName(tokenToBuy, DefaultBondDenom)
		tokenPair2, err := pc.QuerySwapTokenPair(tokenPairName2)
		if err != nil {
			return buyAmount, err
		}

		nativeToken := calculateTokenToBuy(tokenPair1, soldToken, DefaultBondDenom, FeeRate)
		buyAmount = calculateTokenToBuy(tokenPair2, nativeToken, tokenToBuy, FeeRate).Amount
	}

	return buyAmount, nil
}

//calculateTokenToBuy calculates the amount to buy
func calculateTokenToBuy(swapTokenPair types.SwapTokenPair, sellToken sdk.DecCoin, buyTokenDenom string, feeRate sdk.Dec) sdk.DecCoin {
	var inputReserve, outputReserve sdk.Dec
	if buyTokenDenom < sellToken.Denom {
		inputReserve = swapTokenPair.QuotePooledCoin.Amount
		outputReserve = swapTokenPair.BasePooledCoin.Amount
	} else {
		inputReserve = swapTokenPair.BasePooledCoin.Amount
		outputReserve = swapTokenPair.QuotePooledCoin.Amount
	}
	tokenBuyAmt := getInputPrice(sellToken.Amount, inputReserve, outputReserve, feeRate)
	tokenBuy := sdk.NewDecCoinFromDec(buyTokenDenom, tokenBuyAmt)

	return tokenBuy
}

// getInputPrice gets the input price
func getInputPrice(inputAmount, inputReserve, outputReserve, feeRate sdk.Dec) sdk.Dec {
	inputAmountWithFee := inputAmount.MulTruncate(sdk.OneDec().Sub(feeRate).MulTruncate(sdk.NewDec(1000)))
	denominator := inputReserve.MulTruncate(sdk.NewDec(1000)).Add(inputAmountWithFee)
	return mulAndQuo(inputAmountWithFee, outputReserve, denominator)
}

// mulAndQuo returns a * b / c
func mulAndQuo(a, b, c sdk.Dec) sdk.Dec {
	// 10^8
	auxiliaryDec := sdk.NewDecFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(sdk.Precision), nil))
	a = a.MulTruncate(auxiliaryDec)
	return a.MulTruncate(b).QuoTruncate(c).QuoTruncate(auxiliaryDec)
}