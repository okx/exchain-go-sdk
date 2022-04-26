package ibc

import (
	"fmt"
	"github.com/okex/exchain-go-sdk/module/auth"
	ibcmsg "github.com/okex/exchain/libs/cosmos-sdk/types/ibc-adapter"
	"math"

	"github.com/okex/exchain-go-sdk/module"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/client"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	codec_types "github.com/okex/exchain/libs/cosmos-sdk/codec/types"
	cryptotypes "github.com/okex/exchain/libs/cosmos-sdk/crypto/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	signing "github.com/okex/exchain/libs/cosmos-sdk/types/tx/signing"
	ibc_tx "github.com/okex/exchain/libs/cosmos-sdk/x/auth/ibc-tx"
	signing2 "github.com/okex/exchain/libs/cosmos-sdk/x/auth/ibcsigning"
	ibc_type "github.com/okex/exchain/libs/ibc-go/modules/apps/transfer/types"
	client_types "github.com/okex/exchain/libs/ibc-go/modules/core/02-client/types"
)

const (
	src_port = "transfer"
)

var (
	txConfig client.TxConfig
)

func init() {
	txConfig = newTxConfig()
}

func newTxConfig() client.TxConfig {
	interfaceRegistry := codec_types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return ibc_tx.NewTxConfig(marshaler, ibc_tx.DefaultSignModes)
}

func (ibc ibcClient) GetLatestHeight() uint64 {
	status, err := ibc.Status()
	if err != nil {

		fmt.Println(err)
		return 0
	}

	return uint64(status.SyncInfo.LatestBlockHeight)
}

func (ibc ibcClient) Transfer(priKey cryptotypes.PrivKey, srcChannel string, receiver string, amount string, fee sdk.CoinAdapters, memo string, targetRpc string) (resp sdk.TxResponse, err error) {

	// get account info
	accountInfo, err := auth.NewAuthClient(ibc.BaseClient).QueryAccount(priKey.PubKey().Address().String())
	if err != nil {

		return sdk.TxResponse{}, err
	}

	coin, err := sdk.ParseCoinNormalized(amount)
	if err != nil {

		return sdk.TxResponse{}, err
	}

	timeoutHeight, err := getTimeoutHeight(targetRpc)
	if err != nil {

		return sdk.TxResponse{}, err
	}

	// generate msg
	msg := &ibc_type.MsgTransfer{
		SourcePort:       src_port,
		SourceChannel:    srcChannel,
		Token:            sdk.NewCoinAdapter("wei", coin.Amount.Mul(sdk.NewDec(int64(math.Pow10(18)))).RoundInt()),
		Sender:           priKey.PubKey().Address().String(),
		Receiver:         receiver,
		TimeoutHeight:    timeoutHeight,
		TimeoutTimestamp: 0,
	}

	// build unsignedTx
	txb, err := ibc.buildUnsignedTx(msg, memo, fee)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// sign Tx
	tx, err := ibc.signedTx(txb, priKey, accountInfo.GetAccountNumber(), accountInfo.GetSequence())
	if err != nil {

		return sdk.TxResponse{}, err
	}

	// get tx bytes
	txBytes, err := txConfig.TxEncoder()(tx)
	if err != nil {

		return sdk.TxResponse{}, err
	}

	return ibc.Broadcast(txBytes, "sync")
}

// get Client Height from destination chain
func getTimeoutHeight(dstRpc string) (client_types.Height, error) {
	cdc := gosdktypes.NewCodec()

	dstClient := module.NewBaseClient(cdc, &gosdktypes.ClientConfig{
		NodeURI: dstRpc,
	})

	status, err := dstClient.Status()
	if err != nil {

		return client_types.Height{}, err
	}

	latestHeight := status.SyncInfo.LatestBlockHeight

	return client_types.Height{
		RevisionNumber: client_types.ParseChainID(status.NodeInfo.Network),
		RevisionHeight: uint64(latestHeight + 1000),
	}, nil
}

func (ibc ibcClient) buildUnsignedTx(msg ibcmsg.Msg, memo string, fee sdk.CoinAdapters) (client.TxBuilder, error) {
	txb := txConfig.NewTxBuilder()

	// config txb
	txb.SetMemo(memo)
	txb.SetFeeAmount(fee)
	txb.SetGasLimit(ibc.GetConfig().Gas)
	txb.SetTimeoutHeight(0)

	// set ibc msgs
	if err := txb.SetMsgs(msg); err != nil {

		return nil, err
	}

	return txb, nil
}

func (ibc ibcClient) signedTx(txb client.TxBuilder, priKey cryptotypes.PrivKey, accNum uint64, seqNum uint64) (signing2.Tx, error) {
	signMode := txConfig.SignModeHandler().DefaultMode()

	// init signature
	signature := signing.SignatureV2{
		PubKey: priKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode: signMode,
		},
		Sequence: seqNum,
	}

	err := txb.SetSignatures(signature)
	if err != nil {

		return nil, err
	}

	signerData := signing2.SignerData{
		ChainID:       ibc.GetConfig().ChainID,
		AccountNumber: accNum,
		Sequence:      seqNum,
	}

	// bytes to sign
	signBytes, err := txConfig.SignModeHandler().GetSignBytes(signMode, signerData, txb.GetTx())
	if err != nil {
		panic(err)
	}

	// signed bytes
	sigBytes, err := priKey.Sign(signBytes)
	if err != nil {
		panic(err)
	}

	sigData := signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: sigBytes,
	}

	sig := signing.SignatureV2{
		PubKey:   priKey.PubKey(),
		Data:     &sigData,
		Sequence: seqNum,
	}

	// set signature
	err = txb.SetSignatures(sig)
	if err != nil {

		return nil, err
	}

	return txb.GetTx(), nil
}
