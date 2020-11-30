package utils

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"io/ioutil"
)

// GetStdTxFromFile gets the instance of stdTx from a json file
func GetStdTxFromFile(codec gosdktypes.SDKCodec, filePath string) (stdTx authtypes.StdTx, err error) {
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