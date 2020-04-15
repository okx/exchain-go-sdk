package tx

import (
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
)

var (
	Kb keys.Keybase
)

func init() {
	Kb = keys.NewInMemory()
}

// BuildUnsignedStdTxOffline builds a stdTx without signature
func BuildUnsignedStdTxOffline(msgs []types.Msg, memo string) types.StdTx {
	Fee := types.NewStdFee(0, nil)
	return types.NewStdTx(msgs, Fee, nil, memo)
}

// BuildTx builds a common std tx with all the params
func BuildTx(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) (types.StdTx, error) {
	signMsg := StdSignMsg{
		ChainID:       "okchain",
		AccountNumber: accNumber,
		Sequence:      seqNumber,
		Memo:          memo,
		Msgs:          msgs,
		Fee:           types.NewStdFee(0, nil),
	}

	sig, err := makeSignature(Kb, fromName, passphrase, signMsg)
	if err != nil {
		return types.StdTx{}, err
	}

	return types.NewStdTx(signMsg.Msgs, signMsg.Fee, []types.StdSignature{sig}, signMsg.Memo), nil
}
