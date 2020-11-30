package auth

import (
	"errors"
	gosdk "github.com/okex/okexchain-go-sdk"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/auth/types"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

const (
	addr      = "okexchain1kfs5q53jzgzkepqa6ual0z7f97wvxnkamr5vys"
	accPubkey = "okexchainpub1addwnpepq2vs59k5r76j4eazstu2e9dpttkr9enafdvnlhe27l2a88wpc0rsk0xy9zf"
)

func TestAuthClient_QueryAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdk.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	// TODO
	//mockCli.RegisterModule(NewAuthClient(mockCli.MockBaseClient))

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedCdc := mockCli.GetCodec()
	expectedRet := mockCli.BuildAccountBytes(addr, accPubkey, "1024btc,2048.1024okt", 1, 2)
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(types.AccountInfoPath, tmbytes.HexBytes(types.GetAddressStoreKey(accAddr))).Return(expectedRet, nil)

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

	_, err = mockCli.Auth().QueryAccount(addr[1:])
	require.Error(t, err)

	mockCli.EXPECT().Query(types.AccountInfoPath, tmbytes.HexBytes(types.GetAddressStoreKey(accAddr))).Return(nil, nil)
	_, err = mockCli.Auth().QueryAccount(addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.AccountInfoPath, tmbytes.HexBytes(types.GetAddressStoreKey(accAddr))).
		Return(nil, errors.New("default error"))
	_, err = mockCli.Auth().QueryAccount(addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.AccountInfoPath, tmbytes.HexBytes(types.GetAddressStoreKey(accAddr))).
		Return(expectedRet[1:], nil)
	_, err = mockCli.Auth().QueryAccount(addr)
	require.Error(t, err)
}
