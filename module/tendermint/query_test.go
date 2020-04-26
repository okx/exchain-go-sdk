package tendermint

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"
	tmtypes "github.com/tendermint/tendermint/types"
	"testing"
	"time"
)

const (
	addr      = "okchain1alq9na49n9yycysh889rl90g9nhe58lcv27tfj"
	valConsPK = "okchainvalconspub1zcjduepqpjq9n8g6fnjrys5t07cqcdcptu5d06tpxvhdu04mdrc4uc5swmmqfu3wku"
)

func TestTendermintClient_QueryBlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
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
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
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

func TestTendermintClient_QueryCommitResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	height, blockTime := int64(1024), time.Now()
	appHash, blockIDHash := cmn.HexBytes("default app hash"), cmn.HexBytes("default block ID hash")

	expectedRet := mockCli.GetRawCommitResultPointer(true, "default chainID", height, blockTime, appHash, blockIDHash)
	mockCli.EXPECT().Commit(gomock.AssignableToTypeOf(&height)).Return(expectedRet, nil)

	commitResult, err := mockCli.Tendermint().QueryCommitResult(1024)
	require.NoError(t, err)
	require.Equal(t, true, commitResult.CanonicalCommit)
	require.Equal(t, "default chainID", commitResult.ChainID)
	require.Equal(t, appHash, commitResult.AppHash)
	require.Equal(t, height, commitResult.Header.Height)
	require.Equal(t, blockIDHash, commitResult.Commit.BlockID.Hash)
	require.True(t, blockTime.Equal(commitResult.Time))

	mockCli.EXPECT().Commit(gomock.AssignableToTypeOf(&height)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryCommitResult(1024)
	require.Error(t, err)
}

func TestTendermintClient_QueryValidatorsResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	height, votingPower, proposerPriority := int64(1024), int64(2048), int64(-1024)
	consPubkey, err := sdk.GetConsPubKeyBech32(valConsPK)
	require.NoError(t, err)

	expectedRet := mockCli.GetRawValidatorsResultPointer(height, votingPower, proposerPriority, consPubkey)
	mockCli.EXPECT().Validators(gomock.AssignableToTypeOf(&height)).Return(expectedRet, nil)

	valsResult, err := mockCli.Tendermint().QueryValidatorsResult(1024)
	require.NoError(t, err)
	require.Equal(t, height, valsResult.BlockHeight)
	require.Equal(t, proposerPriority, valsResult.Validators[0].ProposerPriority)
	require.Equal(t, votingPower, valsResult.Validators[0].VotingPower)
	require.Equal(t, consPubkey, valsResult.Validators[0].PubKey)

	mockCli.EXPECT().Validators(gomock.AssignableToTypeOf(&height)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryValidatorsResult(1024)
	require.Error(t, err)
}

func TestTendermintClient_QueryTxResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	txHash, tx := []byte("default tx hash"), []byte("default tx")
	height, code := int64(1024), uint32(0)
	log, eventType := "default log", "default event type"

	expectedRet := mockCli.GetRawTxResultPointer(txHash, height, code, log, eventType, tx)
	mockCli.EXPECT().Tx(txHash, true).Return(expectedRet, nil)

	txResult, err := mockCli.Tendermint().QueryTxResult(txHash, true)
	require.NoError(t, err)
	require.Equal(t, height, txResult.Height)
	require.Equal(t, cmn.HexBytes(txHash), txResult.Hash)
	require.Equal(t, tmtypes.Tx(tx), txResult.Tx)
	require.Equal(t, log, txResult.TxResult.Log)
	require.Equal(t, code, txResult.TxResult.Code)
	require.Equal(t, eventType, txResult.TxResult.Events[0].Type)

	mockCli.EXPECT().Tx(txHash, true).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryTxResult(txHash, true)
	require.Error(t, err)
}

func TestTendermintClient_QueryTxsResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	txHash, tx := []byte("default tx hash"), []byte("default tx")
	height, code := int64(1024), uint32(0)
	log, eventType := "default log", "default event type"

	expectedRet := mockCli.GetRawResultTxSearchPointer(1, txHash, height, code, log, eventType, tx)
	mockCli.EXPECT().TxSearch(gomock.AssignableToTypeOf(""), false, gomock.AssignableToTypeOf(0),
		gomock.AssignableToTypeOf(0)).Return(expectedRet, nil)

	queryStr := fmt.Sprintf("message.sender=%s", addr)
	txSearchResult, err := mockCli.Tendermint().QueryTxsResult(queryStr, 1, 30)
	require.NoError(t, err)
	require.Equal(t, 1, txSearchResult.TotalCount)
	require.Equal(t, height, txSearchResult.Txs[0].Height)
	require.Equal(t, cmn.HexBytes(txHash), txSearchResult.Txs[0].Hash)
	require.Equal(t, tmtypes.Tx(tx), txSearchResult.Txs[0].Tx)
	require.Equal(t, log, txSearchResult.Txs[0].TxResult.Log)
	require.Equal(t, code, txSearchResult.Txs[0].TxResult.Code)
	require.Equal(t, eventType, txSearchResult.Txs[0].TxResult.Events[0].Type)

	mockCli.EXPECT().TxSearch(gomock.AssignableToTypeOf(""), false, gomock.AssignableToTypeOf(0),
		gomock.AssignableToTypeOf(0)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryTxsResult(queryStr, 1, 30)
	require.Error(t, err)

	badQueryStr := fmt.Sprintf("message.sender%s", addr)
	_, err = mockCli.Tendermint().QueryTxsResult(badQueryStr, 1, 30)
	require.Error(t, err)

	badQueryStr = fmt.Sprintf("message.sender==%s", addr)
	_, err = mockCli.Tendermint().QueryTxsResult(badQueryStr, 1, 30)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryTxsResult(queryStr, -1, 30)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryTxsResult(queryStr, 1, -30)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryTxsResult("", 1, 30)
	require.Error(t, err)
}
