package dex

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/module/dex/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/tx"
	"io/ioutil"
)

// GenerateUnsignedTransferOwnershipTx generates the unsigned transfer-ownership transaction offline
func (dc dexClient) GenerateUnsignedTransferOwnershipTx(product, fromAddrStr, toAddrStr, memo, outputPath string) error {
	fromAddr, err := sdk.AccAddressFromBech32(fromAddrStr)
	if err != nil {
		return fmt.Errorf("failed. parse Address [%s] error: %s", fromAddrStr, err)
	}

	toAddr, err := sdk.AccAddressFromBech32(toAddrStr)
	if err != nil {
		return fmt.Errorf("failed. parse Address [%s] error: %s", toAddr, err)
	}

	msg := types.NewMsgTransferOwnership(fromAddr, toAddr, product)
	jsonBytes, err := dc.GetCodec().MarshalJSON(tx.BuildUnsignedStdTxOffline([]sdk.Msg{msg}, memo))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, jsonBytes, 0644)
}
