package types

import (
	"github.com/tendermint/tendermint/crypto"
	"time"
)

type DelegatorResp struct {
	DelegatorAddress     AccAddress   `json:"delegator_address"`
	ValidatorAddresses   []ValAddress `json:"validator_address"`
	Shares               Dec          `json:"shares"`
	Tokens               Dec          `json:"tokens" `
	UnbondedTokens       Dec          `json:"unbonded_tokens"`
	CompletionTime       time.Time    `json:"completion_time"`
	IsProxy              bool         `json:"is_proxy"`
	TotalDelegatedTokens Dec          `json:"total_delegated_tokens"`
	ProxyAddress         AccAddress   `json:"proxy_address"`
}

type Delegator struct {
	DelegatorAddress     AccAddress   `json:"delegator_address"`
	ValidatorAddresses   []ValAddress `json:"validator_address"`
	Shares               Dec          `json:"shares"`
	Tokens               Dec          `json:"tokens"`
	IsProxy              bool         `json:"is_proxy"`
	TotalDelegatedTokens Dec          `json:"total_delegated_tokens"`
	ProxyAddress         AccAddress   `json:"proxy_address"`
}

// NewDelegator creates a new instance of Delegator
func NewDelegator(delAddr AccAddress) Delegator {
	return Delegator{
		delAddr,
		nil,
		ZeroDec(),
		ZeroDec(),
		false,
		ZeroDec(),
		nil,
	}
}

type Undelegation struct {
	DelegatorAddress AccAddress `json:"delegator_address"`
	Quantity         Dec        `json:"quantity"`
	CompletionTime   time.Time  `json:"completion_time"`
}

// DefaultUndelegation returns default entity for Undelegation
func DefaultUndelegation() Undelegation {
	return Undelegation{
		nil, ZeroDec(), time.Unix(0, 0).UTC(),
	}
}

type Description struct {
	Moniker  string `json:"moniker"`  // name
	Identity string `json:"identity"` // optional identity signature (ex. UPort or Keybase)
	Website  string `json:"website"`  // optional website link
	Details  string `json:"details"`  // optional details
}

func NewDescription(moniker, identity, website, details string) Description {
	return Description{
		Moniker:  moniker,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

type CommissionRates struct {
	Rate          Dec `json:"rate"`            // the commission rate charged to delegators, as a fraction
	MaxRate       Dec `json:"max_rate"`        // maximum commission rate which validator can ever charge, as a fraction
	MaxChangeRate Dec `json:"max_change_rate"` // maximum daily increase of the validator commission, as a fraction
}

func NewCommissionRates(rate, maxRate, maxChangeRate Dec) CommissionRates {
	return CommissionRates{
		Rate:          rate,
		MaxRate:       maxRate,
		MaxChangeRate: maxChangeRate,
	}
}

// Commission defines a commission parameters for a given validator.
type Commission struct {
	CommissionRates `json:"commission_rates"`
	UpdateTime      time.Time `json:"update_time"` // the last time the commission rate was changed
}

type Validator struct {
	OperatorAddress         ValAddress    `json:"operator_address"`    // address of the validator's operator; bech encoded in JSON
	ConsPubKey              crypto.PubKey `json:"consensus_pubkey"`    // the consensus public key of the validator; bech encoded in JSON
	Jailed                  bool          `json:"jailed"`              // has the validator been jailed from bonded status?
	Status                  byte          `json:"status"`              // validator status (bonded/unbonding/unbonded)
	Tokens                  Int           `json:"tokens"`              // delegated tokens (incl. self-delegation)
	DelegatorShares         Dec           `json:"delegator_shares"`    // total shares issued to a validator's delegators
	Description             Description   `json:"description"`         // description terms for the validator
	UnbondingHeight         int64         `json:"unbonding_height" `   // if unbonding, height at which this validator has begun unbonding
	UnbondingCompletionTime time.Time     `json:"unbonding_time"`      // if unbonding, min time for the validator to complete unbonding
	Commission              Commission    `json:"commission"`          // commission parameters
	MinSelfDelegation       Dec           `json:"min_self_delegation"` // validator's self declared minimum self delegation
}

// this is a helper struct used for JSON de- and encoding only
type bechValidator struct {
	OperatorAddress         ValAddress  `json:"operator_address"`    // address of the validator's operator; bech encoded in JSON
	ConsPubKey              string      `json:"consensus_pubkey"`    // the consensus public key of the validator; bech encoded in JSON
	Jailed                  bool        `json:"jailed"`              // has the validator been jailed from bonded status?
	Status                  byte        `json:"status"`              // validator status (bonded/unbonding/unbonded)
	Tokens                  Int         `json:"tokens"`              // delegated tokens (incl. self-delegation)
	DelegatorShares         Dec         `json:"delegator_shares"`    // total shares issued to a validator's delegators
	Description             Description `json:"description"`         // description terms for the validator
	UnbondingHeight         int64       `json:"unbonding_height" `   // if unbonding, height at which this validator has begun unbonding
	UnbondingCompletionTime time.Time   `json:"unbonding_time"`      // if unbonding, min time for the validator to complete unbonding
	Commission              Commission  `json:"commission"`          // commission parameters
	MinSelfDelegation       Dec         `json:"min_self_delegation"` // validator's self declared minimum self delegation
}

// UnmarshalJSON unmarshals the validator from JSON using Bech32
func (v *Validator) UnmarshalJSON(data []byte) error {
	bv := &bechValidator{}
	if err := MsgCdc.UnmarshalJSON(data, bv); err != nil {
		return err
	}
	consPubKey, err := GetConsPubKeyBech32(bv.ConsPubKey)
	if err != nil {
		return err
	}
	*v = Validator{
		OperatorAddress:         bv.OperatorAddress,
		ConsPubKey:              consPubKey,
		Jailed:                  bv.Jailed,
		Tokens:                  bv.Tokens,
		Status:                  bv.Status,
		DelegatorShares:         bv.DelegatorShares,
		Description:             bv.Description,
		UnbondingHeight:         bv.UnbondingHeight,
		UnbondingCompletionTime: bv.UnbondingCompletionTime,
		Commission:              bv.Commission,
		MinSelfDelegation:       bv.MinSelfDelegation,
	}
	return nil
}
