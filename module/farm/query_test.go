package farm

import (
	"errors"
	"fmt"
	"testing"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/exchain-go-sdk/mocks"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	farmtypes "github.com/okex/exchain/x/farm/types"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/okex/exchain/libs/tendermint/libs/bytes"
)

const (
	addr      = "ex1qj5c07sm6jetjz8f509qtrxgh4psxkv3ddyq7u"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "expub17weu6qepqtfc6zq8dukwc3lhlhx7th2csfjw0g3cqnqvanh7z9c2nhkr8mn5z9uq4q6"
	mnemonic  = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	memo      = "my memo"

	expectedTokenSymbol                    = "abc-d53"
	expectedTokenAmount             int64  = 1024
	expectedStartBlockHeightToYield int64  = 1000
	expectedPoolName                       = "default-pool-name"
	expectedHeight                  int64  = 1024
	expectedReferencePeriod         uint64 = 1
)

func TestFarmClient_QueryPools(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
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

	expectedOwnerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedCdc := mockCli.GetCodec()
	// fixed to all pools query
	expectedParams := expectedCdc.MustMarshalJSON(farmtypes.NewQueryPoolsParams(1, 0))
	expectedPath := fmt.Sprintf("custom/%s/%s", farmtypes.QuerierRoute, farmtypes.QueryPools)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

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
		require.Equal(t, 1, len(pool.TotalAccumulatedRewards))
		require.Equal(t, expectedTokenSymbol, pool.TotalAccumulatedRewards[0].Denom)
		require.True(t, pool.TotalAccumulatedRewards[0].Amount.Equal(expectedDec))
	}

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Farm().QueryPools()
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Farm().QueryPools()
	require.Error(t, err)
}

func TestFarmClient_QueryPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
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

	expectedOwnerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedCdc := mockCli.GetCodec()
	// fixed to all pools query
	expectedParams := expectedCdc.MustMarshalJSON(farmtypes.NewQueryPoolParams(expectedPoolName))
	expectedPath := fmt.Sprintf("custom/%s/%s", farmtypes.QuerierRoute, farmtypes.QueryPool)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

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
	require.Equal(t, 1, len(pool.TotalAccumulatedRewards))
	require.Equal(t, expectedTokenSymbol, pool.TotalAccumulatedRewards[0].Denom)
	require.True(t, pool.TotalAccumulatedRewards[0].Amount.Equal(expectedDec))

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Farm().QueryPool(expectedPoolName)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Farm().QueryPool(expectedPoolName)
	require.Error(t, err)
}

func TestFarmClient_QueryAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	poolName1 := fmt.Sprintf("%s%d", expectedPoolName, 1)
	poolName2 := fmt.Sprintf("%s%d", expectedPoolName, 2)
	poolName3 := fmt.Sprintf("%s%d", expectedPoolName, 3)
	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedRet := mockCli.BuildFarmPoolNameListBytes(poolName1, poolName2, poolName3)
	expectedCdc := mockCli.GetCodec()
	expectedParams := expectedCdc.MustMarshalJSON(farmtypes.NewQueryAccountParams(accAddr))
	expectedPath := fmt.Sprintf("custom/%s/%s", farmtypes.QuerierRoute, farmtypes.QueryAccount)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	poolNameList, err := mockCli.Farm().QueryAccount(addr)
	require.NoError(t, err)

	require.Equal(t, 3, len(poolNameList))
	require.Equal(t, poolName1, poolNameList[0])
	require.Equal(t, poolName2, poolNameList[1])
	require.Equal(t, poolName3, poolNameList[2])

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Farm().QueryAccount(addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Farm().QueryAccount(addr)
	require.Error(t, err)

	_, err = mockCli.Farm().QueryAccount(addr[1:])
	require.Error(t, err)
}

func TestFarmClient_QueryAccountsLockedTo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	expectedRet := mockCli.BuildAccAddrListBytes(accAddr)
	expectedCdc := mockCli.GetCodec()
	expectedParams := expectedCdc.MustMarshalJSON(farmtypes.NewQueryPoolParams(expectedPoolName))
	expectedPath := fmt.Sprintf("custom/%s/%s", farmtypes.QuerierRoute, farmtypes.QueryAccountsLockedTo)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	accAddrList, err := mockCli.Farm().QueryAccountsLockedTo(expectedPoolName)
	require.NoError(t, err)

	require.Equal(t, 1, len(accAddrList))
	require.True(t, accAddrList[0].Equals(accAddr))

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Farm().QueryAccountsLockedTo(expectedPoolName)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(0), nil)
	_, err = mockCli.Farm().QueryAccountsLockedTo(expectedPoolName)
	require.Error(t, err)
}

func TestFarmClient_QueryLockInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	accAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	expectedRet := mockCli.BuildLockInfoBytes(accAddr, expectedPoolName, expectedTokenSymbol, sdk.NewDec(expectedTokenAmount),
		expectedHeight, expectedReferencePeriod)
	expectedCdc := mockCli.GetCodec()

	expectedParams := expectedCdc.MustMarshalJSON(farmtypes.NewQueryPoolAccountParams(expectedPoolName, accAddr))
	expectedPath := fmt.Sprintf("custom/%s/%s", farmtypes.QuerierRoute, farmtypes.QueryLockInfo)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	lockInfo, err := mockCli.Farm().QueryLockInfo(expectedPoolName, addr)
	require.NoError(t, err)

	require.True(t, lockInfo.Owner.Equals(accAddr))
	require.Equal(t, expectedPoolName, lockInfo.PoolName)
	require.Equal(t, expectedTokenSymbol, lockInfo.Amount.Denom)
	require.True(t, lockInfo.Amount.Amount.Equal(sdk.NewDec(expectedTokenAmount)))
	require.Equal(t, expectedHeight, lockInfo.StartBlockHeight)
	require.Equal(t, expectedReferencePeriod, lockInfo.ReferencePeriod)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Farm().QueryLockInfo(expectedPoolName, addr)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Farm().QueryLockInfo(expectedPoolName, addr)
	require.Error(t, err)

	_, err = mockCli.Farm().QueryLockInfo(expectedPoolName, addr[1:])
	require.Error(t, err)
}
