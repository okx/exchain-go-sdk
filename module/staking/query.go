package staking

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/staking/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/okex/okexchain-go-sdk/utils"
	stakingtypes "github.com/okex/okexchain/x/staking/types"
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
//func (sc stakingClient) QueryValidator(valAddrStr string) (val types.Validator, err error) {
//	valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
//	if err != nil {
//		return
//	}
//
//	res, _, err := sc.QueryStore(types.GetValidatorKey(valAddr), ModuleName, "key")
//	if err != nil {
//		return
//	}
//	if len(res) == 0 {
//		return val, fmt.Errorf("failed. no validator found with address %s", valAddrStr)
//	}
//
//	var innerVal types.ValidatorInner
//	sc.GetCodec().MustUnmarshalBinaryLengthPrefixed(res, &innerVal)
//
//	return innerVal.Standardize()
//
//}

// QueryDelegator gets the detail info of a delegator
func (sc stakingClient) QueryDelegator(delAddrStr string) (delResp types.DelegatorResp, err error) {
	delAddr, err := sdk.AccAddressFromBech32(delAddrStr)
	if err != nil {
		return
	}

	resp, _, err := sc.QueryStore(types.GetDelegatorKey(delAddr), types.ModuleName, "key")
	if err != nil {
		return delResp, utils.ErrClientQuery(err.Error())
	}

	delegator, undelegation := types.NewDelegator(delAddr), types.DefaultUndelegation()
	if len(resp) != 0 {
		sc.GetCodec().MustUnmarshalBinaryLengthPrefixed(resp, &delegator)
	}

	// query for the undelegation info
	jsonBytes, err := sc.GetCodec().MarshalJSON(params.NewQueryDelegatorParams(delAddr))
	if err != nil {
		return delResp, utils.ErrMarshalJSON(err.Error())
	}

	res, _, err := sc.Query(types.UnbondDelegationPath, jsonBytes)
	// if err!= nil , we treat it as there's no undelegation of the delegator
	if err == nil {
		if err = sc.GetCodec().UnmarshalJSON(res, &undelegation); err != nil {
			return
		}
	}

	return types.ConvertToDelegatorResp(delegator, undelegation), nil
}
