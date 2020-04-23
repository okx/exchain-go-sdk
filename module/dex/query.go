package dex

import (
	"github.com/okex/okchain-go-sdk/module/dex/types"
	"github.com/okex/okchain-go-sdk/types/params"
	"github.com/okex/okchain-go-sdk/utils"
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
		return tokenPairs, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
