package farm

import (
	"github.com/okex/okexchain-go-sdk/module/farm/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/crypto/keys"
	"github.com/okex/okexchain-go-sdk/types/params"
)

// CreatePool creates a farm pool
func (fc farmClient) CreatePool(fromInfo keys.Info, passWd, poolName, lockToken, yieldToken, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckCreatePoolParams(fromInfo, passWd, poolName, lockToken, yieldToken); err != nil {
		return
	}

	msg := types.NewMsgCreatePool(fromInfo.GetAddress(), poolName, lockToken, yieldToken)

	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// DestroyPool destroys a farm pool
func (fc farmClient) DestroyPool(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckPoolNameParams(fromInfo, passWd, poolName); err != nil {
		return
	}

	msg := types.NewMsgDestroyPool(fromInfo.GetAddress(), poolName)

	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// Provide provides a number of yield tokens into a pool
func (fc farmClient) Provide(fromInfo keys.Info, passWd, poolName, amountStr, yieldPerBlockStr string, startHeightToYield int64,
	memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckPoolNameParams(fromInfo, passWd, poolName); err != nil {
		return
	}

	amount, err := sdk.ParseDecCoin(amountStr)
	if err != nil {
		return
	}

	amountYieldPerBlock, err := sdk.NewDecFromStr(yieldPerBlockStr)
	if err != nil {
		return
	}

	msg := types.NewMsgProvide(fromInfo.GetAddress(), poolName, amount, amountYieldPerBlock, startHeightToYield)

	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// Lock locks a number of tokens for yield farming
func (fc farmClient) Lock(fromInfo keys.Info, passWd, poolName, amountStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckPoolNameParams(fromInfo, passWd, poolName); err != nil {
		return
	}

	amount, err := sdk.ParseDecCoin(amountStr)
	if err != nil {
		return
	}

	msg := types.NewMsgLock(fromInfo.GetAddress(), poolName, amount)

	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// Unlock unlocks a number of tokens from the farm pool and claims the current yield
func (fc farmClient) Unlock(fromInfo keys.Info, passWd, poolName, amountStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckPoolNameParams(fromInfo, passWd, poolName); err != nil {
		return
	}

	amount, err := sdk.ParseDecCoin(amountStr)
	if err != nil {
		return
	}

	msg := types.NewMsgUnlock(fromInfo.GetAddress(), poolName, amount)

	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// Claim claims yield farming rewards
func (fc farmClient) Claim(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckPoolNameParams(fromInfo, passWd, poolName); err != nil {
		return
	}

	msg := types.NewMsgClaim(fromInfo.GetAddress(), poolName)

	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// SetWhite is used for test
// TODO: remove it later
func (fc farmClient) SetWhite(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckPoolNameParams(fromInfo, passWd, poolName); err != nil {
		return
	}

	msg := types.NewMsgSetWhite(poolName, fromInfo.GetAddress())

	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}
