package utils

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	valAddrStr = "okexchainvaloper1ntvyep3suq5z7789g7d5dejwzameu08mmv8pca"
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