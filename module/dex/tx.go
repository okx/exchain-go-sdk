package dex

import (
	"errors"
	"github.com/okex/okchain-go-sdk/module/dex/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
	"github.com/okex/okchain-go-sdk/types/params"
	"github.com/okex/okchain-go-sdk/utils"
)

// List lists a trading pair on dex
func (dc dexClient) List(fromInfo keys.Info, passWd, baseAsset, quoteAsset, initPriceStr, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckDexAssetsParams(fromInfo, passWd, baseAsset, quoteAsset); err != nil {
		return
	}

	initPrice := sdk.MustNewDecFromStr(initPriceStr)
	msg := types.NewMsgList(fromInfo.GetAddress(), baseAsset, quoteAsset, initPrice)

	return dc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// Deposit deposits some tokens to a specific product
func (dc dexClient) Deposit(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckProductParams(fromInfo, passWd, product); err != nil {
		return
	}

	amount, err := sdk.ParseDecCoin(amountStr)
	if err != nil {
		return
	}
	msg := types.NewMsgDeposit(fromInfo.GetAddress(), product, amount)

	return dc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// Withdraw withdraws some tokens from a specific product
func (dc dexClient) Withdraw(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckProductParams(fromInfo, passWd, product); err != nil {
		return
	}

	amount, err := sdk.ParseDecCoin(amountStr)
	if err != nil {
		return
	}
	msg := types.NewMsgWithdraw(fromInfo.GetAddress(), product, amount)

	return dc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// TransferOwnership signs the multi-signed tx from a json file and broadcast
func (dc dexClient) TransferOwnership(fromInfo keys.Info, passWd, inputPath string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	stdTx, err := utils.GetStdTxFromFile(dc.GetCodec(), inputPath)
	if err != nil {
		return
	}

	if len(stdTx.Msgs) == 0 {
		return resp, errors.New("failed. invalid msg type")
	}

	msg, ok := stdTx.Msgs[0].(types.MsgTransferOwnership)
	if !ok {
		return resp, errors.New("failed. invalid msg type")
	}

	return dc.BuildAndBroadcast(fromInfo.GetName(), passWd, stdTx.Memo, []sdk.Msg{msg}, accNum, seqNum)

}
