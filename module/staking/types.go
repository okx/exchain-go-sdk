package staking

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"time"
)

// const
const (
	ModuleName = "staking"

	unbondDelegationPath = "custom/staking/unbondingDelegation"

	defaultMinSelfDelegation = "0.001okt"
)

var (
	msgCdc        = types.NewCodec()
	validatorsKey = []byte{0x21}
)

func init() {
	registerCodec(msgCdc)
}

// validator is the struct of validator's detail info(inner)
type validator struct {
	OperatorAddress         types.ValAddress    `json:"operator_address"`
	ConsPubKey              crypto.PubKey       `json:"consensus_pubkey"`
	Jailed                  bool                `json:"jailed"`
	Status                  byte                `json:"status"`
	Tokens                  types.Int           `json:"tokens"`
	DelegatorShares         types.Dec           `json:"delegator_shares"`
	Description             exposed.Description `json:"description"`
	UnbondingHeight         int64               `json:"unbonding_height"`
	UnbondingCompletionTime time.Time           `json:"unbonding_time"`
	Commission              Commission          `json:"commission"`
	MinSelfDelegation       types.Dec           `json:"min_self_delegation"`
}

func (v validator) standardize() (val exposed.Validator, err error) {
	bechConsPubKey, err := types.Bech32ifyConsPub(v.ConsPubKey)
	if err != nil {
		return
	}
	return exposed.Validator{
		OperatorAddress:         v.OperatorAddress,
		ConsPubKey:              bechConsPubKey,
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		DelegatorShares:         v.DelegatorShares,
		Description:             v.Description,
		UnbondingHeight:         v.UnbondingHeight,
		UnbondingCompletionTime: v.UnbondingCompletionTime,
		MinSelfDelegation:       v.MinSelfDelegation,
	}, err
}

// CommissionRates is a part of Commission
type CommissionRates struct {
	Rate          types.Dec `json:"rate"`
	MaxRate       types.Dec `json:"max_rate"`
	MaxChangeRate types.Dec `json:"max_change_rate"`
}

// Commission defines a commission parameters for a given validator
type Commission struct {
	CommissionRates `json:"commission_rates"`
	UpdateTime      time.Time `json:"update_time"`
}

// NewDescription creates a new instance of Description
func NewDescription(moniker, identity, website, details string) exposed.Description {
	return exposed.Description{
		Moniker:  moniker,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

// Delegator is the struct of the info of a delegator
type Delegator struct {
	DelegatorAddress     types.AccAddress   `json:"delegator_address"`
	ValidatorAddresses   []types.ValAddress `json:"validator_address"`
	Shares               types.Dec          `json:"shares"`
	Tokens               types.Dec          `json:"tokens"`
	IsProxy              bool               `json:"is_proxy"`
	TotalDelegatedTokens types.Dec          `json:"total_delegated_tokens"`
	ProxyAddress         types.AccAddress   `json:"proxy_address"`
}

// NewDelegator creates a new instance of Delegator
func NewDelegator(delAddr types.AccAddress) Delegator {
	return Delegator{
		delAddr,
		nil,
		types.ZeroDec(),
		types.ZeroDec(),
		false,
		types.ZeroDec(),
		nil,
	}
}

// Undelegation is the struct of the info for unbonding
type Undelegation struct {
	DelegatorAddress types.AccAddress `json:"delegator_address"`
	Quantity         types.Dec        `json:"quantity"`
	CompletionTime   time.Time        `json:"completion_time"`
}

// defaultUndelegation returns default entity for Undelegation
func defaultUndelegation() Undelegation {
	return Undelegation{
		nil, types.ZeroDec(), time.Unix(0, 0).UTC(),
	}
}

func getValidatorKey(valAddr types.ValAddress) []byte {
	return append(validatorsKey, valAddr.Bytes()...)
}

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
