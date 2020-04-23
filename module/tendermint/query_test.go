package tendermint

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"
	"testing"
	"time"
)

func TestTendermintClient_QueryBlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	height, blockTime := int64(1024), time.Now()
	appHash, blockIDHash := cmn.HexBytes("default app hash"), cmn.HexBytes("default block ID hash")

	expectedRet := mockCli.GetRawResultBlockPointer("default chainID", height, blockTime, appHash, blockIDHash)
	expectedCdc := mockCli.GetCodec()

	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Block(gomock.AssignableToTypeOf(&height)).Return(expectedRet, nil)

	block, err := mockCli.Tendermint().QueryBlock(1024)
	require.NoError(t, err)
	require.Equal(t, "default chainID", block.ChainID)
	require.Equal(t, appHash, block.AppHash)
	require.Equal(t, height, block.Header.Height)
	require.Equal(t, blockIDHash, block.LastCommit.BlockID.Hash)
	require.True(t, blockTime.Equal(block.Time))

	mockCli.EXPECT().Block(gomock.AssignableToTypeOf(&height)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryBlock(1024)
	require.Error(t, err)
}

func TestTendermintClient_QueryBlockResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	power, height := int64(1000), int64(1024)
	pubkeyType, eventType := "default pubkey type", "default event type"
	kvPairKey := []byte("default kv pair key")
	expectedRet := mockCli.GetRawResultBlockResultsPointer(power, height, pubkeyType, eventType, kvPairKey)

	mockCli.EXPECT().BlockResults(gomock.AssignableToTypeOf(&height)).Return(expectedRet, nil)

	block, err := mockCli.Tendermint().QueryBlockResults(1024)
	require.NoError(t, err)
	require.Equal(t, height, block.Height)
	require.Equal(t, eventType, block.Results.BeginBlock.Events[0].Type)
	require.Equal(t, kvPairKey, block.Results.BeginBlock.Events[0].Attributes[0].Key)
	require.Equal(t, pubkeyType, block.Results.EndBlock.ValidatorUpdates[0].PubKey.Type)
	require.Equal(t, power, block.Results.EndBlock.ValidatorUpdates[0].Power)

	mockCli.EXPECT().BlockResults(gomock.AssignableToTypeOf(&height)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryBlockResults(1024)
	require.Error(t, err)
}
