package ammswap

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/ammswap/types"
	"github.com/okex/okexchain-go-sdk/utils"
	ammswaptypes "github.com/okex/okexchain/x/ammswap/types"
	"math/big"
)

var (
	DefaultBondDenom = "okt"

	FeeRate = sdk.NewDecWithPrec(3, 3)
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
func (pc ammswapClient) QueryBuyAmount(soldToken sdk.DecCoin, tokenToBuy string) (sdk.Dec, error) {
	var buyAmount sdk.Dec

	swapTokenPair := types.GetSwapTokenPairName(soldToken.Denom, tokenToBuy)
	tokenPair, errTokenPair := pc.QuerySwapTokenPair(swapTokenPair)
	if errTokenPair == nil {
		buyAmount = calculateTokenToBuy(tokenPair, soldToken, tokenToBuy, FeeRate).Amount
	} else {
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
