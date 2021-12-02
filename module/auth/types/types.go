package types

import (
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth"
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth/exported"
)

// const
const (
	ModuleName = auth.ModuleName
)

type (
	Account = exported.Account
)
