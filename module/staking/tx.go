package staking

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/transact_params"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/tx"
	"github.com/okex/okchain-go-sdk/utils"
)

// Delegate delegates okt for voting
func (sc stakingClient) Delegate(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	coin, err := utils.ParseDecCoin(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := NewMsgDelegate(fromInfo.GetAddress(), coin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)
}

// Unbond unbonds the delegation on okchain
func (sc stakingClient) Unbond(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	coin, err := utils.ParseDecCoin(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := NewMsgUndelegate(fromInfo.GetAddress(), coin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)
}

// Vote votes to the some specific validators
func (sc stakingClient) Vote(fromInfo keys.Info, passWd string, valAddrsStr []string, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckVoteParams(fromInfo, passWd, valAddrsStr); err != nil {
		return types.TxResponse{}, err
	}

	valAddrs, err := utils.ParseValAddresses(valAddrsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : validator address parsed error: %s", err.Error())
	}

	msg := NewMsgVote(fromInfo.GetAddress(), valAddrs)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)
}

// DestroyValidator deregisters the validator and unbond the min-self-delegation
func (sc stakingClient) DestroyValidator(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := NewMsgDestroyValidator(fromInfo.GetAddress())

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)
}

// CreateValidator creates a new validator
func (sc stakingClient) CreateValidator(fromInfo keys.Info, passWd, pubkeyStr, moniker, identity, website, details, minSelfDelegationStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	pubkey, err := types.GetConsPubKeyBech32(pubkeyStr)
	if err != nil {
		return types.TxResponse{}, err
	}

	description := NewDescription(moniker, identity, website, details)

	minSelfDelegationCoin, err := utils.ParseDecCoin(minSelfDelegationStr)
	if err != nil {
		return types.TxResponse{}, err
	}

	msg := NewMsgCreateValidator(types.ValAddress(fromInfo.GetAddress()), pubkey, description, minSelfDelegationCoin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)
}

// EditValidator edits the description on a validator by the owner
func (sc stakingClient) EditValidator(fromInfo keys.Info, passWd, moniker, identity, website, details, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	description := NewDescription(moniker, identity, website, details)

	msg := NewMsgEditValidator(types.ValAddress(fromInfo.GetAddress()), description)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)

}

// RegisterProxy registers the identity of proxy
func (sc stakingClient) RegisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := NewMsgRegProxy(fromInfo.GetAddress(), true)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)

}

// UnregisterProxy registers the identity of proxy
func (sc stakingClient) UnregisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := NewMsgRegProxy(fromInfo.GetAddress(), false)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)

}

// BindProxy binds the staking tokens to a proxy
func (sc stakingClient) BindProxy(fromInfo keys.Info, passWd, proxyAddrStr, memo string, accNum, seqNum uint64) (
	types.TxResponse, error) {
	if err := transact_params.CheckSendParams(fromInfo, passWd, proxyAddrStr); err != nil {
		return types.TxResponse{}, err
	}

	proxyAddr, err := types.AccAddressFromBech32(proxyAddrStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Address [%s] error: %s", proxyAddrStr, err)
	}

	msg := NewMsgBindProxy(fromInfo.GetAddress(), proxyAddr)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)

}

// UnbindProxy unbinds the staking tokens from a proxy
func (sc stakingClient) UnbindProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (
	types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := NewMsgUnbindProxy(fromInfo.GetAddress())

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return sc.Broadcast(stdBytes, sc.GetConfig().BroadcastMode)

}
