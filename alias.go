package gosdk

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/okex/okexchain-go-sdk/module/auth/types"
	"github.com/okex/okexchain-go-sdk/module/backend"
	"github.com/okex/okexchain-go-sdk/module/dex"
	"github.com/okex/okexchain-go-sdk/module/governance"
	"github.com/okex/okexchain-go-sdk/module/order"
	"github.com/okex/okexchain-go-sdk/module/tendermint"
	token "github.com/okex/okexchain-go-sdk/module/token/types"
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
	Account = auth.Account
	// staking
	//Validator     = staking.Validator
	//DelegatorResp = staking.DelegatorResp
	// token
	TokenResp = token.TokenResp
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
