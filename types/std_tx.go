package types

import (
	"encoding/json"

	"github.com/tendermint/tendermint/crypto"
)

// const
const (
	BroadcastSync  BroadcastMode = "sync"
	BroadcastAsync BroadcastMode = "async"
	BroadcastBlock BroadcastMode = "commit"
)

// BroadcastMode defines different mode to broadcast
type BroadcastMode string

var (
	_ Tx = (*StdTx)(nil)
)

// StdTx is a standard way to wrap a Msg with Fee and Signatures
type StdTx struct {
	Msgs       []Msg          `json:"msg"`
	Fee        StdFee         `json:"fee"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
}

// NewStdTx creates a new instance of StdTx
func NewStdTx(msgs []Msg, fee StdFee, sigs []StdSignature, memo string) StdTx {
	return StdTx{
		Msgs:       msgs,
		Fee:        fee,
		Signatures: sigs,
		Memo:       memo,
	}
}

// nolint
func (st StdTx) GetMsgs() []Msg       { return nil }
func (st StdTx) ValidateBasic() Error { return nil }

// StdFee includes the amount of coins paid in fees and the maximum gas to be used by the transaction
type StdFee struct {
	Amount DecCoins `json:"amount"`
	Gas    uint64   `json:"gas"`
}

// NewStdFee creates a new instance of StdFee
func NewStdFee(gas uint64, amount DecCoins) StdFee {
	return StdFee{
		Amount: amount,
		Gas:    gas,
	}
}

// Bytes for signing later
func (sf StdFee) Bytes() []byte {
	if len(sf.Amount) == 0 {
		sf.Amount = NewDecCoins()
	}
	bz, err := Cdc.MarshalJSON(sf) // TODO
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

// StdSignDoc is replay-prevention structure
// It includes the result of msg.GetSignBytes(), as well as the ChainID (prevent cross chain replay) and the Sequence
// numbers for each signature (prevent inchain replay and enforce tx ordering per account)
type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number"`
	ChainID       string            `json:"chain_id"`
	Fee           json.RawMessage   `json:"fee"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Sequence      uint64            `json:"sequence"`
}

// stdSignBytes returns the bytes to sign for a transaction
func stdSignBytes(chainID string, accnum uint64, sequence uint64, fee StdFee, msgs []Msg, memo string) []byte {
	var msgsBytes []json.RawMessage
	for _, msg := range msgs {
		msgsBytes = append(msgsBytes, json.RawMessage(msg.GetSignBytes()))
	}
	bz, err := Cdc.MarshalJSON(StdSignDoc{
		AccountNumber: accnum,
		ChainID:       chainID,
		Fee:           json.RawMessage(fee.Bytes()),
		Memo:          memo,
		Msgs:          msgsBytes,
		Sequence:      sequence,
	})
	if err != nil {
		panic(err)
	}
	return MustSortJSON(bz)
}
