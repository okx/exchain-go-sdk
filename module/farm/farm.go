package farm

import (
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/farm/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	farm "github.com/okex/exchain/x/farm/types"
)

type farmClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in farm module
func (fc farmClient) RegisterCodec(cdc *codec.Codec) {
	farm.RegisterCodec(cdc)
}

// Name returns the module name
func (farmClient) Name() string {
	return types.ModuleName
}

// NewFarmClient creates a new instance of farm client as implement
func NewFarmClient(baseClient gosdktypes.BaseClient) exposed.Farm {
	return farmClient{baseClient}
}
