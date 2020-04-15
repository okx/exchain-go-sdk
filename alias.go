package sdk

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/module/auth"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	BroadcastSync  = sdk.BroadcastSync
	BroadcastAsync = sdk.BroadcastAsync
	BroadcastBlock = sdk.BroadcastBlock
)

var (
	NewClientConfig = sdk.NewClientConfig
)

type (
	Account = auth.Account
	Validator = exposed.Validator
	DelegatorResp = exposed.DelegatorResp
)
