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
	require.True(t, blockTime.Equal(block.Time))
	require.Equal(t, height, block.Header.Height)
	require.Equal(t, blockIDHash, block.LastCommit.BlockID.Hash)

	mockCli.EXPECT().Block(gomock.AssignableToTypeOf(&height)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryBlock(1024)
	require.Error(t, err)
}
