package tx

import (
	"fmt"
	"github.com/ok-chain/ok-gosdk/crypto/encoding/codec"
	"github.com/ok-chain/ok-gosdk/types"
	"github.com/ok-chain/ok-gosdk/types/msg"
	"github.com/ok-chain/ok-gosdk/utils"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var (
	MsgCdc = codec.New()
)

func init() {
	RegisterMsgCdc(MsgCdc)
}

func buildTx(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) (StdTx, error) {
	signMsg := StdSignMsg{
		ChainID:       "okchain",
		AccountNumber: accNumber,
		Sequence:      seqNumber,
		Memo:          memo,
		Msgs:          msgs,
		Fee:           NewStdFee(0, nil),
	}

	sig, err := makeSignature(utils.Kb, fromName, passphrase, signMsg)
	if err != nil {
		return StdTx{}, err
	}

	return NewStdTx(signMsg.Msgs, signMsg.Fee, []StdSignature{sig}, signMsg.Memo), nil
}

func BuildAndSignAndEncodeStdTx(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) ([]byte, error) {
	stdTx, err := buildTx(fromName, passphrase, memo, msgs, accNumber, seqNumber)
	if err != nil {
		return nil, fmt.Errorf("build stdTx error: %s", err)
	}

	// amino encoded
	txBytes, err := MsgCdc.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return nil, fmt.Errorf("amino encoded stdTx error: %s", err)
	}
	return txBytes, nil
}

func RegisterMsgCdc(cdc *amino.Codec) {
	//cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName, nil)

	cdc.RegisterInterface((*types.Msg)(nil), nil)
	cdc.RegisterConcrete(msg.MsgSend{}, "token/Send", nil)
	cdc.RegisterConcrete(msg.MsgNewOrder{}, "order/new", nil)
	cdc.RegisterConcrete(msg.MsgCancelOrder{}, "order/cancel", nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "auth/StdTx", nil)
}