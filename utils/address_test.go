package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	accAddrWithOKExChainPrefix = "okexchain1qj5c07sm6jetjz8f509qtrxgh4psxkv32x0qas"
)

func TestAccAddrPrefixConvert(t *testing.T) {
	dstAccAddrStr, err := AccAddrPrefixConvert("ex", defaultAddr, "okexchain")
	require.NoError(t, err)
	require.Equal(t, accAddrWithOKExChainPrefix, dstAccAddrStr)

	// error check
	// wrong source prefix
	_, err = AccAddrPrefixConvert("exx", defaultAddr, "okexchain")
	require.Error(t, err)

	// wrong source account address
	_, err = AccAddrPrefixConvert("ex", defaultAddr+"a", "okexchain")
	require.Error(t, err)

	// wrong destination account address
	dstAccAddrStr, err = AccAddrPrefixConvert("ex", defaultAddr, "okexchainx")
	require.NoError(t, err)
	require.NotEqual(t, accAddrWithOKExChainPrefix, dstAccAddrStr)
}
