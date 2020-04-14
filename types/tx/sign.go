package tx

import (
	"encoding/json"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
)

type StdSignMsg struct {
	ChainID       string       `json:"chain_id"`
	AccountNumber uint64       `json:"account_number"`
	Sequence      uint64       `json:"sequence"`
	Fee           types.StdFee `json:"fee"`
	Msgs          []types.Msg  `json:"msgs"`
	Memo          string       `json:"memo"`
}

func (msg StdSignMsg) Bytes() []byte {
	return StdSignBytes(msg.ChainID, msg.AccountNumber, msg.Sequence, msg.Fee, msg.Msgs, msg.Memo)
}

func StdSignBytes(chainID string, accnum uint64, sequence uint64, fee types.StdFee, msgs []types.Msg, memo string) []byte {
	var msgsBytes []json.RawMessage
	for _, msg := range msgs {
		msgsBytes = append(msgsBytes, json.RawMessage(msg.GetSignBytes()))
	}
	bz, err := types.MsgCdc.MarshalJSON(types.StdSignDoc{
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
	return types.MustSortJSON(bz)
}

func makeSignature(keybase keys.Keybase, name, passphrase string,
	msg StdSignMsg) (sig types.StdSignature, err error) {
	sigBytes, pubkey, err := keybase.Sign(name, passphrase, msg.Bytes())
	if err != nil {
		return
	}
	return types.StdSignature{
		PubKey:    pubkey,
		Signature: sigBytes,
	}, nil
}
