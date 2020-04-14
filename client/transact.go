package client

import (
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/common/transact_params"
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

// Send transfers coins to others
func (cli *OKChainClient) Send(fromInfo keys.Info, passWd, toAddr, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckSendParams(fromInfo, passWd, toAddr); err != nil {
		return types.TxResponse{}, err
	}

	to, err := types.AccAddressFromBech32(toAddr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Address [%s] error: %s", toAddr, err)
	}

	coins, err := utils.ParseDecCoins(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse DecCoins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgTokenSend(fromInfo.GetAddress(), to, coins)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// order module

// NewOrders places orders with some detail info
func (cli *OKChainClient) NewOrders(fromInfo keys.Info, passWd, products, sides, prices, quantities, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	productStrs := strings.Split(products, ",")
	sideStrs := strings.Split(sides, ",")
	priceStrs := strings.Split(prices, ",")
	quantityStrs := strings.Split(quantities, ",")
	if err := transact_params.CheckNewOrderParams(fromInfo, passWd, productStrs, sideStrs, priceStrs, quantityStrs);
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
	if err := transact_params.CheckCancelOrderParams(fromInfo, passWd, orderIdStrs); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgCancelOrders(fromInfo.GetAddress(), orderIdStrs)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// staking module

// Delegate delegates okt for voting
func (cli *OKChainClient) Delegate(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	coin, err := utils.ParseDecCoin(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgDelegate(fromInfo.GetAddress(), coin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// Unbond unbonds the delegation on okchain
func (cli *OKChainClient) Unbond(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	coin, err := utils.ParseDecCoin(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgUndelegate(fromInfo.GetAddress(), coin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// Vote votes to the some specific validators
func (cli *OKChainClient) Vote(fromInfo keys.Info, passWd string, valAddrsStr []string, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckVoteParams(fromInfo, passWd, valAddrsStr); err != nil {
		return types.TxResponse{}, err
	}

	valAddrs, err := utils.ParseValAddresses(valAddrsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : validator address parsed error: %s", err.Error())
	}

	msg := types.NewMsgVote(fromInfo.GetAddress(), valAddrs)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// DestroyValidator deregisters the validator and unbond the min-self-delegation
func (cli *OKChainClient) DestroyValidator(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgDestroyValidator(fromInfo.GetAddress())

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// Unjail unjails the own validator which was jailed by slashing module
func (cli *OKChainClient) Unjail(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgUnjail(types.ValAddress(fromInfo.GetAddress()))

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// CreateValidator creates a new validator
func (cli *OKChainClient) CreateValidator(fromInfo keys.Info, passWd, pubkeyStr, moniker, identity, website, details, minSelfDelegationStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	pubkey, err := types.GetConsPubKeyBech32(pubkeyStr)
	if err != nil {
		return types.TxResponse{}, err
	}

	description := types.NewDescription(moniker, identity, website, details)

	minSelfDelegationCoin, err := utils.ParseDecCoin(minSelfDelegationStr)
	if err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgCreateValidator(types.ValAddress(fromInfo.GetAddress()), pubkey, description, minSelfDelegationCoin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// EditValidator edits the description on a validator by the owner
func (cli *OKChainClient) EditValidator(fromInfo keys.Info, passWd, moniker, identity, website, details, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	description := types.NewDescription(moniker, identity, website, details)

	msg := types.NewMsgEditValidator(types.ValAddress(fromInfo.GetAddress()), description)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

// RegisterProxy registers the identity of proxy
func (cli *OKChainClient) RegisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgRegProxy(fromInfo.GetAddress(), true)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

// UnregisterProxy registers the identity of proxy
func (cli *OKChainClient) UnregisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgRegProxy(fromInfo.GetAddress(), false)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

// BindProxy binds the staking tokens to a proxy
func (cli *OKChainClient) BindProxy(fromInfo keys.Info, passWd, proxyAddrStr, memo string, accNum, seqNum uint64) (
	types.TxResponse, error) {
	if err := transact_params.CheckSendParams(fromInfo, passWd, proxyAddrStr); err != nil {
		return types.TxResponse{}, err
	}

	proxyAddr, err := types.AccAddressFromBech32(proxyAddrStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Address [%s] error: %s", proxyAddrStr, err)
	}

	msg := types.NewMsgBindProxy(fromInfo.GetAddress(), proxyAddr)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

// UnbindProxy unbinds the staking tokens from a proxy
func (cli *OKChainClient) UnbindProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (
	types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgUnbindProxy(fromInfo.GetAddress())

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

// token module

// MultiSend multi-sends coins to several receivers
func (cli *OKChainClient) MultiSend(fromInfo keys.Info, passWd string, transfers []types.TransferUnit, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckTransferUnitsParams(fromInfo, passWd, transfers); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgMultiSend(fromInfo.GetAddress(), transfers)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// Issue issues a kind of token
func (cli *OKChainClient) Issue(fromInfo keys.Info, passWd, orgSymbol, wholeName, totalSupply, tokenDesc, memo string,
	mintable bool, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckTokenIssue(fromInfo, passWd, orgSymbol, wholeName, tokenDesc); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgTokenIssue(fromInfo.GetAddress(), tokenDesc, "", orgSymbol, wholeName, totalSupply, mintable)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// dex module

// List lists a trading pair on dex
func (cli *OKChainClient) List(fromInfo keys.Info, passWd, baseAsset, quoteAsset, initPriceStr, memo string, accNum,
	seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckDexAssets(fromInfo, passWd, baseAsset, quoteAsset); err != nil {
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
	if err := transact_params.CheckProduct(fromInfo, passWd, product); err != nil {
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
	if err := transact_params.CheckProduct(fromInfo, passWd, product); err != nil {
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
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
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
