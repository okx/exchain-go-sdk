package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/okex/okchain-go-sdk/common/libs/pkg/errors"
	"github.com/okex/okchain-go-sdk/crypto/go-bip39"
	"github.com/okex/okchain-go-sdk/crypto/keys/hd"
	"github.com/okex/okchain-go-sdk/types"
	"io/ioutil"
)

// GetStdTxFromFile gets the instance of stdTx from a json file
func GetStdTxFromFile(filePath string) (stdTx types.StdTx, err error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	types.MsgCdc.MustUnmarshalJSON(bytes, &stdTx)

	return
}

// ParseValAddresses parses validator address string to types.ValAddress
func ParseValAddresses(valAddrsStr []string) ([]types.ValAddress, error) {
	valLen := len(valAddrsStr)
	valAddrs := make([]types.ValAddress, valLen)
	var err error
	for i := 0; i < valLen; i++ {
		valAddrs[i], err = types.ValAddressFromBech32(valAddrsStr[i])
		if err != nil {
			return nil, fmt.Errorf("invalid validator address: %s", valAddrsStr[i])
		}
	}
	return valAddrs, nil
}

// GeneratePrivateKeyFromMnemo converts mnemonic to private key
func GeneratePrivateKeyFromMnemo(mnemo string) (privKey string, err error) {
	hdPath := hd.NewFundraiserParams(0, 0)
	seed, err := bip39.NewSeedWithErrorChecking(mnemo, "")
	if err != nil {
		return
	}
	masterPrivateKey, ch := hd.ComputeMastersFromSeed(seed)
	derivedPrivateKey, err := hd.DerivePrivateKeyForPath(masterPrivateKey, ch, hdPath.String())
	return hex.EncodeToString(derivedPrivateKey[:]), nil
}

func sliceToArray(s []byte) (byteArray [32]byte, err error) {
	if len(s) != 32 {
		return byteArray, errors.New("failed. byte slice's length is not 32")
	}
	for i := 0; i < 32; i++ {
		byteArray[i] = s[i]
	}
	return
}
