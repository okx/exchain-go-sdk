package types

import (
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

// const
const (
	ModuleName = "staking"

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
func RegisterCodec(cdc gosdktypes.SDKCodec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "okexchain/staking/MsgCreateValidator")
	cdc.RegisterConcrete(MsgEditValidator{}, "okexchain/staking/MsgEditValidator")
	cdc.RegisterConcrete(MsgDeposit{}, "okexchain/staking/MsgDeposit")
	cdc.RegisterConcrete(MsgWithdraw{}, "okexchain/staking/MsgWithdraw")
	cdc.RegisterConcrete(MsgAddShares{}, "okexchain/staking/MsgAddShares")
	cdc.RegisterConcrete(MsgDestroyValidator{}, "okexchain/staking/MsgDestroyValidator")
	cdc.RegisterConcrete(MsgRegProxy{}, "okexchain/staking/MsgRegProxy")
	cdc.RegisterConcrete(MsgBindProxy{}, "okexchain/staking/MsgBindProxy")
	cdc.RegisterConcrete(MsgUnbindProxy{}, "okexchain/staking/MsgUnbindProxy")
}

// ValidatorInner is the struct of validator's detail info(inner)
type ValidatorInner struct {
	OperatorAddress         sdk.ValAddress `json:"operator_address"`
	ConsPubKey              crypto.PubKey  `json:"consensus_pubkey"`
	Jailed                  bool           `json:"jailed"`
	Status                  byte           `json:"status"`
	Tokens                  sdk.Int        `json:"tokens"`
	DelegatorShares         sdk.Dec        `json:"delegator_shares"`
	Description             Description    `json:"description"`
	UnbondingHeight         int64          `json:"unbonding_height"`
	UnbondingCompletionTime time.Time      `json:"unbonding_time"`
	Commission              Commission     `json:"commission"`
	MinSelfDelegation       sdk.Dec        `json:"min_self_delegation"`
}

// Standardize converts the inner validator to the standard one
func (vi ValidatorInner) Standardize() (val Validator, err error) {
	bechConsPubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, vi.ConsPubKey)
	if err != nil {
		return
	}
	return Validator{
		OperatorAddress:         vi.OperatorAddress,
		ConsPubKey:              bechConsPubKey,
		Jailed:                  vi.Jailed,
		Status:                  vi.Status,
		DelegatorShares:         vi.DelegatorShares,
		Description:             vi.Description,
		UnbondingHeight:         vi.UnbondingHeight,
		UnbondingCompletionTime: vi.UnbondingCompletionTime,
		MinSelfDelegation:       vi.MinSelfDelegation,
	}, err
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

// NewDescription creates a new instance of Description
func NewDescription(moniker, identity, website, details string) Description {
	return Description{
		Moniker:  moniker,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
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

// Validator is the struct of standard validator's detail info
type Validator struct {
	OperatorAddress         sdk.ValAddress `json:"operator_address"`
	ConsPubKey              string         `json:"consensus_pubkey"`
	Jailed                  bool           `json:"jailed"`
	Status                  byte           `json:"status"`
	DelegatorShares         sdk.Dec        `json:"delegator_shares"`
	Description             Description    `json:"description"`
	UnbondingHeight         int64          `json:"unbonding_height"`
	UnbondingCompletionTime time.Time      `json:"unbonding_time"`
	MinSelfDelegation       sdk.Dec        `json:"min_self_delegation"`
}

// Description shows the detail info of a validator
type Description struct {
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Website  string `json:"website"`
	Details  string `json:"details"`
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
