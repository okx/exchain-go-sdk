package exposed

import (
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
	"time"
)

// Staking shows the expected behavior for inner staking client
type Staking interface {
	types.Module
	StakingTx
	StakingQuery
}

// StakingTx shows the expected tx behavior for inner staking client
type StakingTx interface {
	CreateValidator(fromInfo keys.Info, passWd, pubkeyStr, moniker, identity, website, details, memo string, accNum,
		seqNum uint64) (types.TxResponse, error)
	DestroyValidator(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (types.TxResponse, error)
	EditValidator(fromInfo keys.Info, passWd, moniker, identity, website, details, memo string, accNum, seqNum uint64) (
		types.TxResponse, error)
	Delegate(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error)
	Unbond(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error)
	Vote(fromInfo keys.Info, passWd string, valAddrsStr []string, memo string, accNum, seqNum uint64) (types.TxResponse, error)
	RegisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error)
	UnregisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error)
	BindProxy(fromInfo keys.Info, passWd, proxyAddrStr, memo string, accNum, seqNum uint64) (types.TxResponse, error)
}

// StakingQuery shows the expected query behavior for inner staking client
type StakingQuery interface {
	QueryValidators() ([]Validator, error)
	QueryValidator(valAddrStr string) (Validator, error)
	QueryDelegator(delAddrStr string) (DelegatorResp, error)
}

// Validator is the struct of exposed validator's detail info
type Validator struct {
	OperatorAddress         types.ValAddress `json:"operator_address"`
	ConsPubKey              string           `json:"consensus_pubkey"`
	Jailed                  bool             `json:"jailed"`
	Status                  byte             `json:"status"`
	DelegatorShares         types.Dec        `json:"delegator_shares"`
	Description             Description      `json:"description"`
	UnbondingHeight         int64            `json:"unbonding_height"`
	UnbondingCompletionTime time.Time        `json:"unbonding_time"`
	MinSelfDelegation       types.Dec        `json:"min_self_delegation"`
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
	DelegatorAddress     types.AccAddress   `json:"delegator_address"`
	ValidatorAddresses   []types.ValAddress `json:"validator_address"`
	Shares               types.Dec          `json:"shares"`
	Tokens               types.Dec          `json:"tokens" `
	UnbondedTokens       types.Dec          `json:"unbonded_tokens"`
	CompletionTime       time.Time          `json:"completion_time"`
	IsProxy              bool               `json:"is_proxy"`
	TotalDelegatedTokens types.Dec          `json:"total_delegated_tokens"`
	ProxyAddress         types.AccAddress   `json:"proxy_address"`
}
