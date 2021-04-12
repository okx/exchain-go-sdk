package distribution

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/distribution/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/x/distribution"
)

var _ gosdktypes.Module = (*distrClient)(nil)

type distrClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in distribution module
func (dc distrClient) RegisterCodec(cdc *codec.Codec) {
	distribution.RegisterCodec(cdc)
}

// Name returns the module name
func (distrClient) Name() string {
	return types.ModuleName
}

// NewDistrClient creates a new instance of distribution client as implement
func NewDistrClient(baseClient gosdktypes.BaseClient) exposed.Distribution {
	return distrClient{baseClient}
}
