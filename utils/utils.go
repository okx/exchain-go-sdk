package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/bartekn/go-bip39"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetStdTxFromFile gets the instance of stdTx from a json file
func GetStdTxFromFile(codec sdk.SDKCodec, filePath string) (stdTx sdk.StdTx, err error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	codec.MustUnmarshalJSON(bytes, &stdTx)
	return
}

// ParseValAddresses parses validator address string to types.ValAddress
func ParseValAddresses(valAddrsStr []string) ([]sdk.ValAddress, error) {
	valLen := len(valAddrsStr)
	valAddrs := make([]sdk.ValAddress, valLen)
	var err error
	for i := 0; i < valLen; i++ {
		valAddrs[i], err = sdk.ValAddressFromBech32(valAddrsStr[i])
		if err != nil {
			return nil, fmt.Errorf("invalid validator address: %s", valAddrsStr[i])
		}
	}
	return valAddrs, nil
}

// GeneratePrivateKeyFromMnemo converts mnemonic to private key
func GeneratePrivateKeyFromMnemo(mnemo string) (privKey string, err error) {
	// TODO
	hdPath := hd.NewFundraiserParams(0, 0, 0)
	seed, err := bip39.NewSeedWithErrorChecking(mnemo, "")
	if err != nil {
		return
	}
	masterPrivateKey, ch := hd.ComputeMastersFromSeed(seed)
	derivedPrivateKey, err := hd.DerivePrivateKeyForPath(masterPrivateKey, ch, hdPath.String())
	if err != nil {
		return
	}
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
