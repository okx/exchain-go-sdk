package dex

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	dextypes "github.com/okex/okexchain/x/dex/types"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"testing"
)

const (
	addr      = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okexchainpub17weu6qepq0ph2t3u697qar7rmdtdtqp4744jcprjd2h356zr0yh5vmw38a3my4vqjx5"
	mnemonic  = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	memo      = "my memo"

	product = "btc-000_okt"
	recAddr = "okexchain193xnjknz3e52mqv2nyufnzjugu3mh65rpxdasn"
)

func TestDexClient_QueryProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient))

	initPrice, err := sdk.NewDecFromStr("10.24")
	require.NoError(t, err)
	minQuantity, err := sdk.NewDecFromStr("1.024")
	require.NoError(t, err)
	ownerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	deposit, err := sdk.ParseDecCoin("1024.1024okt")
	require.NoError(t, err)

	expectedRet := mockCli.BuildTokenPairsResponseBytes("btc", "eth", "okt",
		initPrice, minQuantity, 4, 4, 512, 1024, 2048, 4096,
		true, ownerAddr, deposit)
	expectedCdc := mockCli.GetCodec()
	expectedParams := expectedCdc.MustMarshalJSON(dextypes.NewQueryDexInfoParams(addr, dextypes.DefaultPage, dextypes.DefaultPerPage))
	expectedPath := fmt.Sprintf("custom/%s/%s", dextypes.QuerierRoute, dextypes.QueryProducts)
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(3)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	tokenPairs, err := mockCli.Dex().QueryProducts(addr, dextypes.DefaultPage, dextypes.DefaultPerPage)
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

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Dex().QueryProducts(addr, dextypes.DefaultPage, dextypes.DefaultPerPage)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Dex().QueryProducts(addr, dextypes.DefaultPage, dextypes.DefaultPerPage)
	require.Error(t, err)
}
