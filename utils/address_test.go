package utils

import (
	"testing"

	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	defaultValAddr             = "exvaloper1qj5c07sm6jetjz8f509qtrxgh4psxkv3m2wy6x"
	accAddrWithOKExChainPrefix = "okexchain1qj5c07sm6jetjz8f509qtrxgh4psxkv32x0qas"
	valAddrWithOKExChainPrefix = "okexchainvaloper1qj5c07sm6jetjz8f509qtrxgh4psxkv3tzllpj"
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

	// recover the account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("ex", "expub")
}

func TestValAddrPrefixConvert(t *testing.T) {
	dstValAddrStr, err := ValAddrPrefixConvert("exvaloper", defaultValAddr, "okexchainvaloper")
	require.NoError(t, err)
	require.Equal(t, valAddrWithOKExChainPrefix, dstValAddrStr)

	// error check
	// wrong source prefix
	_, err = ValAddrPrefixConvert("exxvaloper", defaultValAddr, "okexchainvaloper")
	require.Error(t, err)

	// wrong source validator address
	_, err = ValAddrPrefixConvert("exvaloper", defaultValAddr+"a", "okexchainvaloper")
	require.Error(t, err)

	// wrong destination validator address
	dstValAddrStr, err = ValAddrPrefixConvert("exvaloper", defaultValAddr, "okexchainxvaloper")
	require.NoError(t, err)
	require.NotEqual(t, valAddrWithOKExChainPrefix, dstValAddrStr)

	// recover the validator prefix
	sdk.GetConfig().SetBech32PrefixForValidator("exvaloper", "exvaloperpub")
}
