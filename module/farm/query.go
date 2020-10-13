package farm

import "github.com/okex/okexchain-go-sdk/module/farm/types"

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
