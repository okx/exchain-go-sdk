package staking

import (
	"errors"
	"fmt"
	"testing"
	"time"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/exchain-go-sdk/mocks"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	stakingtypes "github.com/okex/exchain/x/staking/types"
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
	valAddr   = "exvaloper1qwuag8gx408m9ej038vzx50ntt0x4yrq8qwdtq"
	valConsPK = "exvalconspub1zcjduepqs2c6xnrfjwxzfclrpq4rh5mxrwlxmncvq6l48ah3ccdew2j6nnfsh3tc5f"
	proxyAddr = "ex1alrwch5sd3wm3np4njz7l754xtnng6cf4z9s5v"

	defaultMoniker  = "default moniker"
	defaultIdentity = "default identity"
	defaultWebsite  = "default website"
	defaultDetails  = "default details"
)

func TestStakingClient_QueryValidators(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewStakingClient(mockCli.MockBaseClient))

	valOperAddr, err := sdk.ValAddressFromBech32(valAddr)
	require.NoError(t, err)
	delegatorShares, err := sdk.NewDecFromStr("1")
	require.NoError(t, err)
	minSelfDelegation, err := sdk.NewDecFromStr("10000")
	require.NoError(t, err)
	unbondingCompletionTime := time.Now()

	expectedRet := mockCli.BuildValidatorsBytes(valOperAddr, valConsPK, defaultMoniker, defaultIdentity, defaultWebsite,
		defaultDetails, 2, delegatorShares, minSelfDelegation, 0, unbondingCompletionTime, false,
		true)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	expectedPath := fmt.Sprintf("custom/%s/%s", stakingtypes.QuerierRoute, stakingtypes.QueryValidators)
	expectedParams, err := expectedCdc.MarshalJSON(stakingtypes.NewQueryValidatorsParams(1, 0, "all"))
	require.NoError(t, err)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	vals, err := mockCli.Staking().QueryValidators()
	require.NoError(t, err)

	require.Equal(t, 1, len(vals))
	require.Equal(t, valOperAddr, vals[0].OperatorAddress)
	expectedValConsPK, err := stakingtypes.GetConsPubKeyBech32(valConsPK)
	require.NoError(t, err)
	require.True(t, expectedValConsPK.Equals(vals[0].ConsPubKey))
	require.Equal(t, false, vals[0].Jailed)
	require.Equal(t, sdk.BondStatus(2), vals[0].Status)
	require.Equal(t, delegatorShares, vals[0].DelegatorShares)
	require.Equal(t, int64(0), vals[0].UnbondingHeight)
	require.Equal(t, minSelfDelegation, vals[0].MinSelfDelegation)
	require.True(t, unbondingCompletionTime.Equal(vals[0].UnbondingCompletionTime))
	require.Equal(t, defaultMoniker, vals[0].Description.Moniker)
	require.Equal(t, defaultIdentity, vals[0].Description.Identity)
	require.Equal(t, defaultWebsite, vals[0].Description.Website)
	require.Equal(t, defaultDetails, vals[0].Description.Details)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Staking().QueryValidators()
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Staking().QueryValidators()
	require.Error(t, err)
}

