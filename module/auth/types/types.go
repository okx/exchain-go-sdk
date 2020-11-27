package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

// const
const (
	ModuleName = "auth"

	AccountInfoPath = "/store/acc/key"
)

var addressStoreKeyPrefix = []byte{0x01}

// GetAddressStoreKey gets the store key for an account
func GetAddressStoreKey(accAddr sdk.AccAddress) []byte {
	return append(addressStoreKeyPrefix, accAddr.Bytes()...)
}

type Account = exported.Account
