package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	defaultMnemonic   = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	defaultAddr       = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	defaultPrivateKey = "EA6D97F31E4B70663594DD6AFC3E3550AAB5FDD9C44305E8F8F2003023B27FDA"
	defaultMemo       = "my memo"
	valConsPK         = "okexchainvalconspub1zcjduepqpjq9n8g6fnjrys5t07cqcdcptu5d06tpxvhdu04mdrc4uc5swmmqttvmqv"
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
}

func TestGenerateMnemonic(t *testing.T) {
	mnemo, err := GenerateMnemonic()
	require.NoError(t, err)
	require.NotNil(t, mnemo)
}
