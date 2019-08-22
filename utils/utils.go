package utils

import (
	"encoding/hex"
	"github.com/ok-chain/gosdk/common/libs/pkg/errors"
	"github.com/ok-chain/gosdk/crypto/go-bip39"
	"github.com/ok-chain/gosdk/crypto/keys/hd"
	"github.com/ok-chain/gosdk/types"
)

var (
	AddressStoreKeyPrefix = []byte{0x01}
)

func AddressStoreKey(addr types.AccAddress) []byte {
	return append(AddressStoreKeyPrefix, addr.Bytes()...)
}

func GeneratePrivateKeyFromMnemo(mnemo string) (string, error) {
	hdPath := hd.NewFundraiserParams(0, 0)
	seed, err := bip39.NewSeedWithErrorChecking(mnemo, "")
	if err != nil {
		return "", err
	}
	masterPrivateKey, ch := hd.ComputeMastersFromSeed(seed)
	derivedPrivateKey, err := hd.DerivePrivateKeyForPath(masterPrivateKey, ch, hdPath.String())
	return hex.EncodeToString(derivedPrivateKey[:]), nil
}

func slice2Array(s []byte) (byteArray [32]byte, err error) {
	if len(s) != 32 {
		return byteArray, errors.New("byte slice's length is not 32")
	}
	for i := 0; i < 32; i++ {
		byteArray[i] = s[i]
	}
	return
}
