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
