package tx

import (
	"github.com/okx/okbchain/app/crypto/hd"
	"github.com/okx/okbchain/libs/cosmos-sdk/crypto/keys"
	authtypes "github.com/okx/okbchain/libs/cosmos-sdk/x/auth/types"
)

var (
	// Kb - global keybase
	Kb keys.Keybase
)

func init() {
	Kb = keys.NewInMemory(hd.EthSecp256k1Options()...)
}

// MakeSignature completes the signature
func MakeSignature(name, passphrase string, msg authtypes.StdSignMsg) (sig authtypes.StdSignature, err error) {
	sigBytes, pubkey, err := Kb.Sign(name, passphrase, msg.Bytes())
	if err != nil {
		return
	}
	return authtypes.StdSignature{
		PubKey:    pubkey,
		Signature: sigBytes,
	}, nil
}
