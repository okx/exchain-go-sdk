package gosdk

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/okex/okexchain-go-sdk/module/auth/types"
	"github.com/okex/okexchain-go-sdk/module/backend"
	"github.com/okex/okexchain-go-sdk/module/dex"
	"github.com/okex/okexchain-go-sdk/module/governance"
	"github.com/okex/okexchain-go-sdk/module/order"
	"github.com/okex/okexchain-go-sdk/module/staking"
	"github.com/okex/okexchain-go-sdk/module/tendermint"
	"github.com/okex/okexchain-go-sdk/module/token"
	"github.com/okex/okexchain-go-sdk/types"
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
	Account = types2.Account
	// staking
	Validator     = staking.Validator
	DelegatorResp = staking.DelegatorResp
	// token
	Token             = token.Token
	AccountTokensInfo = token.AccountTokensInfo
	// dex
	TokenPair = dex.TokenPair
	// order
	BookRes     = order.BookRes
	OrderDetail = order.OrderDetail
	// backend
	Ticker      = backend.Ticker
	MatchResult = backend.MatchResult
	Order       = backend.Order
	Deal        = backend.Deal
	// tendermint
	Block            = tendermint.Block
	BlockResults     = tendermint.BlockResults
	ResultCommit     = tendermint.ResultCommit
	ResultValidators = tendermint.ResultValidators
	ResultTx         = tendermint.ResultTx
	ResultTxs        = tendermint.ResultTxs
	// governance
	Proposal = governance.Proposal
)
