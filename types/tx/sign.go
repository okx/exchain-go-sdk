package tx

import (
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
)

var (
	// Kb - global keybase
	Kb keys.Keybase
)

func init() {
	Kb = keys.NewInMemory()
}

// MakeSignature completes the signature
func MakeSignature(name, passphrase string, msg types.StdSignMsg) (sig types.StdSignature, err error) {
	sigBytes, pubkey, err := Kb.Sign(name, passphrase, msg.Bytes())
	if err != nil {
		return
	}
	return types.StdSignature{
		PubKey:    pubkey,
		Signature: sigBytes,
	}, nil
}
