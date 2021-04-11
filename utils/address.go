package utils

import sdk "github.com/cosmos/cosmos-sdk/types"

// AccAddrPrefixConvert converts the account address between two different prefix
func AccAddrPrefixConvert(srcPrefx, srcAccAddrStr, dstPrefix string) (dstAccAddrStr string, err error) {
	config := sdk.GetConfig()
	// set source prefix
	config.SetBech32PrefixForAccount(srcPrefx, srcPrefx+"pub")
	accAddr, err := sdk.AccAddressFromBech32(srcAccAddrStr)
	if err != nil {
		return
	}

	config.SetBech32PrefixForAccount(dstPrefix, dstPrefix+"pub")
	return accAddr.String(), err
}
