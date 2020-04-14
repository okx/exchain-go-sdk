package utils

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/libs/pkg/errors"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/tx"
	"io/ioutil"
)

// GenerateUnsignedTransferOwnershipTx generates the unsigned transfer-ownership transaction offline
func GenerateUnsignedTransferOwnershipTx(product, fromAddrStr, toAddrStr, memo, outputPath string) error {
	fromAddr, err := types.AccAddressFromBech32(fromAddrStr)
	if err != nil {
		return fmt.Errorf("err : parse Address [%s] error: %s", fromAddrStr, err)
	}

	toAddr, err := types.AccAddressFromBech32(toAddrStr)
	if err != nil {
		return fmt.Errorf("err : parse Address [%s] error: %s", toAddr, err)
	}

	msg := types.NewMsgTransferOwnership(fromAddr, toAddr, product)
	jsonBytes, err := types.MsgCdc.MarshalJSON(tx.BuildUnsignedStdTxOffline([]types.Msg{msg}, memo))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, jsonBytes, 0644)
}

// MultiSign appends signature to the unsigned tx file of transfer-ownership
func MultiSign(fromInfo keys.Info, passWd, inputPath, outputPath string) error {
	stdTx, err := GetStdTxFromFile(inputPath)
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

	msg.ToSignature = types.NewStdSignature(fromInfo.GetPubKey(), signature)
	jsonBytes, err := types.MsgCdc.MarshalJSON(tx.BuildUnsignedStdTxOffline([]types.Msg{msg}, stdTx.Memo))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, jsonBytes, 0644)
}
