package gosdk

import (
	"github.com/okex/okchain-go-sdk/module/auth"
	"github.com/okex/okchain-go-sdk/module/backend"
	"github.com/okex/okchain-go-sdk/module/dex"
	"github.com/okex/okchain-go-sdk/module/order"
	"github.com/okex/okchain-go-sdk/module/staking"
	"github.com/okex/okchain-go-sdk/module/tendermint"
	"github.com/okex/okchain-go-sdk/module/token"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	BroadcastSync  = sdk.BroadcastSync
	BroadcastAsync = sdk.BroadcastAsync
	BroadcastBlock = sdk.BroadcastBlock
)

var (
	// NewClientConfig gives an easy way for the callers to set client config
	NewClientConfig = sdk.NewClientConfig
)

type (
	TxResponse = sdk.TxResponse
	// auth
	Account = auth.Account
	// staking
	Validator = staking.Validator
	DelegatorResp = staking.DelegatorResp
	// token
	Token = token.Token
	AccountTokensInfo = token.AccountTokensInfo
	// dex
	TokenPair = dex.TokenPair
	// order
	BookRes = order.BookRes
	// backend
	Ticker = backend.Ticker
	MatchResult = backend.MatchResult
	Order = backend.Order
	Deal = backend.Deal
	// tendermint
	Block = tendermint.Block
)
