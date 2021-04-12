package dex

import (
	"encoding/json"
	"fmt"

	"github.com/okex/exchain-go-sdk/module/dex/types"
	"github.com/okex/exchain-go-sdk/utils"
	dextypes "github.com/okex/exchain/x/dex/types"
)

// QueryProducts gets token pair info
func (dc dexClient) QueryProducts(ownerAddr string, page, perPage int) (tokenPairs []types.TokenPair, err error) {
	jsonBytes, err := dc.GetCodec().MarshalJSON(dextypes.NewQueryDexInfoParams(ownerAddr, page, perPage))
	if err != nil {
		return
	}

	path := fmt.Sprintf("custom/%s/%s", dextypes.QuerierRoute, dextypes.QueryProducts)
	res, _, err := dc.Query(path, jsonBytes)
	if err != nil {
		return
	}

	var response types.ListResponse
	if err = json.Unmarshal(res, &response); err != nil {
		return tokenPairs, utils.ErrUnmarshalJSON(err.Error())
	}

	// TODO:
	// not friendly
	return response.Data.Data, err
}
