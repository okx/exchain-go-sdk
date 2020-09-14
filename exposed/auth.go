package exposed

import (
	"github.com/okex/okexchain-go-sdk/module/auth/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
)

// Auth shows the expected behavior for inner auth client
type Auth interface {
	sdk.Module
	AuthQuery
}

// AuthQuery shows the expected query behavior for inner auth client
type AuthQuery interface {
	QueryAccount(accAddrStr string) (types.Account, error)
}
