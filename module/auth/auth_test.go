package auth

import (
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/auth/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"
	"testing"
)

const (
	addr      = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	accPubkey = "okchainpub1addwnpepqgzuks5c07kfce85e0t0x8qkuvvxu874965ruafn6svhjrhswt0lgdj85lv"
)

func TestAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewAuthClient(mockCli.MockBaseClient))

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedCdc := mockCli.GetCodec()
	expectedRet := mockCli.BuildAccountBytes(addr, accPubkey, "1024btc,2048.1024okt", 1, 2)
	mockCli.MockBaseClient.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.MockBaseClient.EXPECT().Query(types.AccountInfoPath, cmn.HexBytes(types.GetAddressStoreKey(accAddr))).Return(expectedRet, nil)

	acc, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	require.Equal(t, addr, acc.GetAddress().String())
	require.Equal(t, uint64(1), acc.GetAccountNumber())
	require.Equal(t, uint64(2), acc.GetSequence())
	accPkBech32, err := sdk.Bech32ifyAccPub(acc.GetPubKey())
	require.NoError(t, err)
	require.Equal(t, accPubkey, accPkBech32)

	coins := acc.GetCoins()
	expectedCoin0, err := sdk.ParseDecCoin("1024btc")
	require.NoError(t, err)
	expectedCoin1, err := sdk.ParseDecCoin("2048.1024okt")
	require.NoError(t, err)
	require.Equal(t, expectedCoin0, coins[0])
	require.Equal(t, expectedCoin1, coins[1])

}
