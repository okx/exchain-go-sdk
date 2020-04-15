package staking

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/types"
)

// QueryValidators gets all the validators info
func (sc stakingClient) QueryValidators() (vals []exposed.Validator, err error) {
	resKVs, err := sc.QuerySubspace(types.ValidatorsKey, ModuleName)
	if err != nil {
		return
	}

	for _, kv := range resKVs {
		var innerVal validator
		sc.GetCodec().MustUnmarshalBinaryLengthPrefixed(kv.Value, &innerVal)
		val, err := innerVal.standardize()
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}

	return

}

// QueryValidator gets the info of a specific validator
func (sc stakingClient) QueryValidator(valAddrStr string) (val exposed.Validator, err error) {
	valAddr, err := types.ValAddressFromBech32(valAddrStr)
	if err != nil {
		return
	}

	res, err := sc.QueryStore(getValidatorKey(valAddr), ModuleName, "key")
	if err != nil {
		return
	}
	if len(res) == 0 {
		return val, fmt.Errorf("failed. no validator found with address %s", valAddrStr)
	}

	var innerVal validator
	sc.GetCodec().MustUnmarshalBinaryLengthPrefixed(res, &innerVal)

	return innerVal.standardize()

}

// QueryDelegator gets the detail info of a delegator
func (sc stakingClient) QueryDelegator(delAddrStr string) (delResp exposed.DelegatorResp, err error) {
	delAddr, err := types.AccAddressFromBech32(delAddrStr)
	if err != nil {
		return
	}

	resp, err := sc.QueryStore(types.GetDelegatorKey(delAddr), ModuleName, "key")
	if err != nil {
		return delResp, fmt.Errorf("ok client query error : %s", err.Error())
	}

	delegator, undelegation := NewDelegator(delAddr), defaultUndelegation()
	if len(resp) != 0 {
		sc.GetCodec().MustUnmarshalBinaryLengthPrefixed(resp, &delegator)
	}

	// query for the undelegation info
	jsonBytes, err := sc.GetCodec().MarshalJSON(params.NewQueryDelegatorParams(delAddr))
	if err != nil {
		return delResp, fmt.Errorf("error : QueryDelegatorParams failed in json marshal : %s", err.Error())
	}

	res, err := sc.Query(unbondDelegationPath, jsonBytes)
	// if err!= nil , we treat it as there's no undelegation of the delegator
	if err == nil {
		if err = sc.GetCodec().UnmarshalJSON(res, &undelegation); err != nil {
			return
		}
	}

	return convertToDelegatorResp(delegator, undelegation), nil
}
