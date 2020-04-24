package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	defaultName       = "alice"
	defaultPassWd     = "12345678"
	defaultMnemonic   = "sustain hole urban away boy core lazy brick wait drive tiger tell"
	defaultPrivateKey = "de0e9d9e7bac1366f7d8719a450dab03c9b704172ba43e0a25a7be1d51c69a87"
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

	_, _, err = CreateAccountWithMnemo(defaultMnemonic, "", defaultPassWd)
	require.NoError(t, err)

	_, _, err = CreateAccountWithMnemo(defaultMnemonic, defaultName, "")
	require.NoError(t, err)

	_, _, err = CreateAccountWithMnemo("", defaultName, defaultPassWd)
	require.Error(t, err)

	_, _, err = CreateAccountWithMnemo(defaultPassWd, defaultName, defaultPassWd)
	require.Error(t, err)
}

func TestCreateAccountWithPrivateKey(t *testing.T) {
	privateKeyStr, err := GeneratePrivateKeyFromMnemo(defaultMnemonic)
	require.NoError(t, err)
	require.Equal(t, defaultPrivateKey, privateKeyStr)

	privInfo, err := CreateAccountWithPrivateKey(privateKeyStr, defaultName, defaultPassWd)
	require.NoError(t, err)

	mnemoInfo, _, err := CreateAccountWithMnemo(defaultMnemonic, defaultName, defaultPassWd)
	require.NoError(t, err)
	require.Equal(t, privInfo.GetPubKey(), mnemoInfo.GetPubKey())
	require.Equal(t, privInfo.GetAddress(), mnemoInfo.GetAddress())
	require.Equal(t, privInfo.GetName(), mnemoInfo.GetName())

	_, err = CreateAccountWithPrivateKey("", defaultName, defaultPassWd)
	require.Error(t, err)

	_, err = CreateAccountWithPrivateKey(defaultMnemonic, defaultName, defaultPassWd)
	require.Error(t, err)
}

func TestGenerateMnemonic(t *testing.T) {
	mnemo, err := GenerateMnemonic()
	require.NoError(t, err)
	require.NotNil(t, mnemo)
}
