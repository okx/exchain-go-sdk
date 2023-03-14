package types

import (
	"github.com/okx/okbchain/x/token"
	tokentypes "github.com/okx/okbchain/x/token/types"
)

// const
const (
	ModuleName = token.ModuleName
)

type (
	TransferUnit = tokentypes.TransferUnit
	TokenResp    = tokentypes.TokenResp
)
