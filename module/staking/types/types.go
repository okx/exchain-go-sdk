package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/staking"
	stakingtypes "github.com/okex/okexchain/x/staking/types"
	"time"
)

// const
const (
	ModuleName = stakingtypes.ModuleName

	UnbondDelegationPath = "custom/staking/unbondingDelegation"

	defaultMinSelfDelegation = "0.001okt"
)

var (
	msgCdc = gosdktypes.NewCodec()
	// ValidatorsKey is useful for subspace and store query about validator
	ValidatorsKey = []byte{0x21}
	// DelegatorKey is useful for subspace and store query about delegator
	DelegatorKey = []byte{0x52}
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for staking module
func RegisterCodec(cdc *codec.Codec) {
	staking.RegisterCodec(cdc)
}

// CommissionRates is a part of Commission
type CommissionRates struct {
	Rate          sdk.Dec `json:"rate"`
	MaxRate       sdk.Dec `json:"max_rate"`
	MaxChangeRate sdk.Dec `json:"max_change_rate"`
}

// Commission defines a commission parameters for a given validator
type Commission struct {
	CommissionRates `json:"commission_rates"`
	UpdateTime      time.Time `json:"update_time"`
}

// Delegator is the struct of the info of a delegator
type Delegator struct {
	DelegatorAddress     sdk.AccAddress   `json:"delegator_address"`
	ValidatorAddresses   []sdk.ValAddress `json:"validator_address"`
	Shares               sdk.Dec          `json:"shares"`
	Tokens               sdk.Dec          `json:"tokens"`
	IsProxy              bool             `json:"is_proxy"`
	TotalDelegatedTokens sdk.Dec          `json:"total_delegated_tokens"`
	ProxyAddress         sdk.AccAddress   `json:"proxy_address"`
}

// NewDelegator creates a new instance of Delegator
func NewDelegator(delAddr sdk.AccAddress) Delegator {
	return Delegator{
		delAddr,
		nil,
		sdk.ZeroDec(),
		sdk.ZeroDec(),
		false,
		sdk.ZeroDec(),
		nil,
	}
}

// Undelegation is the struct of the info for unbonding
type Undelegation struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address"`
	Quantity         sdk.Dec        `json:"quantity"`
	CompletionTime   time.Time      `json:"completion_time"`
}

// DefaultUndelegation returns default entity for Undelegation
func DefaultUndelegation() Undelegation {
	return Undelegation{
		nil, sdk.ZeroDec(), time.Unix(0, 0).UTC(),
	}
}

// GetValidatorKey builds the store key for a specific validator
func GetValidatorKey(valAddr sdk.ValAddress) []byte {
	return append(ValidatorsKey, valAddr.Bytes()...)
}

// GetDelegatorKey builds the store key for a specific delegator
func GetDelegatorKey(delAddr sdk.AccAddress) []byte {
	return append(DelegatorKey, delAddr.Bytes()...)
}

// ConvertToDelegatorResp builds DelegatorResp with the info of Delegator and Undelegation
func ConvertToDelegatorResp(delegator Delegator, undelegation Undelegation) DelegatorResp {
	return DelegatorResp{
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

// DelegatorResp is designed only for delegator query
type DelegatorResp struct {
	DelegatorAddress     sdk.AccAddress   `json:"delegator_address"`
	ValidatorAddresses   []sdk.ValAddress `json:"validator_address"`
	Shares               sdk.Dec          `json:"shares"`
	Tokens               sdk.Dec          `json:"tokens" `
	UnbondedTokens       sdk.Dec          `json:"unbonded_tokens"`
	CompletionTime       time.Time        `json:"completion_time"`
	IsProxy              bool             `json:"is_proxy"`
	TotalDelegatedTokens sdk.Dec          `json:"total_delegated_tokens"`
	ProxyAddress         sdk.AccAddress   `json:"proxy_address"`
}
