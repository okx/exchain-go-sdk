package mocks

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/exposed"
	auth "github.com/okex/okchain-go-sdk/module/auth/types"
	backend "github.com/okex/okchain-go-sdk/module/backend/types"
	dex "github.com/okex/okchain-go-sdk/module/dex/types"
	order "github.com/okex/okchain-go-sdk/module/order/types"
	slashing "github.com/okex/okchain-go-sdk/module/slashing/types"
	staking "github.com/okex/okchain-go-sdk/module/staking/types"
	tendermint "github.com/okex/okchain-go-sdk/module/tendermint/types"
	token "github.com/okex/okchain-go-sdk/module/token/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

// MockClient - structure of the mock client for gosdk testing
type MockClient struct {
	t *testing.T
	*sdk.MockBaseClient
	config  sdk.ClientConfig
	cdc     sdk.SDKCodec
	modules map[string]sdk.Module
}

// NewMockClient creates a new instance of MockClient
func NewMockClient(t *testing.T, ctrl *gomock.Controller, config sdk.ClientConfig) MockClient {
	cdc := sdk.NewCodec()
	pMockClient := &MockClient{
		t:              t,
		MockBaseClient: sdk.NewMockBaseClient(ctrl),
		config:         config,
		cdc:            cdc,
		modules:        make(map[string]sdk.Module),
	}

	return *pMockClient
}

// RegisterModule registers the specific module for MockClient
func (mc *MockClient) RegisterModule(mods ...sdk.Module) {
	for _, mod := range mods {
		moduleName := mod.Name()
		if _, ok := mc.modules[moduleName]; ok {
			panic(fmt.Sprintf("duplicated module: %s", moduleName))
		}
		// register codec by each module
		mod.RegisterCodec(mc.cdc)
		mc.modules[moduleName] = mod
	}
	sdk.RegisterBasicCodec(mc.cdc)
	mc.cdc.Seal()
}

// GetConfig returns the client config
func (mc *MockClient) GetConfig() sdk.ClientConfig {
	return mc.config
}

// GetCodec returns the client codec
func (mc *MockClient) GetCodec() sdk.SDKCodec {
	return mc.cdc
}

// nolint
func (mc *MockClient) Auth() exposed.Auth {
	return mc.modules[auth.ModuleName].(exposed.Auth)
}
func (mc *MockClient) Backend() exposed.Backend {
	return mc.modules[backend.ModuleName].(exposed.Backend)
}
func (mc *MockClient) Dex() exposed.Dex {
	return mc.modules[dex.ModuleName].(exposed.Dex)
}
func (mc *MockClient) Order() exposed.Order {
	return mc.modules[order.ModuleName].(exposed.Order)
}
func (mc *MockClient) Staking() exposed.Staking {
	return mc.modules[staking.ModuleName].(exposed.Staking)
}
func (mc *MockClient) Slashing() exposed.Slashing {
	return mc.modules[slashing.ModuleName].(exposed.Slashing)
}
func (mc *MockClient) Token() exposed.Token {
	return mc.modules[token.ModuleName].(exposed.Token)
}
func (mc *MockClient) Tendermint() exposed.Tendermint {
	return mc.modules[tendermint.ModuleName].(exposed.Tendermint)
}

// BuildAccountBytes generates the account bytes for test
func (mc *MockClient) BuildAccountBytes(accAddrStr, accPubkeyStr, coinsStr string, accNum, seqNum uint64) []byte {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	require.NoError(mc.t, err)
	accPubkey, err := sdk.GetAccPubKeyBech32(accPubkeyStr)
	require.NoError(mc.t, err)
	coins, err := sdk.ParseDecCoins(coinsStr)
	require.NoError(mc.t, err)
	account := auth.BaseAccount{
		Address:       accAddr,
		Coins:         coins,
		PubKey:        accPubkey,
		AccountNumber: accNum,
		Sequence:      seqNum,
	}

	bytes, err := mc.cdc.MarshalBinaryBare(account)
	require.NoError(mc.t, err)

	return bytes
}

// BuildTokenPairBytes generates the token pairs bytes for test
func (mc *MockClient) BuildTokenPairsBytes(baseAssetSymbol1, baseAssetSymbol2, quoteAssetSymbol string, initPrice,
	minQuantity sdk.Dec, maxPriceDigit, maxQuantityDigit, blockHeight1, blockHeight2 int64, ID1, ID2 uint64, delisting bool,
	owner sdk.AccAddress, deposits sdk.DecCoin) []byte {

	var tokenPairs []dex.TokenPair

	tokenPairs = append(tokenPairs, dex.TokenPair{
		BaseAssetSymbol:  baseAssetSymbol1,
		QuoteAssetSymbol: quoteAssetSymbol,
		InitPrice:        initPrice,
		MaxPriceDigit:    maxPriceDigit,
		MaxQuantityDigit: maxQuantityDigit,
		MinQuantity:      minQuantity,
		ID:               ID1,
		Delisting:        delisting,
		Owner:            owner,
		Deposits:         deposits,
		BlockHeight:      blockHeight1,
	})

	tokenPairs = append(tokenPairs, dex.TokenPair{
		BaseAssetSymbol:  baseAssetSymbol2,
		QuoteAssetSymbol: quoteAssetSymbol,
		InitPrice:        initPrice,
		MaxPriceDigit:    maxPriceDigit,
		MaxQuantityDigit: maxQuantityDigit,
		MinQuantity:      minQuantity,
		ID:               ID2,
		Delisting:        delisting,
		Owner:            owner,
		Deposits:         deposits,
		BlockHeight:      blockHeight2,
	})

	return mc.cdc.MustMarshalJSON(tokenPairs)
}

// BuildOrderDetailBytes generates the order detail bytes for test
func (mc *MockClient) BuildOrderDetailBytes(txHash, orderID, extraInfo, product, side string, status, timestamp,
	orderExpireBlocks int64, sender sdk.AccAddress, price, quantity, filledAvgPrice, remainQuantity, remainLocked sdk.Dec,
	feePerBlock sdk.DecCoin) []byte {
	orderDetail := order.OrderDetail{
		TxHash:            txHash,
		OrderID:           orderID,
		Sender:            sender,
		Product:           product,
		Side:              side,
		Price:             price,
		Quantity:          quantity,
		Status:            status,
		FilledAvgPrice:    filledAvgPrice,
		RemainQuantity:    remainQuantity,
		RemainLocked:      remainLocked,
		Timestamp:         timestamp,
		OrderExpireBlocks: orderExpireBlocks,
		FeePerBlock:       feePerBlock,
		ExtraInfo:         extraInfo,
	}

	return mc.cdc.MustMarshalJSON(orderDetail)
}
