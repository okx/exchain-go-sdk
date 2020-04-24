package dex

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/dex/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/params"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"
)

const (
	addr      = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okchainpub1addwnpepqgzuks5c07kfce85e0t0x8qkuvvxu874965ruafn6svhjrhswt0lgdj85lv"
	mnemonic  = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	memo      = "my memo"

	product = "btc-000_okt"

	recMnemonic = "pepper basket run install fury scheme journey worry tumble toddler swap change"
	recAddr     = "okchain1wux20ku36ntgtxpgm7my9863xy3fqs0xgh66d7"
)

func TestDexClient_QueryProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
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

	expectedRet := mockCli.BuildTokenPairsBytes("btc", "eth", "okt",
		initPrice, minQuantity, 4, 4, 512, 1024, 2048, 4096,
		true, ownerAddr, deposit)
	expectedCdc := mockCli.GetCodec()

	queryParams, err := params.NewQueryDexInfoParams(addr, 1, 30)
	require.NoError(t, err)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(types.ProductsPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

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

	mockCli.EXPECT().Query(types.ProductsPath, cmn.HexBytes(queryBytes)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Dex().QueryProducts(addr, 1, 30)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.ProductsPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	_, err = mockCli.Dex().QueryProducts(addr, 1, 30)
	require.Error(t, err)

}
