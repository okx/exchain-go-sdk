package farm

import (
	"github.com/okex/okexchain-go-sdk/module/farm/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
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

// QueryPool gets the farm pool info by its pool name
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

// QueryAccount gets the name of pools that an account has locked coins in
func (fc farmClient) QueryAccount(accAddrStr string) (poolNames []string, err error) {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	if err != nil {
		return
	}

	accParams := params.NewQueryAccountParams(accAddr)
	jsonBytes, err := fc.GetCodec().MarshalJSON(accParams)
	if err != nil {
		return poolNames, utils.ErrMarshalJSON(err.Error())
	}

	res, err := fc.Query(types.QueryAccountPath, jsonBytes)
	if err != nil {
		return poolNames, utils.ErrClientQuery(err.Error())
	}

	fc.GetCodec().MustUnmarshalJSON(res, &poolNames)
	return
}

// QueryAccountsLockedTo gets all addresses of accounts that have locked coins in a pool
func (fc farmClient) QueryAccountsLockedTo(poolName string) (accAddrs []sdk.AccAddress, err error) {
	poolParams := params.NewQueryPoolParams(poolName)
	jsonBytes, err := fc.GetCodec().MarshalJSON(poolParams)
	if err != nil {
		return accAddrs, utils.ErrMarshalJSON(err.Error())
	}

	res, err := fc.Query(types.QueryAccountsLockedToPath, jsonBytes)
	if err != nil {
		return accAddrs, utils.ErrClientQuery(err.Error())
	}

	fc.GetCodec().MustUnmarshalJSON(res, &accAddrs)
	return
}
