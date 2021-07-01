package evm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/evm/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	evm "github.com/okex/exchain/x/evm/types"
)

var _ gosdktypes.Module = (*evmClient)(nil)

type evmClient struct {
	bc gosdktypes.BaseClient
	ec *ethclient.Client
}

// RegisterCodec registers the msg type in evm module
func (ec evmClient) RegisterCodec(cdc *codec.Codec) {
	evm.RegisterCodec(cdc)
}

// Name returns the module name
func (evmClient) Name() string {
	return types.ModuleName
}

// NewEvmClient creates a new instance of evm client as implement
func NewEvmClient(baseClient gosdktypes.BaseClient) exposed.Evm {
	var client, err = ethclient.Dial(baseClient.GetConfig().NodeURI)
	if err != nil {
		client = nil
	}
	return evmClient{baseClient, client}
}