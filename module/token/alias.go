package token

import "github.com/okex/okexchain-go-sdk/module/token/types"

// const
const (
	ModuleName = types.ModuleName
)

type (
	// nolint
	Token             = types.Token
	AccountTokensInfo = types.AccountTokensInfo
	TransferUnit      = types.TransferUnit
)

var (
	// NewTransferUnit is alias for NewTransferUnit function in types file
	NewTransferUnit = types.NewTransferUnit
)
