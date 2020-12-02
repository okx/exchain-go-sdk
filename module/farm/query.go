package farm

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// QueryAccount gets the name of pools that an account has locked coins in
func (fc farmClient) QueryAccount(accAddrStr string) (poolNames []string, err error) {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	if err != nil {
		return
	}

	jsonBytes, err := fc.GetCodec().MarshalJSON(farmtypes.NewQueryAccountParams(accAddr))
	if err != nil {
		return poolNames, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", farmtypes.QuerierRoute, farmtypes.QueryAccount)
	res, _, err := fc.Query(path, jsonBytes)
	if err != nil {
		return poolNames, utils.ErrClientQuery(err.Error())
	}

	if err = fc.GetCodec().UnmarshalJSON(res, &poolNames); err != nil {
		return poolNames, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
