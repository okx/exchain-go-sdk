package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	defaultMnemonic   = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	defaultAddr       = "okexchain1qj5c07sm6jetjz8f509qtrxgh4psxkv32x0qas"
	defaultAddrEth    = "0x04A987fa1Bd4b2B908e9A3Ca058cc8BD43035991"
	defaultPrivateKey = "89C81C304704E9890025A5A91898802294658D6E4034A11C6116F4B129EA12D3"
)

func TestCreateAccount(t *testing.T) {
	info, mnemo, err := CreateAccount("", "")
	require.NoError(t, err)
	require.Equal(t, defaultName, info.GetName())
	require.NotNil(t, mnemo)
}

func TestCreateAccountWithMnemo(t *testing.T) {
	info, mnemo, err := CreateAccountWithMnemo(defaultMnemonic, defaultName, defaultPassWd)
	require.NoError(t, err)
	require.Equal(t, defaultName, info.GetName())
	require.Equal(t, defaultMnemonic, mnemo)
	require.Equal(t, defaultAddr, info.GetAddress().String())

	_, _, err = CreateAccountWithMnemo(defaultMnemonic, "", defaultPassWd)
	require.NoError(t, err)

	_, _, err = CreateAccountWithMnemo(defaultMnemonic, defaultName, "")
	require.NoError(t, err)

	_, _, err = CreateAccountWithMnemo("", defaultName, defaultPassWd)
	require.Error(t, err)

	_, _, err = CreateAccountWithMnemo(defaultPassWd, defaultName, defaultPassWd)
	require.Error(t, err)

	invalidMnemo := fmt.Sprintf("%s abandon", defaultMnemonic)
	_, _, err = CreateAccountWithMnemo(invalidMnemo, defaultName, defaultPassWd)
	require.Error(t, err)
	fmt.Println(err)
}

func TestCreateAccountWithPrivateKey(t *testing.T) {
	infoByMnemo, _, err := CreateAccountWithMnemo(defaultMnemonic, defaultName, defaultPassWd)
	require.NoError(t, err)
	infoByPriv, err := CreateAccountWithPrivateKey(defaultPrivateKey, defaultName, defaultPassWd)
	require.NoError(t, err)

	require.Equal(t, infoByMnemo.GetName(), infoByPriv.GetName())
	require.Equal(t, infoByMnemo.GetAddress(), infoByPriv.GetAddress())
	require.Equal(t, infoByMnemo.GetAlgo(), infoByPriv.GetAlgo())
	require.Equal(t, infoByMnemo.GetPubKey(), infoByPriv.GetPubKey())
	require.Equal(t, infoByMnemo.GetType(), infoByPriv.GetType())

	_, err = CreateAccountWithPrivateKey("", defaultName, defaultPassWd)
	require.Error(t, err)

	_, err = CreateAccountWithPrivateKey(defaultPrivateKey, "", "")
	require.NoError(t, err)
}

func TestGenerateMnemonic(t *testing.T) {
	mnemo, err := GenerateMnemonic()
	require.NoError(t, err)
	require.NotNil(t, mnemo)
}

func TestGeneratePrivateKeyFromMnemo(t *testing.T) {
	priKey, err := GeneratePrivateKeyFromMnemo(defaultMnemonic)
	require.NoError(t, err)
	require.Equal(t, defaultPrivateKey, priKey)

	_, err = GeneratePrivateKeyFromMnemo("")
	require.Error(t, err)

	// bad mnemonic, add one word in it
	// https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt
	invalidMnemo := fmt.Sprintf("%s abandon", defaultMnemonic)
	_, err = GeneratePrivateKeyFromMnemo(invalidMnemo)
	require.Error(t, err)
}
