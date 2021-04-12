package types

import (
	stakingcli "github.com/okex/exchain/x/staking/client/cli"
	stakingtypes "github.com/okex/exchain/x/staking/types"
)

// const
const (
	ModuleName = stakingtypes.ModuleName
)

type (
	Validator         = stakingtypes.Validator
	DelegatorResponse = stakingcli.DelegatorResponse
)
