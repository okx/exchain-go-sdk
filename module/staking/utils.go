package staking

import "github.com/okex/okchain-go-sdk/exposed"

func convertToDelegatorResp(delegator Delegator, undelegation Undelegation) exposed.DelegatorResp {
	return exposed.DelegatorResp{
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
