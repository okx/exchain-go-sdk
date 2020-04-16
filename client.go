package sdk

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/module"
	"github.com/okex/okchain-go-sdk/module/auth"
	"github.com/okex/okchain-go-sdk/module/backend"
	"github.com/okex/okchain-go-sdk/module/dex"
	"github.com/okex/okchain-go-sdk/module/order"
	"github.com/okex/okchain-go-sdk/module/slashing"
	"github.com/okex/okchain-go-sdk/module/staking"
	"github.com/okex/okchain-go-sdk/module/token"
	"github.com/okex/okchain-go-sdk/types"
)

// Client - structure of the main client of okchain gosdk
type Client struct {
	cdc     types.SDKCodec
	modules map[string]types.Module
}

// NewClient creates a new instance of Client
func NewClient(config types.ClientConfig) Client {
	cdc := types.NewCodec()
	pBaseClient := module.NewBaseClient(cdc, config)

	pClient := &Client{
		cdc:     cdc,
		modules: make(map[string]types.Module),
	}

	pClient.registerModule(
		auth.NewAuthClient(pBaseClient),
		backend.NewBackendClient(pBaseClient),
		dex.NewDexClient(pBaseClient),
		order.NewOrderClient(pBaseClient),
		staking.NewStakingClient(pBaseClient),
		slashing.NewSlashingClient(pBaseClient),
		token.NewTokenClient(pBaseClient),
	)

	return *pClient
}

func (cli *Client) registerModule(mods ...types.Module) {
	for _, mod := range mods {
		moduleName := mod.Name()
		if _, ok := cli.modules[mod.Name()]; ok {
			panic(fmt.Sprintf("duplicated module: %s", moduleName))
		}
		// register codec by each module
		mod.RegisterCodec(cli.cdc)
		cli.modules[moduleName] = mod
	}
	types.RegisterBasicCodec(cli.cdc)
}

// nolint
func (cli *Client) Auth() exposed.Auth {
	return cli.modules[auth.ModuleName].(exposed.Auth)
}
func (cli *Client) Backend() exposed.Backend {
	return cli.modules[backend.ModuleName].(exposed.Backend)
}
func (cli *Client) Dex() exposed.Dex {
	return cli.modules[dex.ModuleName].(exposed.Dex)
}
func (cli *Client) Order() exposed.Order {
	return cli.modules[order.ModuleName].(exposed.Order)
}
func (cli *Client) Staking() exposed.Staking {
	return cli.modules[staking.ModuleName].(exposed.Staking)
}
func (cli *Client) Slashing() exposed.Slashing {
	return cli.modules[slashing.ModuleName].(exposed.Slashing)
}
func (cli *Client) Token() exposed.Token {
	return cli.modules[token.ModuleName].(exposed.Token)
}
