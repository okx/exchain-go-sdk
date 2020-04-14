package sdk

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/types"
)

// const
const (
	BroadcastSync  = types.BroadcastSync
	BroadcastAsync = types.BroadcastAsync
	BroadcastBlock = types.BroadcastBlock
)

var (
	NewClientConfig = types.NewClientConfig
)

type (
	Validator = exposed.Validator
	DelegatorResp = exposed.DelegatorResp
)
