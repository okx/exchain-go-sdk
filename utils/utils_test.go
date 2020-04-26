package utils

import (
	"fmt"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	valAddrStr = "okchainvaloper1alq9na49n9yycysh889rl90g9nhe58lcs50wu5"
)

func TestParseValAddresses(t *testing.T) {
	valAddrsStr := []string{valAddrStr}
	valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
	require.NoError(t, err)

	valAddrs, err := ParseValAddresses(valAddrsStr)
	require.NoError(t, err)
	require.Equal(t, 1, len(valAddrs))
	require.Equal(t, valAddr, valAddrs[0])

	// bad val address
	valAddrsStr = append(valAddrsStr, valAddrStr[1:])
	_, err = ParseValAddresses(valAddrsStr)
	require.Error(t, err)
}

func TestGeneratePrivateKeyFromMnemo(t *testing.T) {
	priKey, err := GeneratePrivateKeyFromMnemo(defaultMnemonic)
	require.NoError(t, err)
	require.Equal(t, defaultPrivateKey, priKey)

	// bad mnemonic, add one word in it
	_, err = GeneratePrivateKeyFromMnemo(fmt.Sprintf("%s %s", defaultMnemonic, "offer"))
	require.Error(t, err)
}
