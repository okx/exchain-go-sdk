package slashing

import (
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
)

// Unjail unjails the own validator which was jailed by slashing module
func (sc slashingClient) Unjail(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (resp types.TxResponse,
	err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	msg := NewMsgUnjail(types.ValAddress(fromInfo.GetAddress()))

	return sc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	
}
