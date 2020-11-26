package farm

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/farm/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"
	"testing"
)

const (
	addr      = "okexchain1kfs5q53jzgzkepqa6ual0z7f97wvxnkamr5vys"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okexchainpub1addwnpepq2vs59k5r76j4eazstu2e9dpttkr9enafdvnlhe27l2a88wpc0rsk0xy9zf"
	mnemonic  = "view acid farm come spike since hour width casino cause mom sheriff"
	memo      = "my memo"

	expectedTokenSymbol                    = "abc-d53"
	expectedTokenAmount             int64  = 1024
	expectedAmountYieldPerBlock     int64  = 50
	expectedStartBlockHeightToYield int64  = 1000
	expectedPoolName                       = "default-pool-name"
	expectedHeight                  int64  = 1024
	expectedReferencePeriod         uint64 = 1
)

func TestFarmClient_QueryPools(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	expectedPoolName0, expectedPoolName1 := fmt.Sprintf("%s%d", expectedPoolName, 0), fmt.Sprintf("%s%d", expectedPoolName, 1)
	expectedDec := sdk.NewDec(expectedTokenAmount)
	expectedRet := mockCli.BuildFarmPoolsBytes(
		expectedPoolName0,
		expectedPoolName1,
		addr,
		expectedTokenSymbol,
		expectedStartBlockHeightToYield,
		expectedDec,
	)
	expectedCdc := mockCli.GetCodec()

	expectedOwnerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	// fixed to all pools query
	queryParams := params.NewQueryPoolsParams(1, 0)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(types.QueryPoolsPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	pools, err := mockCli.Farm().QueryPools()
	require.NoError(t, err)

	require.Equal(t, 2, len(pools))
	for i, pool := range pools {
		require.Equal(t, fmt.Sprintf("%s%d", expectedPoolName, i), pool.Name)
		require.Equal(t, expectedOwnerAddr, pool.Owner)
		require.Equal(t, expectedTokenSymbol, pool.MinLockAmount.Denom)
		require.True(t, pool.MinLockAmount.Amount.Equal(expectedDec))
		require.Equal(t, 1, len(pool.YieldedTokenInfos))
		require.Equal(t, expectedTokenSymbol, pool.YieldedTokenInfos[0].RemainingAmount.Denom)
		require.True(t, pool.YieldedTokenInfos[0].RemainingAmount.Amount.Equal(expectedDec))
		require.Equal(t, expectedStartBlockHeightToYield, pool.YieldedTokenInfos[0].StartBlockHeightToYield)
		require.True(t, pool.YieldedTokenInfos[0].AmountYieldedPerBlock.Equal(expectedDec))
		require.Equal(t, expectedTokenSymbol, pool.DepositAmount.Denom)
		require.True(t, pool.DepositAmount.Amount.Equal(expectedDec))
		require.Equal(t, expectedTokenSymbol, pool.TotalValueLocked.Denom)
		require.True(t, pool.TotalValueLocked.Amount.Equal(expectedDec))
	}

	mockCli.EXPECT().Query(types.QueryPoolsPath, cmn.HexBytes(queryBytes)).Return(nil, errors.New("default error"))
	_, err = mockCli.Farm().QueryPools()
	require.Error(t, err)

	mockCli.EXPECT().Query(types.QueryPoolsPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	require.Panics(t, func() {
		_, _ = mockCli.Farm().QueryPools()
	})
}

func TestFarmClient_QueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	expectedDec := sdk.NewDec(expectedTokenAmount)
	expectedRet := mockCli.BuildFarmPoolBytes(
		expectedPoolName,
		addr,
		expectedTokenSymbol,
		expectedStartBlockHeightToYield,
		expectedDec,
	)
	expectedCdc := mockCli.GetCodec()

	expectedOwnerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	queryParams := params.NewQueryPoolParams(expectedPoolName)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(types.QueryPoolPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	pool, err := mockCli.Farm().QueryPool(expectedPoolName)
	require.NoError(t, err)

	require.Equal(t, expectedPoolName, pool.Name)
	require.Equal(t, expectedOwnerAddr, pool.Owner)
	require.Equal(t, expectedTokenSymbol, pool.MinLockAmount.Denom)
	require.True(t, pool.MinLockAmount.Amount.Equal(expectedDec))
	require.Equal(t, 1, len(pool.YieldedTokenInfos))
	require.Equal(t, expectedTokenSymbol, pool.YieldedTokenInfos[0].RemainingAmount.Denom)
	require.True(t, pool.YieldedTokenInfos[0].RemainingAmount.Amount.Equal(expectedDec))
	require.Equal(t, expectedStartBlockHeightToYield, pool.YieldedTokenInfos[0].StartBlockHeightToYield)
	require.True(t, pool.YieldedTokenInfos[0].AmountYieldedPerBlock.Equal(expectedDec))
	require.Equal(t, expectedTokenSymbol, pool.DepositAmount.Denom)
	require.True(t, pool.DepositAmount.Amount.Equal(expectedDec))
	require.Equal(t, expectedTokenSymbol, pool.TotalValueLocked.Denom)
	require.True(t, pool.TotalValueLocked.Amount.Equal(expectedDec))

	mockCli.EXPECT().Query(types.QueryPoolPath, cmn.HexBytes(queryBytes)).Return(nil, errors.New("default error"))
	_, err = mockCli.Farm().QueryPool(expectedPoolName)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.QueryPoolPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	require.Panics(t, func() {
		_, _ = mockCli.Farm().QueryPool(expectedPoolName)
	})
}

func TestFarmClient_QueryAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	poolName1 := fmt.Sprintf("%s%d", expectedPoolName, 1)
	poolName2 := fmt.Sprintf("%s%d", expectedPoolName, 2)
	poolName3 := fmt.Sprintf("%s%d", expectedPoolName, 3)
	expectedRet := mockCli.BuildFarmPoolNameListBytes(poolName1, poolName2, poolName3)
	expectedCdc := mockCli.GetCodec()

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	queryParams := params.NewQueryAccountParams(accAddr)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(types.QueryAccountPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	poolNameList, err := mockCli.Farm().QueryAccount(addr)
	require.NoError(t, err)

	require.Equal(t, 3, len(poolNameList))
	require.Equal(t, poolName1, poolNameList[0])
	require.Equal(t, poolName2, poolNameList[1])
	require.Equal(t, poolName3, poolNameList[2])

	mockCli.EXPECT().Query(types.QueryAccountPath, cmn.HexBytes(queryBytes)).Return(nil, errors.New("default error"))
	_, err = mockCli.Farm().QueryAccount(addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.QueryAccountPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	require.Panics(t, func() {
		_, _ = mockCli.Farm().QueryAccount(addr)
	})

	_, err = mockCli.Farm().QueryAccount(addr[1:])
	require.Error(t, err)
}

func TestFarmClient_QueryAccountsLockedTo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	expectedRet := mockCli.BuildAccAddrListBytes(accAddr)
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryPoolParams(expectedPoolName)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(types.QueryAccountsLockedToPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	accAddrList, err := mockCli.Farm().QueryAccountsLockedTo(expectedPoolName)
	require.NoError(t, err)

	require.Equal(t, 1, len(accAddrList))
	require.True(t, accAddrList[0].Equals(accAddr))

	mockCli.EXPECT().Query(types.QueryAccountsLockedToPath, cmn.HexBytes(queryBytes)).Return(nil, errors.New("default error"))
	_, err = mockCli.Farm().QueryAccountsLockedTo(expectedPoolName)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.QueryAccountsLockedToPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	require.Panics(t, func() {
		_, _ = mockCli.Farm().QueryAccountsLockedTo(expectedPoolName)
	})
}

func TestFarmClient_QueryLockInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedRet := mockCli.BuildLockInfoBytes(accAddr, expectedPoolName, expectedTokenSymbol, sdk.NewDec(expectedTokenAmount),
		expectedHeight, expectedReferencePeriod)
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryPoolAccountParams(expectedPoolName, accAddr)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(types.QueryLockInfoPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	lockInfo, err := mockCli.Farm().QueryLockInfo(expectedPoolName, addr)
	require.NoError(t, err)

	require.True(t, lockInfo.Owner.Equals(accAddr))
	require.Equal(t, expectedPoolName, lockInfo.PoolName)
	require.Equal(t, expectedTokenSymbol, lockInfo.Amount.Denom)
	require.True(t, lockInfo.Amount.Amount.Equal(sdk.NewDec(expectedTokenAmount)))
	require.Equal(t, expectedHeight, lockInfo.StartBlockHeight)
	require.Equal(t, expectedReferencePeriod, lockInfo.ReferencePeriod)

	mockCli.EXPECT().Query(types.QueryLockInfoPath, cmn.HexBytes(queryBytes)).Return(nil, errors.New("default error"))
	_, err = mockCli.Farm().QueryLockInfo(expectedPoolName, addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.QueryLockInfoPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	require.Panics(t, func() {
		_, _ = mockCli.Farm().QueryLockInfo(expectedPoolName, addr)
	})

	_, err = mockCli.Farm().QueryLockInfo(expectedPoolName, addr[1:])
	require.Error(t, err)
}
