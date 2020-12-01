package utils

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	stakingcli "github.com/okex/okexchain/x/staking/client/cli"
	stakingtypes "github.com/okex/okexchain/x/staking/types"
	"io/ioutil"
)

// GetStdTxFromFile gets the instance of stdTx from a json file
func GetStdTxFromFile(codec gosdktypes.SDKCodec, filePath string) (stdTx authtypes.StdTx, err error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	codec.MustUnmarshalJSON(bytes, &stdTx)
	return
}

// ParseValAddresses parses validator address string to types.ValAddress
func ParseValAddresses(valAddrsStr []string) ([]sdk.ValAddress, error) {
	valLen := len(valAddrsStr)
	valAddrs := make([]sdk.ValAddress, valLen)
	var err error
	for i := 0; i < valLen; i++ {
		valAddrs[i], err = sdk.ValAddressFromBech32(valAddrsStr[i])
		if err != nil {
			return nil, fmt.Errorf("invalid validator address: %s", valAddrsStr[i])
		}
	}
	return valAddrs, nil
}

// ConvertToDelegatorResponse builds DelegatorResponse with the info of Delegator and UndelegationInfo
func ConvertToDelegatorResponse(delegator stakingtypes.Delegator, undelegation stakingtypes.UndelegationInfo) stakingcli.DelegatorResponse {
	return stakingcli.DelegatorResponse{
		DelegatorAddress:     delegator.DelegatorAddress,
		ValidatorAddresses:   delegator.ValidatorAddresses,
		Shares:               delegator.Shares,
		Tokens:               delegator.Tokens,
		UnbondedTokens:       undelegation.Quantity,
		CompletionTime:       undelegation.CompletionTime,
		IsProxy:              delegator.IsProxy,
		TotalDelegatedTokens: delegator.TotalDelegatedTokens,
		ProxyAddress:         delegator.ProxyAddress,
	}
}
