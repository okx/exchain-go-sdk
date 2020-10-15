package farm

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/farm/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
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

	expectedTokenSymbol                   = "air-bb4"
	expectedTokenAmount             int64 = 1024
	expectedAmountYieldPerBlock     int64 = 50
	expectedStartBlockHeightToYield int64 = 100
)

var (
	// an extremely strict way to check
	rawPoolBytes = []byte{124, 10, 20, 178, 97, 64, 82, 50, 18, 5, 108, 132, 29, 215, 59, 247, 139, 201, 47, 156, 195, 78, 221, 18, 17, 100, 101, 102, 97, 117, 108, 116, 45, 112, 111, 111, 108, 45, 110, 97, 109, 101, 26, 7, 97, 105, 114, 45, 98, 98, 52, 34, 39, 10, 23, 10, 7, 97, 105, 114, 45, 98, 98, 52, 18, 12, 49, 48, 50, 52, 48, 48, 48, 48, 48, 48, 48, 48, 16, 100, 26, 10, 53, 48, 48, 48, 48, 48, 48, 48, 48, 48, 42, 17, 10, 3, 111, 107, 116, 18, 10, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 50, 12, 10, 7, 97, 105, 114, 45, 98, 98, 52, 18, 1, 48}
)

func TestFarmClient_QueryPools(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient))

	expectedOwnerAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	// build expected return of the slice of cmn.KVPair
	expectedRet := []cmn.KVPair{
		{
			Key:   append(types.FarmPoolPrefix, []byte("default-pool-name")...),
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
	require.Equal(t, expectedTokenSymbol, pools[0].SymbolLocked)
	require.Equal(t, 1, len(pools[0].YieldedTokenInfos))
	require.Equal(t, expectedStartBlockHeightToYield, pools[0].YieldedTokenInfos[0].StartBlockHeightToYield)
	require.Equal(t, expectedTokenSymbol, pools[0].YieldedTokenInfos[0].RemainingAmount.Denom)
	require.True(t, pools[0].YieldedTokenInfos[0].RemainingAmount.Amount.Equal(sdk.NewDec(expectedTokenAmount)))
	require.True(t, pools[0].YieldedTokenInfos[0].AmountYieldedPerBlock.Equal(sdk.NewDec(expectedAmountYieldPerBlock)))
	require.Equal(t, "okt", pools[0].DepositAmount.Denom)
	require.True(t, pools[0].DepositAmount.Amount.Equal(sdk.NewDec(10)))
	require.Equal(t, expectedTokenSymbol, pools[0].TotalValueLocked.Denom)
	require.Equal(t, 0, len(pools[0].AmountYielded))

	mockCli.EXPECT().QuerySubspace(types.FarmPoolPrefix, types.ModuleName).Return(nil, errors.New("default error"))
	_, err = mockCli.Farm().QueryPools()
	require.Error(t, err)

	badRet := []cmn.KVPair{
		{
			Key: append(types.FarmPoolPrefix, []byte("default-pool-name")...),
			// bad encoded bytes
			Value: rawPoolBytes[1:],
		},
	}
	mockCli.EXPECT().QuerySubspace(types.FarmPoolPrefix, types.ModuleName).Return(badRet, nil)
	require.Panics(t, func() {
		_, _ = mockCli.Farm().QueryPools()
	})
}
