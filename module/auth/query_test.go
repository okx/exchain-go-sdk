package auth

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"testing"
)

const (
	addr         = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	accPubkey    = "okexchainpub17weu6qepq0ph2t3u697qar7rmdtdtqp4744jcprjd2h356zr0yh5vmw38a3my4vqjx5"
	mockCodeHash = "1234567890abcdef"
)

func TestAuthClient_QueryAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewAuthClient(mockCli.MockBaseClient))

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedCdc := mockCli.GetCodec()
	expectedRet := mockCli.BuildAccountBytes(addr, accPubkey, mockCodeHash, "1024btc,2048.1024okt", 1, 2)
	expectedPath := fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount)
	expectedParams, err := expectedCdc.MarshalJSON(auth.NewQueryAccountParams(accAddr))
	require.NoError(t, err)
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(6)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	acc, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	require.Equal(t, addr, acc.GetAddress().String())
	require.Equal(t, uint64(1), acc.GetAccountNumber())
	require.Equal(t, uint64(2), acc.GetSequence())
	accPkBech32, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, acc.GetPubKey())
	require.NoError(t, err)
	require.Equal(t, accPubkey, accPkBech32)

	coins := acc.GetCoins()
	expectedCoin0, err := sdk.ParseDecCoin("1024btc")
	require.NoError(t, err)
	expectedCoin1, err := sdk.ParseDecCoin("2048.1024okt")
	require.NoError(t, err)
	require.Equal(t, expectedCoin0, coins[0])
	require.Equal(t, expectedCoin1, coins[1])

	_, err = mockCli.Auth().QueryAccount(addr[1:])
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), nil)
	_, err = mockCli.Auth().QueryAccount(addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Auth().QueryAccount(addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Auth().QueryAccount(addr)
	require.Error(t, err)
}
