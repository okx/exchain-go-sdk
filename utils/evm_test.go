package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToCosmosAddress(t *testing.T) {
	expectedAccAddr, err := sdk.AccAddressFromBech32(defaultAddr)
	require.NoError(t, err)

	accAddr, err := ToCosmosAddress(defaultAddr)
	require.NoError(t, err)
	require.True(t, accAddr.Equals(expectedAccAddr))

	accAddr, err = ToCosmosAddress(defaultAddrEth)
	require.NoError(t, err)
	require.True(t, accAddr.Equals(expectedAccAddr))

	accAddr, err = ToCosmosAddress(defaultAddrEth[2:])
	require.NoError(t, err)
	require.True(t, accAddr.Equals(expectedAccAddr))

	_, err = ToCosmosAddress(defaultAddr + "a")
	require.Error(t, err)

	_, err = ToCosmosAddress(defaultAddrEth + "g")
	require.Error(t, err)
}
