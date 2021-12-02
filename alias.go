package gosdk

import (
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	ammswap "github.com/okex/exchain-go-sdk/module/ammswap/types"
	auth "github.com/okex/exchain-go-sdk/module/auth/types"
	backend "github.com/okex/exchain-go-sdk/module/backend/types"
	dex "github.com/okex/exchain-go-sdk/module/dex/types"
	evm "github.com/okex/exchain-go-sdk/module/evm/types"
	governance "github.com/okex/exchain-go-sdk/module/governance/types"
	order "github.com/okex/exchain-go-sdk/module/order/types"
	staking "github.com/okex/exchain-go-sdk/module/staking/types"
	tendermint "github.com/okex/exchain-go-sdk/module/tendermint/types"
	token "github.com/okex/exchain-go-sdk/module/token/types"
	"github.com/okex/exchain-go-sdk/types"
	farm "github.com/okex/exchain/x/farm/types"
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
	// ammswap
	SwapTokenPair = ammswap.SwapTokenPair
	// auth
	Account = auth.Account
	// staking
	Validator         = staking.Validator
	DelegatorResponse = staking.DelegatorResponse
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
	Transaction = backend.Transaction
	// tendermint
	Block            = tendermint.Block
	BlockResults     = tendermint.ResultBlockResults
	ResultCommit     = tendermint.ResultCommit
	ResultValidators = tendermint.ResultValidators
	ResultTx         = tendermint.ResultTx
	ResultTxSearch   = tendermint.ResultTxSearch
	// governance
	Proposal = governance.Proposal
	// farm
	FarmPool = farm.FarmPool
	LockInfo = farm.LockInfo
	// evm
	QueryResCode    = evm.QueryResCode
	QueryResStorage = evm.QueryResStorage
)
