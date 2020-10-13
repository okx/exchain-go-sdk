package farm

import (
	"github.com/okex/okexchain-go-sdk/module/farm/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/okex/okexchain-go-sdk/utils"
)

// QueryPools gets all farm pools info
func (fc farmClient) QueryPools() (farmPools []types.FarmPool, err error) {
	resKVs, err := fc.QuerySubspace(types.FarmPoolPrefix, ModuleName)
	if err != nil {
		return
	}

	for _, kv := range resKVs {
		var farmPool types.FarmPool
		fc.GetCodec().MustUnmarshalBinaryLengthPrefixed(kv.Value, &farmPool)
		farmPools = append(farmPools, farmPool)
	}

	return
}

// QueryPools gets the farm pool info by its pool name
func (fc farmClient) QueryPool(poolName string) (farmPool types.FarmPool, err error) {
	poolParams := params.NewQueryPoolParams(poolName)
	jsonBytes, err := fc.GetCodec().MarshalJSON(poolParams)
	if err != nil {
		return farmPool, utils.ErrMarshalJSON(err.Error())
	}

	res, err := fc.Query(types.QueryPoolPath, jsonBytes)
	if err != nil {
		return farmPool, utils.ErrClientQuery(err.Error())
	}

	fc.GetCodec().MustUnmarshalJSON(res, &farmPool)
	return
}
