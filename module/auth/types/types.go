package types

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

// const
const (
	ModuleName = auth.ModuleName
)

type (
	Account = exported.Account
)
