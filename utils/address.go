package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
)

// AccAddrPrefixConvert converts the account address between two different prefixes
func AccAddrPrefixConvert(srcPrefx, srcAccAddrStr, dstPrefix string) (dstAccAddrStr string, err error) {
	config := sdk.GetConfig()
	// set source prefix
	config.SetBech32PrefixForAccount(srcPrefx, fmt.Sprintf("%s%s", srcPrefx, sdk.PrefixPublic))
	accAddr, err := sdk.AccAddressFromBech32(srcAccAddrStr)
	if err != nil {
		return
	}

	// set destination prefix
	config.SetBech32PrefixForAccount(dstPrefix, fmt.Sprintf("%s%s", dstPrefix, sdk.PrefixPublic))
	return accAddr.String(), err
}

// ValAddrPrefixConvert converts the validator address between two different prefixes
func ValAddrPrefixConvert(srcPrefx, srcValAddrStr, dstPrefix string) (dstValAddrStr string, err error) {
	config := sdk.GetConfig()
	// set source prefix
	config.SetBech32PrefixForValidator(srcPrefx, fmt.Sprintf("%s%s", srcPrefx, sdk.PrefixPublic))
	valAddr, err := sdk.ValAddressFromBech32(srcValAddrStr)
	if err != nil {
		return
	}

	// set destination prefix
	config.SetBech32PrefixForValidator(dstPrefix, fmt.Sprintf("%s%s", dstPrefix, sdk.PrefixPublic))
	return valAddr.String(), err
}

// IsValidHexAdress
func IsValidHexAdress(address string) bool {
	if address == "" {
		return false
	}

	if !strings.HasPrefix(address, "0x") {
		return false
	}

	_, err := hex.DecodeString(address[2:])
	if err != nil {
		return false
	}

	return true
}