func TestStakingClient_QueryValidator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewStakingClient(mockCli.MockBaseClient))

	valOperAddr, err := sdk.ValAddressFromBech32(valAddr)
	require.NoError(t, err)
	delegatorShares, err := sdk.NewDecFromStr("1")
	require.NoError(t, err)
	minSelfDelegation, err := sdk.NewDecFromStr("10000")
	require.NoError(t, err)
	unbondingCompletionTime := time.Now()

	expectedRet := mockCli.BuildValidatorsBytes(valOperAddr, valConsPK, defaultMoniker, defaultIdentity, defaultWebsite,
		defaultDetails, 2, delegatorShares, minSelfDelegation, 0,
		unbondingCompletionTime, false, false)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	expectedPath := fmt.Sprintf("custom/%s/%s", stakingtypes.QuerierRoute, stakingtypes.QueryValidator)
	expectedParams, err := expectedCdc.MarshalJSON(stakingtypes.NewQueryValidatorParams(valOperAddr))
	require.NoError(t, err)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	val, err := mockCli.Staking().QueryValidator(valAddr)
	require.NoError(t, err)

	require.Equal(t, valOperAddr, val.OperatorAddress)
	expectedValConsPK, err := stakingtypes.GetConsPubKeyBech32(valConsPK)
	require.NoError(t, err)
	require.True(t, expectedValConsPK.Equals(val.ConsPubKey))
	require.Equal(t, false, val.Jailed)
	require.Equal(t, sdk.BondStatus(2), val.Status)
	require.Equal(t, delegatorShares, val.DelegatorShares)
	require.Equal(t, int64(0), val.UnbondingHeight)
	require.Equal(t, minSelfDelegation, val.MinSelfDelegation)
	require.True(t, unbondingCompletionTime.Equal(val.UnbondingCompletionTime))
	require.Equal(t, defaultMoniker, val.Description.Moniker)
	require.Equal(t, defaultIdentity, val.Description.Identity)
	require.Equal(t, defaultWebsite, val.Description.Website)
	require.Equal(t, defaultDetails, val.Description.Details)

	_, err = mockCli.Staking().QueryValidator(valAddr[1:])
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Staking().QueryValidator(valAddr)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Staking().QueryValidator(valAddr)
	require.Error(t, err)
}

func TestStakingClient_QueryDelegator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewStakingClient(mockCli.MockBaseClient))

	delAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	proxyAddr, err := sdk.AccAddressFromBech32(proxyAddr)
	require.NoError(t, err)
	valAddr, err := sdk.ValAddressFromBech32(valAddr)
	require.NoError(t, err)
	shares, err := sdk.NewDecFromStr("10240000.1024")
	require.NoError(t, err)
	tokens, err := sdk.NewDecFromStr("10.24")
	require.NoError(t, err)
	totalDelegatedTokens, err := sdk.NewDecFromStr("20.48")
	require.NoError(t, err)
	quantity, err := sdk.NewDecFromStr("40.96")
	require.NoError(t, err)
	completionTime := time.Now()

	expectedRet1 := mockCli.BuildDelegatorBytes(delAddr, proxyAddr, []sdk.ValAddress{valAddr}, shares, tokens,
		totalDelegatedTokens, false)
	expectedRet2 := mockCli.BuildUndelegationBytes(delAddr, quantity, completionTime)
	expectedCdc := mockCli.GetCodec()
	expectedParams, err := expectedCdc.MarshalJSON(stakingtypes.NewQueryDelegatorParams(delAddr))
	require.NoError(t, err)
	expectedPath := fmt.Sprintf("custom/%s/%s", stakingtypes.RouterKey, stakingtypes.QueryUnbondingDelegation)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(6)
	mockCli.EXPECT().QueryStore(tmbytes.HexBytes(stakingtypes.GetDelegatorKey(delAddr)), stakingtypes.StoreKey, "key").
		Return(expectedRet1, int64(1024), nil).Times(2)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet2, int64(1024), nil)

	delResp, err := mockCli.Staking().QueryDelegator(addr)
	require.NoError(t, err)
	require.Equal(t, delAddr, delResp.DelegatorAddress)
	require.Equal(t, totalDelegatedTokens, delResp.TotalDelegatedTokens)
	require.Equal(t, quantity, delResp.UnbondedTokens)
	require.Equal(t, valAddr, delResp.ValidatorAddresses[0])
	require.Equal(t, shares, delResp.Shares)
	require.Equal(t, tokens, delResp.Tokens)
	require.Equal(t, false, delResp.IsProxy)
	require.Equal(t, proxyAddr, delResp.ProxyAddress)
	require.True(t, completionTime.Equal(delResp.CompletionTime))

	_, err = mockCli.Staking().QueryDelegator(addr[1:])
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet2[1:], int64(1024), nil)
	_, err = mockCli.Staking().QueryDelegator(addr)
	require.Error(t, err)

	mockCli.EXPECT().QueryStore(tmbytes.HexBytes(stakingtypes.GetDelegatorKey(delAddr)), stakingtypes.StoreKey, "key").
		Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Staking().QueryDelegator(addr)
	require.Error(t, err)
}
