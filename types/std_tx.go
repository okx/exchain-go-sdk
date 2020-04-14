package types

import (
	"encoding/json"
	"github.com/tendermint/tendermint/crypto"
)

type StdTx struct {
	Msgs       []Msg          `json:"msg"`
	Fee        StdFee         `json:"-"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
}

func NewStdTx(msgs []Msg, fee StdFee, sigs []StdSignature, memo string) StdTx {
	return StdTx{
		Msgs:       msgs,
		Fee:        fee,
		Signatures: sigs,
		Memo:       memo,
	}
}

type StdFee struct {
	Amount Coins  `json:"amount"`
	Gas    uint64 `json:"gas"`
}

func NewStdFee(gas uint64, amount Coins) StdFee {
	return StdFee{
		Amount: amount,
		Gas:    gas,
	}
}

func (fee StdFee) Bytes() []byte {
	if len(fee.Amount) == 0 {
		fee.Amount = NewCoins()
	}
	bz, err := MsgCdc.MarshalJSON(fee) // TODO
	if err != nil {
		panic(err)
	}
	return bz
}

// StdSignature is the struct of signature in stdTx
type StdSignature struct {
	crypto.PubKey `json:"pub_key"`
	Signature     []byte `json:"signature"`
}

// NewStdSignature creates a new instance of std signature
func NewStdSignature(pubkey crypto.PubKey, signature []byte) StdSignature {
	return StdSignature{
		PubKey:    pubkey,
		Signature: signature,
	}
}

type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number"`
	ChainID       string            `json:"chain_id"`
	Fee           json.RawMessage   `json:"-"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Sequence      uint64            `json:"sequence"`
}
