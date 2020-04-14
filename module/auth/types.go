package auth

import "github.com/okex/okchain-go-sdk/types"

// const
const (
	ModuleName = "auth"

	accountInfoPath = "/store/acc/key"
)

var addressStoreKeyPrefix = []byte{0x01}

func addressStoreKey(accAddr types.AccAddress) []byte {
	return append(addressStoreKeyPrefix, accAddr.Bytes()...)
}
