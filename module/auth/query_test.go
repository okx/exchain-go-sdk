package auth

import (
	"errors"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

const (
	addr      = "ex1qj5c07sm6jetjz8f509qtrxgh4psxkv3ddyq7u"
	accPubkey = "expub17weu6qepqtfc6zq8dukwc3lhlhx7th2csfjw0g3cqnqvanh7z9c2nhkr8mn5z9uq4q6"
)

func TestAuthClient_QueryAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewAuthClient(mockCli.MockBaseClient))

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedCdc := mockCli.GetCodec()
	expectedRet := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024btc,2048.1024okt", 1, 2)
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
