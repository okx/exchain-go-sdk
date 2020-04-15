package dex

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/module/dex/types"
)

// QueryProducts gets token pair info
func (dc dexClient) QueryProducts(ownerAddr string, page, perPage int) (tokenPairs []types.TokenPair, err error) {
	queryParams, err := params.NewQueryDexInfoParams(ownerAddr, page, perPage)
	if err != nil {
		return
	}

	jsonBytes, err := dc.GetCodec().MarshalJSON(queryParams)
	if err != nil {
		return
	}

	res, err := dc.Query(types.ProductsPath, jsonBytes)
	if err != nil {
		return
	}

	if err = dc.GetCodec().UnmarshalJSON(res, &tokenPairs); err != nil {
		return tokenPairs, fmt.Errorf("failed. unmarshal JSON error: %s", err.Error())
	}

	return
}
