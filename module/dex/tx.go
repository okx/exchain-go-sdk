package dex

import (
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/module/dex/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// List lists a trading pair on dex
func (dc dexClient) List(fromInfo keys.Info, passWd, baseAsset, quoteAsset, initPriceStr, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckDexAssets(fromInfo, passWd, baseAsset, quoteAsset); err != nil {
		return
	}

	initPrice := sdk.MustNewDecFromStr(initPriceStr)
	msg := types.NewMsgList(fromInfo.GetAddress(), baseAsset, quoteAsset, initPrice)

	return dc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}
