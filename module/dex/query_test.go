package dex

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/dex/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"testing"
)

const (
	addr      = "okexchain1kfs5q53jzgzkepqa6ual0z7f97wvxnkamr5vys"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okexchainpub1addwnpepq2vs59k5r76j4eazstu2e9dpttkr9enafdvnlhe27l2a88wpc0rsk0xy9zf"
	mnemonic  = "view acid farm come spike since hour width casino cause mom sheriff"
	memo      = "my memo"

	product = "btc-000_okt"

	recMnemonic = "length borrow act busy blur mouse salad suspect demise dizzy obey rookie"
	recAddr     = "okexchain16zgvph7qc3n4jvamq0lkv3y37k0hc5pw9hhhrs"
)

func TestDexClient_QueryProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	// TODO
	//mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient))

	initPrice, err := sdk.NewDecFromStr("10.24")
	require.NoError(t, err)
	minQuantity, err := sdk.NewDecFromStr("1.024")
	require.NoError(t, err)
	ownerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	deposit, err := sdk.ParseDecCoin("1024.1024okt")
	require.NoError(t, err)

	expectedRet := mockCli.BuildTokenPairsBytes("btc", "eth", "okt",
		initPrice, minQuantity, 4, 4, 512, 1024, 2048, 4096,
		true, ownerAddr, deposit)
	expectedCdc := mockCli.GetCodec()

	queryParams, err := params.NewQueryDexInfoParams(addr, 1, 30)
	require.NoError(t, err)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(types.ProductsPath, tmbytes.HexBytes(queryBytes)).Return(expectedRet, nil)

	tokenPairs, err := mockCli.Dex().QueryProducts(addr, 1, 30)
	require.NoError(t, err)

	require.Equal(t, 2, len(tokenPairs))
	require.Equal(t, "btc", tokenPairs[0].BaseAssetSymbol)
	require.Equal(t, "eth", tokenPairs[1].BaseAssetSymbol)
	require.Equal(t, "okt", tokenPairs[0].QuoteAssetSymbol)
	require.Equal(t, initPrice, tokenPairs[0].InitPrice)
	require.Equal(t, int64(4), tokenPairs[0].MaxQuantityDigit)
	require.Equal(t, int64(4), tokenPairs[0].MaxQuantityDigit)
	require.Equal(t, uint64(2048), tokenPairs[0].ID)
	require.Equal(t, uint64(4096), tokenPairs[1].ID)
	require.Equal(t, true, tokenPairs[0].Delisting)
	require.Equal(t, ownerAddr, tokenPairs[0].Owner)
	require.Equal(t, deposit, tokenPairs[0].Deposits)
	require.Equal(t, int64(512), tokenPairs[0].BlockHeight)
	require.Equal(t, int64(1024), tokenPairs[1].BlockHeight)

	_, err = mockCli.Dex().QueryProducts(addr, -1, 30)
	require.Error(t, err)

	_, err = mockCli.Dex().QueryProducts(addr, 1, -30)
	require.Error(t, err)

	_, err = mockCli.Dex().QueryProducts(addr[1:], 1, 30)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.ProductsPath, tmbytes.HexBytes(queryBytes)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Dex().QueryProducts(addr, 1, 30)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.ProductsPath, tmbytes.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	_, err = mockCli.Dex().QueryProducts(addr, 1, 30)
	require.Error(t, err)

}
