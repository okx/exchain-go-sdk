package gosdk

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module"
	"github.com/okex/okexchain-go-sdk/module/ammswap"
	"github.com/okex/okexchain-go-sdk/module/auth"
	authtypes "github.com/okex/okexchain-go-sdk/module/auth/types"
	"github.com/okex/okexchain-go-sdk/module/backend"
	"github.com/okex/okexchain-go-sdk/module/dex"
	"github.com/okex/okexchain-go-sdk/module/distribution"
	"github.com/okex/okexchain-go-sdk/module/governance"
	"github.com/okex/okexchain-go-sdk/module/order"
	"github.com/okex/okexchain-go-sdk/module/slashing"
	slashingtypes "github.com/okex/okexchain-go-sdk/module/slashing/types"
	"github.com/okex/okexchain-go-sdk/module/staking"
	"github.com/okex/okexchain-go-sdk/module/tendermint"
	"github.com/okex/okexchain-go-sdk/module/token"
	tokentypes "github.com/okex/okexchain-go-sdk/module/token/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// Client - structure of the main client of OKExChain GoSDK
type Client struct {
	config  gosdktypes.ClientConfig
	cdc     *codec.Codec
	modules map[string]gosdktypes.Module
}

// NewClient creates a new instance of Client
func NewClient(config gosdktypes.ClientConfig) Client {
	cdc := gosdktypes.NewCodec()
	pClient := &Client{
		config:  config,
		cdc:     cdc,
		modules: make(map[string]gosdktypes.Module),
	}
	pBaseClient := module.NewBaseClient(cdc, &pClient.config)

	pClient.registerModule(
		auth.NewAuthClient(pBaseClient),
		backend.NewBackendClient(pBaseClient),
		dex.NewDexClient(pBaseClient),
		distribution.NewDistrClient(pBaseClient),
		governance.NewGovClient(pBaseClient),
		order.NewOrderClient(pBaseClient),
		staking.NewStakingClient(pBaseClient),
		slashing.NewSlashingClient(pBaseClient),
		token.NewTokenClient(pBaseClient),
		tendermint.NewTendermintClient(pBaseClient),
		ammswap.NewAmmSwapClient(pBaseClient),
	)

	return *pClient
}

func (cli *Client) registerModule(mods ...gosdktypes.Module) {
	for _, mod := range mods {
		moduleName := mod.Name()
		if _, ok := cli.modules[moduleName]; ok {
			panic(fmt.Sprintf("duplicated module: %s", moduleName))
		}
		// register codec by each module
		mod.RegisterCodec(cli.cdc)
		cli.modules[moduleName] = mod
	}
	gosdktypes.RegisterBasicCodec(cli.cdc)
	cli.cdc.Seal()
}

// GetConfig returns the client config
func (cli *Client) GetConfig() gosdktypes.ClientConfig {
	return cli.config
}

// nolint
func (cli *Client) Auth() exposed.Auth {
	return cli.modules[authtypes.ModuleName].(exposed.Auth)
}
func (cli *Client) Backend() exposed.Backend {
	return cli.modules[backend.ModuleName].(exposed.Backend)
}
func (cli *Client) Dex() exposed.Dex {
	return cli.modules[dex.ModuleName].(exposed.Dex)
}
func (cli *Client) Distribution() exposed.Distribution {
	return cli.modules[distribution.ModuleName].(exposed.Distribution)
}
func (cli *Client) Governance() exposed.Governance {
	return cli.modules[governance.ModuleName].(exposed.Governance)
}
func (cli *Client) Order() exposed.Order {
	return cli.modules[order.ModuleName].(exposed.Order)
}
func (cli *Client) Staking() exposed.Staking {
	return cli.modules[staking.ModuleName].(exposed.Staking)
}
func (cli *Client) Slashing() exposed.Slashing {
	return cli.modules[slashingtypes.ModuleName].(exposed.Slashing)
}
func (cli *Client) Token() exposed.Token {
	return cli.modules[tokentypes.ModuleName].(exposed.Token)
}
func (cli *Client) Tendermint() exposed.Tendermint {
	return cli.modules[tendermint.ModuleName].(exposed.Tendermint)
}
func (cli *Client) AmmSwap() exposed.AmmSwap {
	return cli.modules[ammswap.ModuleName].(exposed.AmmSwap)
}
