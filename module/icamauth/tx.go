package icamauth

import (
	"github.com/okex/exchain-go-sdk/types/params"
	ttx "github.com/okex/exchain-go-sdk/types/tx"
	"github.com/okex/exchain/libs/cosmos-sdk/client"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	codec_types "github.com/okex/exchain/libs/cosmos-sdk/codec/types"
	"github.com/okex/exchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	ibcmsg "github.com/okex/exchain/libs/cosmos-sdk/types/ibc-adapter"
	"github.com/okex/exchain/libs/cosmos-sdk/types/tx/signing"
	ibc_tx "github.com/okex/exchain/libs/cosmos-sdk/x/auth/ibc-tx"
	signing2 "github.com/okex/exchain/libs/cosmos-sdk/x/auth/ibcsigning"
	"github.com/okex/exchain/libs/cosmos-sdk/x/bank"
	tmcrypto "github.com/okex/exchain/libs/tendermint/crypto"
	icatypes "github.com/okex/exchain/x/icamauth/types"
)

var (
	txConfig client.TxConfig
)

func init() {
	txConfig = newTxConfig()
}

func newTxConfig() client.TxConfig {
	interfaceRegistry := codec_types.NewInterfaceRegistry()
	bank.RegisterInterface(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return ibc_tx.NewTxConfig(marshaler, ibc_tx.DefaultSignModes)
}

// AddLiquidity adds the number of liquidity of a token pair
func (ac icamauthClient) SubmitTx(fromInfo keys.Info, passWd, connectionID,
	memo string, data []byte, accNum, seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	var txMsg sdk.MsgProtoAdapter
	if err = ac.protoCdc.UnmarshalInterfaceJSON(data, &txMsg); err != nil {
		return
	}

	msg, err := icatypes.NewMsgSubmitTx(txMsg, connectionID, fromInfo.GetAddress().String())
	if err != nil {
		return
	}

	fees := ac.CovertToCoinAdapter()
	txb, err := ac.buildUnsignedTx(msg, memo, fees)
	if err != nil {
		return
	}

	pri, err := ttx.ExportPrivateKeyObject(fromInfo.GetName(), passWd)
	if err != nil {
		return
	}

	tx, err := ac.signedTx(txb, pri, accNum, seqNum)
	if err != nil {
		return
	}

	// get tx bytes
	txBytes, err := txConfig.TxEncoder()(tx)
	if err != nil {
		return
	}

	return ac.Broadcast(txBytes, ac.GetConfig().BroadcastMode)
}

func (ac icamauthClient) CovertToCoinAdapter() []sdk.CoinAdapter {
	//fees := make([]sdk.CoinAdapter, len(ac.GetConfig().Fees))
	//for i, fee := range ac.GetConfig().Fees {
	//	if fee.Denom == "okt" {
	//		weiFee := fee.Amount.MulInt64(1e18)
	//		fees[i] = sdk.NewCoinAdapter("wei", weiFee.RoundInt())
	//	} else {
	//		fees[i] = sdk.NewCoinAdapter("wei", sdk.NewInt(45000000000000))
	//	}
	//}
	//return fees
	fee := sdk.NewCoinAdapter("wei", sdk.NewInt(45000000000000))
	return []sdk.CoinAdapter{fee}
}

//func (ac icamauthClient) BuildAndBroadcast(fromName, passphrase, memo string, msgs []sdk.Msg, accNumber,
//	seqNumber uint64) (resp sdk.TxResponse, err error) {
//
//	utils.CompleteAndBroadcastTxCLI
//	txBytes, err = utils.PbTxBuildAndSign(cliCtx, txConfig, txBldr, keys.DefaultKeyPass, msgs)
//	if err != nil {
//		panic(err)
//	}
//	stdTx, err := bc.BuildStdTx(fromName, passphrase, memo, msgs, accNumber, seqNumber)
//	if err != nil {
//		return resp, fmt.Errorf("failed. build stdTx error: %s", err)
//	}
//
//	bytes, err := bc.cdc.MarshalBinaryLengthPrefixed(stdTx)
//	if err != nil {
//		return resp, fmt.Errorf("failed. encoded stdTx error: %s", err)
//	}
//
//	return bc.Broadcast(bytes, bc.GetConfig().BroadcastMode)
//}

func (ac icamauthClient) buildUnsignedTx(msg ibcmsg.Msg, memo string, fee sdk.CoinAdapters) (client.TxBuilder, error) {
	txb := txConfig.NewTxBuilder()

	// config txb
	txb.SetMemo(memo)
	txb.SetFeeAmount(fee)
	txb.SetGasLimit(ac.GetConfig().Gas)
	txb.SetTimeoutHeight(0)

	// set ibc msgs
	if err := txb.SetMsgs(msg); err != nil {
		return nil, err
	}

	return txb, nil
}

func (ac icamauthClient) signedTx(txb client.TxBuilder, priKey tmcrypto.PrivKey, accNum uint64, seqNum uint64) (signing2.Tx, error) {
	signMode := txConfig.SignModeHandler().DefaultMode()

	pubKey := ibc_tx.LagacyKey2PbKey(priKey.PubKey())
	// init signature
	signature := signing.SignatureV2{
		PubKey: pubKey,
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
		ChainID:       ac.GetConfig().ChainID,
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
		PubKey:   pubKey,
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
