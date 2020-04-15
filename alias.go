package sdk

import (
	"github.com/okex/okchain-go-sdk/exposed"
	authtypes "github.com/okex/okchain-go-sdk/module/auth/types"
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
	Account = authtypes.Account
	Validator = exposed.Validator
	DelegatorResp = exposed.DelegatorResp
)
