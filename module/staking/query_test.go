package staking

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/staking/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"
	"testing"
	"time"
)

const (
	addr      = "okchain1alq9na49n9yycysh889rl90g9nhe58lcv27tfj"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okchainpub1addwnpepqgzuks5c07kfce85e0t0x8qkuvvxu874965ruafn6svhjrhswt0lgdj85lv"
	mnemonic  = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	memo      = "my memo"
	valAddr   = "okchainvaloper1alq9na49n9yycysh889rl90g9nhe58lcs50wu5"
	valConsPK = "okchainvalconspub1zcjduepqpjq9n8g6fnjrys5t07cqcdcptu5d06tpxvhdu04mdrc4uc5swmmqfu3wku"
)

var (
	// an extremely strict way to check
	rawValBytes = []byte{123, 10, 20, 239, 192, 89, 246, 165, 153, 72, 76, 18, 23, 57, 202, 63, 149, 232, 44, 239, 154, 31, 248, 18, 37, 22, 36, 222, 100, 32, 12, 128, 89, 157, 26, 76, 228, 50, 66, 139, 127, 176, 12, 55, 1, 95, 40, 215, 233, 97, 51, 46, 222, 62, 187, 104, 241, 94, 98, 144, 118, 246, 32, 2, 42, 1, 48, 50, 9, 49, 48, 48, 48, 48, 48, 48, 48, 48, 58, 7, 10, 5, 110, 111, 100, 101, 50, 82, 27, 10, 25, 10, 9, 49, 48, 48, 48, 48, 48, 48, 48, 48, 18, 9, 49, 48, 48, 48, 48, 48, 48, 48, 48, 26, 1, 48, 90, 6, 49, 48, 48, 48, 48, 48}
)

func TestStakingClient_QueryValidators(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewStakingClient(mockCli.MockBaseClient))

	valOperAddr, err := sdk.ValAddressFromBech32(valAddr)
	require.NoError(t, err)
	delegatorShares, err := sdk.NewDecFromStr("1")
	require.NoError(t, err)
	minSelfDelegation, err := sdk.NewDecFromStr("0.001")
	require.NoError(t, err)

	// build expected return of the slice of cmn.KVPair
	expectedRet := []cmn.KVPair{
		{
			Key:   append(types.ValidatorsKey, valOperAddr.Bytes()...),
			Value: rawValBytes,
		},
	}
	expectedCdc := mockCli.GetCodec()

	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().QuerySubspace(types.ValidatorsKey, types.ModuleName).Return(expectedRet, nil)

	vals, err := mockCli.Staking().QueryValidators()
	require.NoError(t, err)

	// an extremely strict way to check by raw bytes
	require.Equal(t, valOperAddr, vals[0].OperatorAddress)
	require.Equal(t, valConsPK, vals[0].ConsPubKey)
	require.Equal(t, false, vals[0].Jailed)
	require.Equal(t, byte(2), vals[0].Status)
	require.Equal(t, delegatorShares, vals[0].DelegatorShares)
	require.Equal(t, int64(0), vals[0].UnbondingHeight)
	require.Equal(t, minSelfDelegation, vals[0].MinSelfDelegation)
	require.True(t, time.Unix(0, 0).UTC().Equal(vals[0].UnbondingCompletionTime))

}

func TestStakingClient_QueryValidator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewStakingClient(mockCli.MockBaseClient))

	valOperAddr, err := sdk.ValAddressFromBech32(valAddr)
	require.NoError(t, err)
	delegatorShares, err := sdk.NewDecFromStr("1")
	require.NoError(t, err)
	minSelfDelegation, err := sdk.NewDecFromStr("0.001")
	require.NoError(t, err)
	unbondingCompletionTime := time.Now()

	expectedRet := mockCli.BuildValidatorBytes(valOperAddr, valConsPK, "default moniker", "default identity",
		"default website", "default details", 2, delegatorShares, minSelfDelegation, 0,
		unbondingCompletionTime, false)
	expectedCdc := mockCli.GetCodec()

	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().QueryStore(cmn.HexBytes(types.GetValidatorKey(valOperAddr)), ModuleName, "key").Return(expectedRet, nil)

	val, err := mockCli.Staking().QueryValidator(valAddr)
	require.NoError(t, err)

	require.Equal(t, valOperAddr, val.OperatorAddress)
	require.Equal(t, valConsPK, val.ConsPubKey)
	require.Equal(t, false, val.Jailed)
	require.Equal(t, byte(2), val.Status)
	require.Equal(t, delegatorShares, val.DelegatorShares)
	require.Equal(t, "default moniker", val.Description.Moniker)
	require.Equal(t, "default identity", val.Description.Identity)
	require.Equal(t, "default website", val.Description.Website)
	require.Equal(t, "default details", val.Description.Details)
	require.Equal(t, int64(0), val.UnbondingHeight)
	require.Equal(t, minSelfDelegation, val.MinSelfDelegation)
	require.True(t, unbondingCompletionTime.Equal(val.UnbondingCompletionTime))

	_, err = mockCli.Staking().QueryValidator(valAddr[1:])
	require.Error(t, err)

	mockCli.EXPECT().QueryStore(cmn.HexBytes(types.GetValidatorKey(valOperAddr)), ModuleName, "key").Return([]byte{1}, errors.New("default error"))
	_, err = mockCli.Staking().QueryValidator(valAddr)
	require.Error(t, err)

	mockCli.EXPECT().QueryStore(cmn.HexBytes(types.GetValidatorKey(valOperAddr)), ModuleName, "key").Return(nil, nil)
	_, err = mockCli.Staking().QueryValidator(valAddr)
	require.Error(t, err)

}
