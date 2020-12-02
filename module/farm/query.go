package farm

import (
	"fmt"
	"github.com/okex/okexchain-go-sdk/module/farm/types"
	"github.com/okex/okexchain-go-sdk/utils"
	farmtypes "github.com/okex/okexchain/x/farm/types"
)

// QueryPools gets all farm pools info
func (fc farmClient) QueryPools() (farmPools []types.FarmPool, err error) {
	// fixed to all pools query
	jsonBytes, err := fc.GetCodec().MarshalJSON(farmtypes.NewQueryPoolsParams(1, 0))
	if err != nil {
		return farmPools, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", farmtypes.QuerierRoute, farmtypes.QueryPools)
	res, _, err := fc.Query(path, jsonBytes)
	if err != nil {
		return farmPools, utils.ErrClientQuery(err.Error())
	}

	if err = fc.GetCodec().UnmarshalJSON(res, &farmPools); err != nil {
		return farmPools, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

// QueryPool gets the farm pool info by its pool name
func (fc farmClient) QueryPool(poolName string) (farmPool types.FarmPool, err error) {
	jsonBytes, err := fc.GetCodec().MarshalJSON(farmtypes.NewQueryPoolParams(poolName))
	if err != nil {
		return farmPool, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", farmtypes.QuerierRoute, farmtypes.QueryPool)
	res, _, err := fc.Query(path, jsonBytes)
	if err != nil {
		return farmPool, utils.ErrClientQuery(err.Error())
	}

	if err = fc.GetCodec().UnmarshalJSON(res, &farmPool); err != nil {
		return farmPool, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
