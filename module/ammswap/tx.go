package ammswap

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/ammswap/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	ammswaptypes "github.com/okex/okexchain/x/ammswap/types"
	"time"
)

// AddLiquidity adds the number of liquidity of a token pair
func (pc ammswapClient) AddLiquidity(fromInfo keys.Info, passWd, minLiquidity, maxBaseAmount, quoteAmount, deadlineDuration,
	memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	minLiquidityDec, err := sdk.NewDecFromStr(minLiquidity)
	if err != nil {
		return
	}

	maxBaseAmountDecCoin, err := sdk.ParseDecCoin(maxBaseAmount)
	if err != nil {
		return
	}

	quoteAmountDecCoin, err := sdk.ParseDecCoin(quoteAmount)
	if err != nil {
		return
	}

	duration, err := time.ParseDuration(deadlineDuration)
	if err != nil {
		return
	}
	deadline := time.Now().Add(duration).Unix()

	msg := ammswaptypes.NewMsgAddLiquidity(minLiquidityDec, maxBaseAmountDecCoin, quoteAmountDecCoin, deadline, fromInfo.GetAddress())
	return pc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// RemoveLiquidity removes the number of liquidity of a token pair
func (pc ammswapClient) RemoveLiquidity(fromInfo keys.Info, passWd, liquidity, minBaseAmount, minQuoteAmount, deadlineDuration,
	memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	liquidityDec, err := sdk.NewDecFromStr(liquidity)
	if err != nil {
		return
	}

	minBaseAmountDecCoin, err := sdk.ParseDecCoin(minBaseAmount)
	if err != nil {
		return
	}

	minQuoteAmountDecCoin, err := sdk.ParseDecCoin(minQuoteAmount)
	if err != nil {
		return
	}

	duration, err := time.ParseDuration(deadlineDuration)
	if err != nil {
		return
	}
	deadline := time.Now().Add(duration).Unix()

	msg := ammswaptypes.NewMsgRemoveLiquidity(liquidityDec, minBaseAmountDecCoin, minQuoteAmountDecCoin, deadline, fromInfo.GetAddress())
	return pc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// CreateExchange creates a token pair in swap module
func (pc ammswapClient) CreateExchange(fromInfo keys.Info, passWd, baseToken, quoteToken, memo string,
	accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := types.NewMsgCreateExchange(baseToken, quoteToken, fromInfo.GetAddress())
	return pc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// TokenSwap swaps the number of specific token with another type token
func (pc ammswapClient) TokenSwap(fromInfo keys.Info, passWd, soldTokenAmount, minBoughtTokenAmount, recipient, deadlineDuration, memo string,
	accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	var soldTokenDecCoin, minBoughtTokenDecCoin sdk.DecCoin
	if soldTokenDecCoin, err = sdk.ParseDecCoin(soldTokenAmount); err != nil {
		return
	}
	if minBoughtTokenDecCoin, err = sdk.ParseDecCoin(minBoughtTokenAmount); err != nil {
		return
	}

	var duration time.Duration
	if duration, err = time.ParseDuration(deadlineDuration); err != nil {
		return
	}
	deadline := time.Now().Add(duration).Unix()

	var recip sdk.AccAddress
	if recipient == "" {
		recip = fromInfo.GetAddress()
	} else {
		if recip, err = sdk.AccAddressFromBech32(recipient); err != nil {
			return
		}
	}

	msg := types.NewMsgTokenToNativeToken(soldTokenDecCoin, minBoughtTokenDecCoin, deadline, recip, fromInfo.GetAddress())
	return pc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}
