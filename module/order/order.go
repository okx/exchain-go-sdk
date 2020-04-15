package order

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/module/order/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

var _ sdk.Module = (*orderClient)(nil)

type orderClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in order module
func (orderClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (orderClient) Name() string {
	return types.ModuleName
}

// NewOrderClient creates a new instance of order client as implement
func NewOrderClient(baseClient sdk.BaseClient) exposed.Order {
	return orderClient{baseClient}
}
