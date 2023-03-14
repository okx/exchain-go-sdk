package gosdk

import (
	auth "github.com/okex/exchain-go-sdk/module/auth/types"
	evm "github.com/okex/exchain-go-sdk/module/evm/types"
	governance "github.com/okex/exchain-go-sdk/module/governance/types"
	staking "github.com/okex/exchain-go-sdk/module/staking/types"
	tendermint "github.com/okex/exchain-go-sdk/module/tendermint/types"
	token "github.com/okex/exchain-go-sdk/module/token/types"
	"github.com/okex/exchain-go-sdk/types"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
)

// const
const (
	BroadcastSync  = types.BroadcastSync
	BroadcastAsync = types.BroadcastAsync
	BroadcastBlock = types.BroadcastBlock

	// vote for the proposal
	VoteYes        = "yes"
	VoteAbstain    = "abstain"
	VoteNo         = "no"
	VoteNoWithVeto = "no_with_veto"
)

var (
	// NewClientConfig gives an easy way for the callers to set client config
	NewClientConfig = types.NewClientConfig
)

// nolint
type (
	TxResponse = sdk.TxResponse
	// auth
	Account = auth.Account
	// staking
	Validator         = staking.Validator
	DelegatorResponse = staking.DelegatorResponse
	// token
	TokenResp = token.TokenResp
	// tendermint
	Block            = tendermint.Block
	BlockResults     = tendermint.ResultBlockResults
	ResultCommit     = tendermint.ResultCommit
	ResultValidators = tendermint.ResultValidators
	ResultTx         = tendermint.ResultTx
	ResultTxSearch   = tendermint.ResultTxSearch
	// governance
	Proposal = governance.Proposal
	// evm
	QueryResCode    = evm.QueryResCode
	QueryResStorage = evm.QueryResStorage
)
