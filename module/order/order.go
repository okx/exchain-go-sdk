package order

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/order/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/order"
)

var _ gosdktypes.Module = (*orderClient)(nil)

type orderClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in order module
func (orderClient) RegisterCodec(cdc *codec.Codec) {
	order.RegisterCodec(cdc)
}

// Name returns the module name
func (orderClient) Name() string {
	return types.ModuleName
}

// NewOrderClient creates a new instance of order client as implement
func NewOrderClient(baseClient gosdktypes.BaseClient) exposed.Order {
	return orderClient{baseClient}
}
