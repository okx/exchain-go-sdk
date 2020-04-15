package staking

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
)

// Delegate delegates okt for voting
func (sc stakingClient) Delegate(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	coin, err := utils.ParseDecCoin(coinsStr)
	if err != nil {
		return resp, fmt.Errorf("failed : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := NewMsgDelegate(fromInfo.GetAddress(), coin)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
}

// Unbond unbonds the delegation on okchain
func (sc stakingClient) Unbond(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	coin, err := utils.ParseDecCoin(coinsStr)
	if err != nil {
		return resp, fmt.Errorf("failed : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := NewMsgUndelegate(fromInfo.GetAddress(), coin)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}

// Vote votes to the some specific validators
func (sc stakingClient) Vote(fromInfo keys.Info, passWd string, valAddrsStr []string, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckVoteParams(fromInfo, passWd, valAddrsStr); err != nil {
		return
	}

	valAddrs, err := utils.ParseValAddresses(valAddrsStr)
	if err != nil {
		return resp, fmt.Errorf("failed. validator address parsed error: %s", err.Error())
	}

	msg := NewMsgVote(fromInfo.GetAddress(), valAddrs)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}

// DestroyValidator deregisters the validator and unbond the min-self-delegation
func (sc stakingClient) DestroyValidator(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := NewMsgDestroyValidator(fromInfo.GetAddress())

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
}

// CreateValidator creates a new validator
func (sc stakingClient) CreateValidator(fromInfo keys.Info, passWd, pubkeyStr, moniker, identity, website, details,
	memo string, accNum, seqNum uint64) (resp types.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	pubkey, err := types.GetConsPubKeyBech32(pubkeyStr)
	if err != nil {
		return
	}

	description := NewDescription(moniker, identity, website, details)
	msg := NewMsgCreateValidator(types.ValAddress(fromInfo.GetAddress()), pubkey, description)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}

// EditValidator edits the description on a validator by the owner
func (sc stakingClient) EditValidator(fromInfo keys.Info, passWd, moniker, identity, website, details, memo string, accNum,
	seqNum uint64) (resp types.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	description := NewDescription(moniker, identity, website, details)
	msg := NewMsgEditValidator(types.ValAddress(fromInfo.GetAddress()), description)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}

// RegisterProxy registers the identity of proxy
func (sc stakingClient) RegisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := NewMsgRegProxy(fromInfo.GetAddress(), true)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}

// UnregisterProxy registers the identity of proxy
func (sc stakingClient) UnregisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := NewMsgRegProxy(fromInfo.GetAddress(), false)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}

// BindProxy binds the staking tokens to a proxy
func (sc stakingClient) BindProxy(fromInfo keys.Info, passWd, proxyAddrStr, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckSendParams(fromInfo, passWd, proxyAddrStr); err != nil {
		return
	}

	proxyAddr, err := types.AccAddressFromBech32(proxyAddrStr)
	if err != nil {
		return resp, fmt.Errorf("failed. parse Address [%s] error: %s", proxyAddrStr, err)
	}

	msg := NewMsgBindProxy(fromInfo.GetAddress(), proxyAddr)

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}

// UnbindProxy unbinds the staking tokens from a proxy
func (sc stakingClient) UnbindProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := NewMsgUnbindProxy(fromInfo.GetAddress())

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}
