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
	recAddr = "okchain1wux20ku36ntgtxpgm7my9863xy3fqs0xgh66d7"

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

func TestTokenClient_QueryTokenInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient))

	originalTotalSupply, err := sdk.NewDecFromStr("10000000000")
	require.NoError(t, err)
	totalSupply, err := sdk.NewDecFromStr("20000000000")
	require.NoError(t, err)
	ownerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedRet := mockCli.BuildTokenInfoBytes("default description", tokenSymbol, "default original symbol",
		"default whole name", originalTotalSupply, totalSupply, ownerAddr, true, false)
	expectedCdc := mockCli.GetCodec()

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(fmt.Sprintf("custom/%s/info/%s", types.ModuleName, tokenSymbol), nil).Return(expectedRet, nil)

	tokensInfo, err := mockCli.Token().QueryTokenInfo("", tokenSymbol)
	require.NoError(t, err)

	require.Equal(t, ownerAddr, tokensInfo[0].Owner)
	require.Equal(t, "default description", tokensInfo[0].Description)
	require.Equal(t, tokenSymbol, tokensInfo[0].Symbol)
	require.Equal(t, "default original symbol", tokensInfo[0].OriginalSymbol)
	require.Equal(t, "default whole name", tokensInfo[0].WholeName)
	require.Equal(t, originalTotalSupply, tokensInfo[0].OriginalTotalSupply)
	require.Equal(t, totalSupply, tokensInfo[0].TotalSupply)
	require.Equal(t, true, tokensInfo[0].Mintable)
	require.Equal(t, ownerAddr, tokensInfo[0].Owner)

	expectedRet = mockCli.BuildTokenInfoBytes("default description", tokenSymbol, "default original symbol",
		"default whole name", originalTotalSupply, totalSupply, ownerAddr, true, true)

	mockCli.EXPECT().Query(fmt.Sprintf("custom/%s/tokens/%s", types.ModuleName, addr), nil).Return(expectedRet, nil)

	tokensInfo, err = mockCli.Token().QueryTokenInfo(addr, "")
	require.NoError(t, err)

	require.Equal(t, ownerAddr, tokensInfo[0].Owner)
	require.Equal(t, "default description", tokensInfo[0].Description)
	require.Equal(t, tokenSymbol, tokensInfo[0].Symbol)
	require.Equal(t, "default original symbol", tokensInfo[0].OriginalSymbol)
	require.Equal(t, "default whole name", tokensInfo[0].WholeName)
	require.Equal(t, originalTotalSupply, tokensInfo[0].OriginalTotalSupply)
	require.Equal(t, totalSupply, tokensInfo[0].TotalSupply)
	require.Equal(t, true, tokensInfo[0].Mintable)
	require.Equal(t, ownerAddr, tokensInfo[0].Owner)

	_, err = mockCli.Token().QueryTokenInfo("", "")
	require.Error(t, err)
}
