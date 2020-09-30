package farm

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/farm/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
)

type farmClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in farm module
func (fc farmClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (farmClient) Name() string {
	return types.ModuleName
}

// NewFarmClient creates a new instance of farm client as implement
func NewFarmClient(baseClient sdk.BaseClient) exposed.Farm {
	return farmClient{baseClient}
}