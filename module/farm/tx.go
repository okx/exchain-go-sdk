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
