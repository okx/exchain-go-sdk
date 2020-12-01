package staking

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/staking/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"testing"
	"time"
)

const (
	addr      = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okexchainpub17weu6qepq0ph2t3u697qar7rmdtdtqp4744jcprjd2h356zr0yh5vmw38a3my4vqjx5"
	mnemonic  = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	memo      = "my memo"
	valAddr   = "okexchainvaloper1ntvyep3suq5z7789g7d5dejwzameu08mmv8pca"
	valConsPK = "okexchainvalconspub1zcjduepq24jtmdyzapg50mevhfnhjl09q876xe5dj4ajsda9q6at2dtrpvmse0tav6"
	proxyAddr = "okexchain193xnjknz3e52mqv2nyufnzjugu3mh65rpxdasn"
)

var (
	// an extremely strict way to check
	rawValBytes = []byte{123, 10, 20, 239, 192, 89, 246, 165, 153, 72, 76, 18, 23, 57, 202, 63, 149, 232, 44, 239, 154, 31, 248, 18, 37, 22, 36, 222, 100, 32, 12, 128, 89, 157, 26, 76, 228, 50, 66, 139, 127, 176, 12, 55, 1, 95, 40, 215, 233, 97, 51, 46, 222, 62, 187, 104, 241, 94, 98, 144, 118, 246, 32, 2, 42, 1, 48, 50, 9, 49, 48, 48, 48, 48, 48, 48, 48, 48, 58, 7, 10, 5, 110, 111, 100, 101, 50, 82, 27, 10, 25, 10, 9, 49, 48, 48, 48, 48, 48, 48, 48, 48, 18, 9, 49, 48, 48, 48, 48, 48, 48, 48, 48, 26, 1, 48, 90, 6, 49, 48, 48, 48, 48, 48}
)

func TestStakingClient_QueryValidators(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	//TODO
	//mockCli.RegisterModule(NewStakingClient(mockCli.MockBaseClient))

	//valOperAddr, err := sdk.ValAddressFromBech32(valAddr)
	//require.NoError(t, err)
	//delegatorShares, err := sdk.NewDecFromStr("1")
	//require.NoError(t, err)
	//minSelfDelegation, err := sdk.NewDecFromStr("0.001")
	//require.NoError(t, err)

	// build expected return of the slice of cmn.KVPair
	//expectedRet := []kv.Pair{
	//	{
	//		Key:   append(types.ValidatorsKey, valOperAddr.Bytes()...),
	//		Value: rawValBytes,
	//	},
	//}
	expectedCdc := mockCli.GetCodec()

	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	//mockCli.EXPECT().QuerySubspace(types.ValidatorsKey, types.ModuleName).Return(expectedRet, nil)

	//vals, err := mockCli.Staking().QueryValidators()
	require.NoError(t, err)

	// an extremely strict way to check by raw bytes
	//require.Equal(t, valOperAddr, vals[0].OperatorAddress)
	//require.Equal(t, valConsPK, vals[0].ConsPubKey)
	//require.Equal(t, false, vals[0].Jailed)
	//require.Equal(t, byte(2), vals[0].Status)
	//require.Equal(t, delegatorShares, vals[0].DelegatorShares)
	//require.Equal(t, int64(0), vals[0].UnbondingHeight)
	//require.Equal(t, minSelfDelegation, vals[0].MinSelfDelegation)
	//require.True(t, time.Unix(0, 0).UTC().Equal(vals[0].UnbondingCompletionTime))

	//mockCli.EXPECT().QuerySubspace(types.ValidatorsKey, types.ModuleName).Return(expectedRet, errors.New("default error"))
	//_, err = mockCli.Staking().QueryValidators()
	require.Error(t, err)
}

func TestStakingClient_QueryValidator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	//TODO
	//mockCli.RegisterModule(NewStakingClient(mockCli.MockBaseClient))

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
	mockCli.EXPECT().QueryStore(tmbytes.HexBytes(types.GetValidatorKey(valOperAddr)), ModuleName, "key").
		Return(expectedRet, nil)

	//val, err := mockCli.Staking().QueryValidator(valAddr)
	require.NoError(t, err)

	//require.Equal(t, valOperAddr, val.OperatorAddress)
	//require.Equal(t, valConsPK, val.ConsPubKey)
	//require.Equal(t, false, val.Jailed)
	//require.Equal(t, byte(2), val.Status)
	//require.Equal(t, delegatorShares, val.DelegatorShares)
	//require.Equal(t, "default moniker", val.Description.Moniker)
	//require.Equal(t, "default identity", val.Description.Identity)
	//require.Equal(t, "default website", val.Description.Website)
	//require.Equal(t, "default details", val.Description.Details)
	//require.Equal(t, int64(0), val.UnbondingHeight)
	//require.Equal(t, minSelfDelegation, val.MinSelfDelegation)
	//require.True(t, unbondingCompletionTime.Equal(val.UnbondingCompletionTime))
	//
	//_, err = mockCli.Staking().QueryValidator(valAddr[1:])
	require.Error(t, err)

	mockCli.EXPECT().QueryStore(tmbytes.HexBytes(types.GetValidatorKey(valOperAddr)), ModuleName, "key").
		Return([]byte{1}, errors.New("default error"))
	//_, err = mockCli.Staking().QueryValidator(valAddr)
	require.Error(t, err)

	mockCli.EXPECT().QueryStore(tmbytes.HexBytes(types.GetValidatorKey(valOperAddr)), ModuleName, "key").Return(nil, nil)
	//_, err = mockCli.Staking().QueryValidator(valAddr)
	require.Error(t, err)
}

func TestStakingClient_QueryDelegator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	// TODO
	//mockCli.RegisterModule(NewStakingClient(mockCli.MockBaseClient))

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
	queryBytes := expectedCdc.MustMarshalJSON(params.NewQueryDelegatorParams(delAddr))

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(6)
	mockCli.EXPECT().QueryStore(tmbytes.HexBytes(types.GetDelegatorKey(delAddr)), ModuleName, "key").
		Return(expectedRet1, nil).Times(2)
	mockCli.EXPECT().Query(types.UnbondDelegationPath, tmbytes.HexBytes(queryBytes)).Return(expectedRet2, nil)

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

	mockCli.EXPECT().Query(types.UnbondDelegationPath, tmbytes.HexBytes(queryBytes)).Return(expectedRet2[1:], nil)
	_, err = mockCli.Staking().QueryDelegator(addr)
	require.Error(t, err)

	mockCli.EXPECT().QueryStore(tmbytes.HexBytes(types.GetDelegatorKey(delAddr)), ModuleName, "key").
		Return(expectedRet1, errors.New("default error"))
	_, err = mockCli.Staking().QueryDelegator(addr)
	require.Error(t, err)
}
