package token

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/token/types"
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
	recAddr   = "okexchain16zgvph7qc3n4jvamq0lkv3y37k0hc5pw9hhhrs"

	tokenSymbol = "btc-000"
)

func TestTokenClient_QueryAccountTokensInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient))

	expectedRet := mockCli.BuildAccountTokensInfoBytes(addr, tokenSymbol, "1024.1024", "2048,2048", "10.24")
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryAccTokenParams("", "all")
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.AccountTokensInfoPath, addr), tmbytes.HexBytes(queryBytes)).Return(expectedRet, nil)

	accTokensInfo, err := mockCli.Token().QueryAccountTokensInfo(addr)
	require.NoError(t, err)

	require.Equal(t, addr, accTokensInfo.Address)
	require.Equal(t, tokenSymbol, accTokensInfo.Currencies[0].Symbol)
	require.Equal(t, "1024.1024", accTokensInfo.Currencies[0].Available)
	require.Equal(t, "2048,2048", accTokensInfo.Currencies[0].Freeze)
	require.Equal(t, "10.24", accTokensInfo.Currencies[0].Locked)

	_, err = mockCli.Token().QueryAccountTokensInfo(addr[1:])
	require.Error(t, err)

	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.AccountTokensInfoPath, addr), tmbytes.HexBytes(queryBytes)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Token().QueryAccountTokensInfo(addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.AccountTokensInfoPath, addr), tmbytes.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	_, err = mockCli.Token().QueryAccountTokensInfo(addr)
	require.Error(t, err)
}

func TestTokenClient_QueryAccountTokenInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient))

	expectedRet := mockCli.BuildAccountTokensInfoBytes(addr, tokenSymbol, "1024.1024", "2048,2048", "10.24")
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryAccTokenParams(tokenSymbol, "partial")
	require.NoError(t, err)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.AccountTokensInfoPath, addr), tmbytes.HexBytes(queryBytes)).Return(expectedRet, nil)

	accTokensInfo, err := mockCli.Token().QueryAccountTokenInfo(addr, tokenSymbol)
	require.NoError(t, err)

	require.Equal(t, addr, accTokensInfo.Address)
	require.Equal(t, tokenSymbol, accTokensInfo.Currencies[0].Symbol)
	require.Equal(t, "1024.1024", accTokensInfo.Currencies[0].Available)
	require.Equal(t, "2048,2048", accTokensInfo.Currencies[0].Freeze)
	require.Equal(t, "10.24", accTokensInfo.Currencies[0].Locked)

	_, err = mockCli.Token().QueryAccountTokenInfo(addr[1:], tokenSymbol)
	require.Error(t, err)

	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.AccountTokensInfoPath, addr), tmbytes.HexBytes(queryBytes)).
		Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Token().QueryAccountTokenInfo(addr, tokenSymbol)
	require.Error(t, err)

	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.AccountTokensInfoPath, addr), tmbytes.HexBytes(queryBytes)).
		Return(expectedRet[1:], nil)
	_, err = mockCli.Token().QueryAccountTokenInfo(addr, tokenSymbol)
	require.Error(t, err)

}

func TestTokenClient_QueryTokenInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
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

	mockCli.EXPECT().Query(fmt.Sprintf("custom/%s/info/%s", types.ModuleName, tokenSymbol), nil).
		Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Token().QueryTokenInfo("", tokenSymbol)
	require.Error(t, err)

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

	mockCli.EXPECT().Query(fmt.Sprintf("custom/%s/tokens/%s", types.ModuleName, addr), nil).
		Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Token().QueryTokenInfo(addr, "")
	require.Error(t, err)
}
