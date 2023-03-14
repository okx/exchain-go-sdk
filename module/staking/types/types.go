package types

import (
	stakingcli "github.com/okx/okbchain/x/staking/client/cli"
	stakingtypes "github.com/okx/okbchain/x/staking/types"
)

// const
const (
	ModuleName = stakingtypes.ModuleName
)

type (
	Validator         = stakingtypes.Validator
	DelegatorResponse = stakingcli.DelegatorResponse
)
