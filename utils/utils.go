package utils

import "github.com/ok-chain/gosdk/types"

var(
	AddressStoreKeyPrefix = []byte{0x01}
)

func AddressStoreKey(addr types.AccAddress) []byte {
	return append(AddressStoreKeyPrefix, addr.Bytes()...)
}

