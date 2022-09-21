package wasm

import (
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/token/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/client/context"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain/x/token"
	wasm "github.com/okex/exchain/x/wasm/types"
)

var _ gosdktypes.Module = (*wasmClient)(nil)

type wasmClient struct {
	gosdktypes.BaseClient

	wasm.QueryClient
}

// RegisterCodec registers the msg type in token module
func (wasmClient) RegisterCodec(cdc *codec.Codec) {
	token.RegisterCodec(cdc)
}

// Name returns the module name
func (wasmClient) Name() string {
	return types.ModuleName
}

// NewTokenClient creates a new instance of token client as implement
func NewWasmClient(baseClient gosdktypes.BaseClient) exposed.Wasm {
	clientCtx := context.NewCLIContext().WithNodeURI(baseClient.GetConfig().NodeURI)
	return wasmClient{baseClient, wasm.NewQueryClient(clientCtx)}
}
