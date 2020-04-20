package dex

import (
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/module/dex/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/tx"
	"github.com/okex/okchain-go-sdk/utils"
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
	jsonBytes, err := dc.GetCodec().MarshalJSON(dc.BuildUnsignedStdTxOffline([]sdk.Msg{msg}, memo))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, jsonBytes, 0644)
}

// MultiSign appends signature to the unsigned tx file of transfer-ownership
func (dc dexClient) MultiSign(fromInfo keys.Info, passWd, inputPath, outputPath string) error {
	stdTx, err := utils.GetStdTxFromFile(dc.GetCodec(),inputPath)
	if err != nil {
		return err
	}

	if len(stdTx.Msgs) == 0 {
		return errors.New("failed. msg is empty")
	}

	msg, ok := stdTx.Msgs[0].(types.MsgTransferOwnership)
	if !ok {
		return errors.New("failed. invalid msg type")
	}

	signature, _, err := tx.Kb.Sign(fromInfo.GetName(), passWd, msg.GetSignBytes())
	if err != nil {
		return fmt.Errorf("failed. sign error: %s", err.Error())
	}

	msg.ToSignature = sdk.NewStdSignature(fromInfo.GetPubKey(), signature)
	jsonBytes, err := dc.GetCodec().MarshalJSON(dc.BuildUnsignedStdTxOffline([]sdk.Msg{msg}, stdTx.Memo))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, jsonBytes, 0644)
}
