package staking

import (
	"fmt"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain-go-sdk/module/staking/types"
	"github.com/okex/exchain-go-sdk/utils"
	stakingtypes "github.com/okex/exchain/x/staking/types"
)

// QueryValidators gets all the validators info
func (sc stakingClient) QueryValidators() (vals []types.Validator, err error) {
	jsonBytes, err := sc.GetCodec().MarshalJSON(stakingtypes.NewQueryValidatorsParams(1, 0, "all"))
	if err != nil {
		return vals, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", stakingtypes.QuerierRoute, stakingtypes.QueryValidators)
	res, _, err := sc.Query(path, jsonBytes)
	if err != nil {
		return vals, utils.ErrClientQuery(err.Error())
	}

	if err = sc.GetCodec().UnmarshalJSON(res, &vals); err != nil {
		return vals, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

// QueryValidator gets the info of a specific validator
func (sc stakingClient) QueryValidator(valAddrStr string) (val types.Validator, err error) {
	valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
	if err != nil {
		return
	}

	jsonBytes, err := sc.GetCodec().MarshalJSON(stakingtypes.NewQueryValidatorParams(valAddr))
	if err != nil {
		return val, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", stakingtypes.QuerierRoute, stakingtypes.QueryValidator)
	res, _, err := sc.Query(path, jsonBytes)
	if err != nil {
		return val, utils.ErrClientQuery(err.Error())
	}

	if err = sc.GetCodec().UnmarshalJSON(res, &val); err != nil {
		return val, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

// QueryDelegator gets the detail info of a delegator
func (sc stakingClient) QueryDelegator(delAddrStr string) (delResp types.DelegatorResponse, err error) {
	delAddr, err := sdk.AccAddressFromBech32(delAddrStr)
	if err != nil {
		return
	}

	delegator, undelegation := stakingtypes.NewDelegator(delAddr), stakingtypes.DefaultUndelegation()
	resp, _, err := sc.QueryStore(stakingtypes.GetDelegatorKey(delAddr), stakingtypes.StoreKey, "key")
	if err != nil {
		return delResp, utils.ErrClientQuery(err.Error())
	}
	if len(resp) != 0 {
		sc.GetCodec().MustUnmarshalBinaryLengthPrefixed(resp, &delegator)
	}

	// query for the undelegation info
	jsonBytes, err := sc.GetCodec().MarshalJSON(stakingtypes.NewQueryDelegatorParams(delAddr))
	if err != nil {
		return delResp, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/%s", stakingtypes.RouterKey, stakingtypes.QueryUnbondingDelegation)
	res, _, err := sc.Query(path, jsonBytes)
	// if err!= nil , we treat it as there's no undelegation of the delegator
	if err == nil {
		if err = sc.GetCodec().UnmarshalJSON(res, &undelegation); err != nil {
			return
		}
	}

	return utils.ConvertToDelegatorResponse(delegator, undelegation), nil
}
