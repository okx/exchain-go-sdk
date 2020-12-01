package farm

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	farmtypes "github.com/okex/okexchain/x/farm/types"
)

// CreatePool creates a farm pool
func (fc farmClient) CreatePool(fromInfo keys.Info, passWd, poolName, minLockAmountStr, yieldToken, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckCreatePoolParams(fromInfo, passWd, poolName, minLockAmountStr, yieldToken); err != nil {
		return
	}

	minLockAmount, err := sdk.ParseDecCoin(minLockAmountStr)
	if err != nil {
		return
	}

	msg := farmtypes.NewMsgCreatePool(fromInfo.GetAddress(), poolName, minLockAmount, yieldToken)
	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// DestroyPool destroys a farm pool
func (fc farmClient) DestroyPool(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckPoolNameParams(fromInfo, passWd, poolName); err != nil {
		return
	}

	msg := farmtypes.NewMsgDestroyPool(fromInfo.GetAddress(), poolName)
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

	msg := farmtypes.NewMsgProvide(poolName, fromInfo.GetAddress(), amount, amountYieldPerBlock, startHeightToYield)
	return fc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}
