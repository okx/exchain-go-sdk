package farm

import (
	"github.com/okex/okexchain-go-sdk/module/farm/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/okex/okexchain-go-sdk/utils"
)

// QueryPools gets all farm pools info
func (fc farmClient) QueryPools() (farmPools []types.FarmPool, err error) {
	// fixed to all pools query
	poolsParams := params.NewQueryPoolsParams(1, 0)
	jsonBytes, err := fc.GetCodec().MarshalJSON(poolsParams)
	if err != nil {
		return farmPools, utils.ErrMarshalJSON(err.Error())
	}

	res, err := fc.Query(types.QueryPoolsPath, jsonBytes)
	if err != nil {
		return farmPools, utils.ErrClientQuery(err.Error())
	}

	fc.GetCodec().MustUnmarshalJSON(res, &farmPools)
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

// QueryLockInfo gets the locked info of an account in a specific pool
func (fc farmClient) QueryLockInfo(poolName, accAddrStr string) (lockInfo types.LockInfo, err error) {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	if err != nil {
		return
	}
	poolParams := params.NewQueryPoolAccountParams(poolName, accAddr)
	jsonBytes, err := fc.GetCodec().MarshalJSON(poolParams)
	if err != nil {
		return lockInfo, utils.ErrMarshalJSON(err.Error())
	}

	res, err := fc.Query(types.QueryLockInfoPath, jsonBytes)
	if err != nil {
		return lockInfo, utils.ErrClientQuery(err.Error())
	}

	fc.GetCodec().MustUnmarshalJSON(res, &lockInfo)
	return
}
