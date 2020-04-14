package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	name   = "alice"
	passWd = "12345678"
	// sender's mnemonic
	mnemonic = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	addr     = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	// target address
	addr1 = "okchain1g7c3nvac7mjgn2m9mqllgat8wwd3aptdqket5k"
)

func TestGenerateUnsignedTransferOwnershipTx(t *testing.T) {
	err := GenerateUnsignedTransferOwnershipTx("btc-e68_okt", addr, addr1, "my memo", "./unsignedTx.json")
	require.NoError(t, err)
}

func TestMultiSign(t *testing.T) {
	fromInfo, _, err := CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	err = MultiSign(fromInfo, passWd, "./unsignedTx.json", "./signedTx.json")
	require.NoError(t, err)
}
