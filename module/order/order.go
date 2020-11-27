package order

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/order/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

var _ gosdktypes.Module = (*orderClient)(nil)

type orderClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in order module
func (orderClient) RegisterCodec(cdc gosdktypes.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (orderClient) Name() string {
	return types.ModuleName
}

// NewOrderClient creates a new instance of order client as implement
func NewOrderClient(baseClient gosdktypes.BaseClient) exposed.Order {
	return orderClient{baseClient}
}
