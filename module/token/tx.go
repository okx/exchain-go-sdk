package token

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/module/token/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/params"
)

// Send transfers coins to other receiver
func (tc tokenClient) Send(fromInfo keys.Info, passWd, toAddrStr, coinsStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckSendParams(fromInfo, passWd, toAddrStr); err != nil {
		return
	}

	toAddr, err := sdk.AccAddressFromBech32(toAddrStr)
	if err != nil {
		return resp, fmt.Errorf("failed. parse Address [%s] error: %s", toAddrStr, err)
	}

	coins, err := sdk.ParseDecCoins(coinsStr)
	if err != nil {
		return resp, fmt.Errorf("failed. parse DecCoins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgTokenSend(fromInfo.GetAddress(), toAddr, coins)

	return tc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// MultiSend multi-sends coins to several receivers
func (tc tokenClient) MultiSend(fromInfo keys.Info, passWd string, transfers []types.TransferUnit, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckTransferUnitsParams(fromInfo, passWd, transfers); err != nil {
		return
	}

	msg := types.NewMsgMultiSend(fromInfo.GetAddress(), transfers)

	return tc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// Issue issues a kind of token
func (tc tokenClient) Issue(fromInfo keys.Info, passWd, orgSymbol, wholeName, totalSupply, tokenDesc, memo string,
	mintable bool, accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckTokenIssueParams(fromInfo, passWd, orgSymbol, wholeName, tokenDesc); err != nil {
		return
	}

	msg := types.NewMsgTokenIssue(fromInfo.GetAddress(), tokenDesc, "", orgSymbol, wholeName, totalSupply, mintable)

	return tc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}
