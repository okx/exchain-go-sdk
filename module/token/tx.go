package token

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	types2 "github.com/okex/okchain-go-sdk/module/token/types"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
)

// Send transfers coins to other receiver
func (tc tokenClient) Send(fromInfo keys.Info, passWd, toAddrStr, coinsStr, memo string, accNum, seqNum uint64) (
	resp types.TxResponse, err error) {
	if err = params.CheckSendParams(fromInfo, passWd, toAddrStr); err != nil {
		return
	}

	toAddr, err := types.AccAddressFromBech32(toAddrStr)
	if err != nil {
		return resp, fmt.Errorf("failed. parse Address [%s] error: %s", toAddrStr, err)
	}

	coins, err := utils.ParseDecCoins(coinsStr)
	if err != nil {
		return resp, fmt.Errorf("failed. parse DecCoins [%s] error: %s", coinsStr, err)
	}

	msg := types2.NewMsgTokenSend(fromInfo.GetAddress(), toAddr, coins)

	return tc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)

}
