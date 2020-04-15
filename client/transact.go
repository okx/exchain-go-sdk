package client

import (
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/tx"
	"github.com/okex/okchain-go-sdk/utils"
	"strings"
)

// broadcast mode
const (
	BroadcastBlock = "block"
	BroadcastSync  = "sync"
	BroadcastAsync = "async"
)


// order module

// NewOrders places orders with some detail info
func (cli *OKChainClient) NewOrders(fromInfo keys.Info, passWd, products, sides, prices, quantities, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	productStrs := strings.Split(products, ",")
	sideStrs := strings.Split(sides, ",")
	priceStrs := strings.Split(prices, ",")
	quantityStrs := strings.Split(quantities, ",")
	if err := params.CheckNewOrderParams(fromInfo, passWd, productStrs, sideStrs, priceStrs, quantityStrs);
		err != nil {
		return types.TxResponse{}, err
	}

	orderItems := types.BuildOrderItems(productStrs, sideStrs, priceStrs, quantityStrs)
	msg := types.NewMsgNewOrders(fromInfo.GetAddress(), orderItems)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

// CancelOrders cancels orders by orderIds
func (cli *OKChainClient) CancelOrders(fromInfo keys.Info, passWd, orderIds, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	orderIdStrs := strings.Split(orderIds, ",")
	if err := params.CheckCancelOrderParams(fromInfo, passWd, orderIdStrs); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgCancelOrders(fromInfo.GetAddress(), orderIdStrs)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// token module


// dex module

// List lists a trading pair on dex
func (cli *OKChainClient) List(fromInfo keys.Info, passWd, baseAsset, quoteAsset, initPriceStr, memo string, accNum,
	seqNum uint64) (types.TxResponse, error) {
	if err := params.CheckDexAssets(fromInfo, passWd, baseAsset, quoteAsset); err != nil {
		return types.TxResponse{}, err
	}

	initPrice := types.MustNewDecFromStr(initPriceStr)
	msg := types.NewMsgList(fromInfo.GetAddress(), baseAsset, quoteAsset, initPrice)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// Deposit deposits some tokens to a specific product
func (cli *OKChainClient) Deposit(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (
	types.TxResponse, error) {
	if err := params.CheckProduct(fromInfo, passWd, product); err != nil {
		return types.TxResponse{}, err
	}

	amount, err := utils.ParseDecCoin(amountStr)
	if err != nil {
		return types.TxResponse{}, err
	}
	msg := types.NewMsgDeposit(fromInfo.GetAddress(), product, amount)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// Withdraw withdraws some tokens from a specific product
func (cli *OKChainClient) Withdraw(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (
	types.TxResponse, error) {
	if err := params.CheckProduct(fromInfo, passWd, product); err != nil {
		return types.TxResponse{}, err
	}

	amount, err := utils.ParseDecCoin(amountStr)
	if err != nil {
		return types.TxResponse{}, err
	}
	msg := types.NewMsgWithdraw(fromInfo.GetAddress(), product, amount)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// TransferOwnership signs the multi-signed tx from a json file and broadcast
func (cli *OKChainClient) TransferOwnership(fromInfo keys.Info, passWd, inputPath string, accNum, seqNum uint64) (
	types.TxResponse, error) {
	if err := params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	stdTx, err := utils.GetStdTxFromFile(inputPath)
	if err != nil {
		return types.TxResponse{}, err
	}

	if len(stdTx.Msgs) == 0 {
		return types.TxResponse{}, errors.New("failed. invalid msg type")
	}

	msg, ok := stdTx.Msgs[0].(types.MsgTransferOwnership)
	if !ok {
		return types.TxResponse{}, errors.New("failed. invalid msg type")
	}

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, stdTx.Memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}
