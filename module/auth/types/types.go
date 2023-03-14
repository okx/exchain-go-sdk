package types

import (
	"github.com/okx/okbchain/libs/cosmos-sdk/x/auth"
	"github.com/okx/okbchain/libs/cosmos-sdk/x/auth/exported"
)

// const
const (
	ModuleName = auth.ModuleName
)

type (
	Account = exported.Account
)
