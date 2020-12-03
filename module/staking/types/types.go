package types

import (
	stakingcli "github.com/okex/okexchain/x/staking/client/cli"
	stakingtypes "github.com/okex/okexchain/x/staking/types"
)

// const
const (
	ModuleName = stakingtypes.ModuleName
)

type (
	Validator         = stakingtypes.Validator
	DelegatorResponse = stakingcli.DelegatorResponse
)
