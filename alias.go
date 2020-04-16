package sdk

import (
	"github.com/okex/okchain-go-sdk/module/auth"
	"github.com/okex/okchain-go-sdk/module/dex"
	"github.com/okex/okchain-go-sdk/module/staking"
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

	Account = auth.Account

	Validator = staking.Validator
	DelegatorResp = staking.DelegatorResp

	Token = token.Token
	AccountTokensInfo = token.AccountTokensInfo

	TokenPair = dex.TokenPair
)
