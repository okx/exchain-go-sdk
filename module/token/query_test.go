package token

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/token/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/params"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"

	"testing"
)

const (
	addr      = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okchainpub1addwnpepqgzuks5c07kfce85e0t0x8qkuvvxu874965ruafn6svhjrhswt0lgdj85lv"
	mnemonic  = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	memo      = "my memo"

	tokenSymbol = "btc-000"
)

func TestTokenClient_QueryAccountTokensInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient))

	expectedRet := mockCli.BuildAccountTokensInfoBytes(addr, tokenSymbol, "1024.1024", "2048,2048", "10.24")
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryAccTokenParams("", "all")
	require.NoError(t, err)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.AccountTokensInfoPath, addr), cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	accTokensInfo, err := mockCli.Token().QueryAccountTokensInfo(addr)
	require.NoError(t, err)

	require.Equal(t, addr, accTokensInfo.Address)
	require.Equal(t, tokenSymbol, accTokensInfo.Currencies[0].Symbol)
	require.Equal(t, "1024.1024", accTokensInfo.Currencies[0].Available)
	require.Equal(t, "2048,2048", accTokensInfo.Currencies[0].Freeze)
	require.Equal(t, "10.24", accTokensInfo.Currencies[0].Locked)

	_, err = mockCli.Token().QueryAccountTokensInfo(addr[1:])
	require.Error(t, err)

}

func TestTokenClient_QueryAccountTokenInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient))

	expectedRet := mockCli.BuildAccountTokensInfoBytes(addr, tokenSymbol, "1024.1024", "2048,2048", "10.24")
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryAccTokenParams(tokenSymbol, "partial")
	require.NoError(t, err)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.AccountTokensInfoPath, addr), cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	accTokensInfo, err := mockCli.Token().QueryAccountTokenInfo(addr, tokenSymbol)
	require.NoError(t, err)

	require.Equal(t, addr, accTokensInfo.Address)
	require.Equal(t, tokenSymbol, accTokensInfo.Currencies[0].Symbol)
	require.Equal(t, "1024.1024", accTokensInfo.Currencies[0].Available)
	require.Equal(t, "2048,2048", accTokensInfo.Currencies[0].Freeze)
	require.Equal(t, "10.24", accTokensInfo.Currencies[0].Locked)

	_, err = mockCli.Token().QueryAccountTokenInfo(addr[1:], tokenSymbol)
	require.Error(t, err)
}
