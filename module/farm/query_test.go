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

var (
	// an extremely strict way to check
	rawPoolBytes = []byte{141, 1, 10, 20, 178, 97, 64, 82, 50, 18, 5, 108, 132, 29, 215, 59, 247, 139, 201, 47, 156, 195, 78, 221, 18, 17, 100, 101, 102, 97, 117, 108, 116, 45, 112, 111, 111, 108, 45, 110, 97, 109, 101, 26, 23, 10, 7, 97, 98, 99, 45, 100, 53, 51, 18, 12, 49, 48, 50, 52, 48, 48, 48, 48, 48, 48, 48, 48, 34, 17, 10, 3, 111, 107, 116, 18, 10, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 42, 12, 10, 7, 97, 98, 99, 45, 100, 53, 51, 18, 1, 48, 50, 40, 10, 23, 10, 7, 97, 98, 99, 45, 100, 53, 51, 18, 12, 49, 48, 50, 52, 48, 48, 48, 48, 48, 48, 48, 48, 16, 232, 7, 26, 10, 53, 48, 48, 48, 48, 48, 48, 48, 48, 48}
)

func TestFarmClient_QueryPools(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	expectedDec := sdk.NewDec(expectedTokenAmount)
	expectedOwnerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	// build expected return of the slice of cmn.KVPair
	expectedRet := []cmn.KVPair{
		{
			Key:   append(types.FarmPoolPrefix, []byte(expectedPoolName)...),
			Value: rawPoolBytes,
		},
	}
	expectedCdc := mockCli.GetCodec()

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().QuerySubspace(types.FarmPoolPrefix, types.ModuleName).Return(expectedRet, nil)

	pools, err := mockCli.Farm().QueryPools()
	require.NoError(t, err)

	// an extremely strict way to check by raw bytes
	require.Equal(t, 1, len(pools))
	require.Equal(t, "default-pool-name", pools[0].Name)
	require.Equal(t, expectedOwnerAddr, pools[0].Owner)
	require.Equal(t, expectedTokenSymbol, pools[0].MinLockAmount.Denom)
	require.Equal(t, "okt", pools[0].DepositAmount.Denom)
	require.True(t, pools[0].DepositAmount.Amount.Equal(sdk.NewDec(10)))
	require.Equal(t, expectedTokenSymbol, pools[0].TotalValueLocked.Denom)
	require.Equal(t, 1, len(pools[0].YieldedTokenInfos))
	require.Equal(t, expectedTokenSymbol, pools[0].YieldedTokenInfos[0].RemainingAmount.Denom)
	require.True(t, pools[0].YieldedTokenInfos[0].RemainingAmount.Amount.Equal(expectedDec))
	require.Equal(t, expectedStartBlockHeightToYield, pools[0].YieldedTokenInfos[0].StartBlockHeightToYield)
	require.True(t, pools[0].YieldedTokenInfos[0].AmountYieldedPerBlock.Equal(sdk.NewDec(expectedAmountYieldPerBlock)))

	mockCli.EXPECT().QuerySubspace(types.FarmPoolPrefix, types.ModuleName).Return(nil, errors.New("default error"))
	_, err = mockCli.Farm().QueryPools()
	require.Error(t, err)

	badRet := []cmn.KVPair{
		{
			Key: append(types.FarmPoolPrefix, []byte(expectedPoolName)...),
			// bad encoded bytes
			Value: rawPoolBytes[1:],
		},
	}
	mockCli.EXPECT().QuerySubspace(types.FarmPoolPrefix, types.ModuleName).Return(badRet, nil)
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
