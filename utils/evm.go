package utils

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

// ToCosmosAddress converts string address of cosmos and ethereum style to cosmos address
func ToCosmosAddress(addrStr string) (toAddr sdk.AccAddress, err error) {
	if strings.HasPrefix(addrStr, sdk.GetConfig().GetBech32AccountAddrPrefix()) {
		toAddr, err = sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return toAddr, fmt.Errorf("failed. invalid bech32 formatted address: %s", err)
		}
		return
	}

	// strip 0x prefix if exists
	addrStr = strings.TrimPrefix(addrStr, "0x")
	return sdk.AccAddressFromHex(addrStr)
}
