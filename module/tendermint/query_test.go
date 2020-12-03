package tendermint

import (
	"encoding/hex"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	okexchain "github.com/okex/okexchain/app/types"
	stakingtypes "github.com/okex/okexchain/x/staking/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

const (
	addr      = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	valConsPK = "okexchainvalconspub1zcjduepq24jtmdyzapg50mevhfnhjl09q876xe5dj4ajsda9q6at2dtrpvmse0tav6"
)

func TestTendermintClient_QueryBlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	height, blockTime := int64(1024), time.Now()
	appHash, blockIDHash := tmbytes.HexBytes("default app hash"), tmbytes.HexBytes("default block ID hash")

	expectedRet := mockCli.GetRawResultBlockPointer("default chainID", height, blockTime, appHash, blockIDHash)
	mockCli.EXPECT().Block(gomock.AssignableToTypeOf(&height)).Return(expectedRet, nil)

	block, err := mockCli.Tendermint().QueryBlock(height)
	require.NoError(t, err)
	require.Equal(t, "default chainID", block.ChainID)
	require.Equal(t, appHash, block.AppHash)
	require.Equal(t, height, block.Header.Height)
	require.Equal(t, blockIDHash, block.LastCommit.BlockID.Hash)
	require.True(t, blockTime.Equal(block.Time))

	mockCli.EXPECT().Block(gomock.AssignableToTypeOf(&height)).Return(nil, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryBlock(height)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryBlock(-1)
	require.Error(t, err)
}

func TestTendermintClient_QueryBlockResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	power, height := int64(1000), int64(1024)
	pubkeyType, eventType := "default pubkey type", "default event type"
	kvPairKey := []byte("default kv pair key")
	expectedRet := mockCli.GetRawResultBlockResultsPointer(power, height, pubkeyType, eventType, kvPairKey)

	mockCli.EXPECT().BlockResults(gomock.AssignableToTypeOf(&height)).Return(expectedRet, nil)

	blockResults, err := mockCli.Tendermint().QueryBlockResults(height)
	require.NoError(t, err)
	require.Equal(t, height, blockResults.Height)
	require.Equal(t, 1, len(blockResults.TxsResults))
	require.Equal(t, 1, len(blockResults.TxsResults[0].Events))
	require.Equal(t, eventType, blockResults.TxsResults[0].Events[0].Type)
	require.Equal(t, 1, len(blockResults.TxsResults[0].Events[0].Attributes))
	require.Equal(t, kvPairKey, blockResults.TxsResults[0].Events[0].Attributes[0].Key)
	require.Equal(t, 1, len(blockResults.BeginBlockEvents))
	require.Equal(t, eventType, blockResults.BeginBlockEvents[0].Type)
	require.Equal(t, 1, len(blockResults.BeginBlockEvents[0].Attributes))
	require.Equal(t, kvPairKey, blockResults.BeginBlockEvents[0].Attributes[0].Key)
	require.Equal(t, 1, len(blockResults.EndBlockEvents))
	require.Equal(t, eventType, blockResults.EndBlockEvents[0].Type)
	require.Equal(t, 1, len(blockResults.EndBlockEvents[0].Attributes))
	require.Equal(t, kvPairKey, blockResults.EndBlockEvents[0].Attributes[0].Key)
	require.Equal(t, 1, len(blockResults.ValidatorUpdates))
	require.Equal(t, power, blockResults.ValidatorUpdates[0].Power)
	require.Equal(t, pubkeyType, blockResults.ValidatorUpdates[0].PubKey.Type)

	mockCli.EXPECT().BlockResults(gomock.AssignableToTypeOf(&height)).Return(nil, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryBlockResults(height)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryBlock(-1)
	require.Error(t, err)
}

func TestTendermintClient_QueryCommitResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	height, blockTime := int64(1024), time.Now()
	appHash, blockIDHash := tmbytes.HexBytes("default app hash"), tmbytes.HexBytes("default block ID hash")

	expectedRet := mockCli.GetRawCommitResultPointer(true, "default chainID", height, blockTime, appHash, blockIDHash)
	mockCli.EXPECT().Commit(gomock.AssignableToTypeOf(&height)).Return(expectedRet, nil)

	commitResult, err := mockCli.Tendermint().QueryCommitResult(height)
	require.NoError(t, err)
	require.Equal(t, true, commitResult.CanonicalCommit)
	require.Equal(t, "default chainID", commitResult.ChainID)
	require.Equal(t, appHash, commitResult.AppHash)
	require.Equal(t, height, commitResult.Header.Height)
	require.Equal(t, blockIDHash, commitResult.Commit.BlockID.Hash)
	require.True(t, blockTime.Equal(commitResult.Time))

	mockCli.EXPECT().Commit(gomock.AssignableToTypeOf(&height)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryCommitResult(height)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryCommitResult(-1)
	require.Error(t, err)
}

func TestTendermintClient_QueryValidatorsResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	height, votingPower, proposerPriority := int64(1024), int64(2048), int64(-1024)
	okexchain.SetBech32Prefixes(sdk.GetConfig())
	consPubkey, err := stakingtypes.GetConsPubKeyBech32(valConsPK)
	require.NoError(t, err)

	expectedRet := mockCli.GetRawValidatorsResultPointer(height, votingPower, proposerPriority, consPubkey)
	mockCli.EXPECT().Validators(gomock.AssignableToTypeOf(&height), gomock.AssignableToTypeOf(0), gomock.AssignableToTypeOf(0)).
		Return(expectedRet, nil)

	valsResult, err := mockCli.Tendermint().QueryValidatorsResult(height)
	require.NoError(t, err)
	require.Equal(t, height, valsResult.BlockHeight)
	require.Equal(t, proposerPriority, valsResult.Validators[0].ProposerPriority)
	require.Equal(t, votingPower, valsResult.Validators[0].VotingPower)
	require.Equal(t, consPubkey, valsResult.Validators[0].PubKey)

	mockCli.EXPECT().Validators(gomock.AssignableToTypeOf(&height), gomock.AssignableToTypeOf(0), gomock.AssignableToTypeOf(0)).
		Return(nil, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryValidatorsResult(height)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryValidatorsResult(-1)
	require.Error(t, err)
}

func TestTendermintClient_QueryTxResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	hashHexStr := "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	expectedHashBytes, err := hex.DecodeString(hashHexStr)
	require.NoError(t, err)
	tx := []byte("default tx")
	height, code := int64(1024), uint32(0)
	log, eventType := "default log", "default event type"

	expectedRet := mockCli.GetRawTxResultPointer(height, code, log, hashHexStr, eventType, tx)
	mockCli.EXPECT().Tx(expectedHashBytes, true).Return(expectedRet, nil)

	txResult, err := mockCli.Tendermint().QueryTxResult(hashHexStr, true)
	require.NoError(t, err)
	require.Equal(t, height, txResult.Height)
	require.Equal(t, tmbytes.HexBytes(expectedHashBytes), txResult.Hash)
	require.Equal(t, tmtypes.Tx(tx), txResult.Tx)
	require.Equal(t, log, txResult.TxResult.Log)
	require.Equal(t, code, txResult.TxResult.Code)
	require.Equal(t, eventType, txResult.TxResult.Events[0].Type)

	mockCli.EXPECT().Tx(expectedHashBytes, true).Return(nil, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryTxResult(hashHexStr, true)
	require.Error(t, err)
}

func TestTendermintClient_QueryTxsByEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTendermintClient(mockCli.MockBaseClient))

	hashHexStr := "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	expectedHashBytes, err := hex.DecodeString(hashHexStr)
	tx := []byte("default tx")
	height, code := int64(1024), uint32(0)
	log, eventType := "default log", "default event type"

	expectedRet := mockCli.GetRawResultTxSearchPointer(1, height, code, log, hashHexStr, eventType, tx)
	mockCli.EXPECT().TxSearch(gomock.AssignableToTypeOf(""), false, gomock.AssignableToTypeOf(0),
		gomock.AssignableToTypeOf(0), gomock.AssignableToTypeOf("")).Return(expectedRet, nil)

	queryStr := fmt.Sprintf("message.sender=%s", addr)
	txSearchResult, err := mockCli.Tendermint().QueryTxsByEvents(queryStr, 1, 30)
	require.NoError(t, err)
	require.Equal(t, 1, txSearchResult.TotalCount)
	require.Equal(t, height, txSearchResult.Txs[0].Height)
	require.Equal(t, tmbytes.HexBytes(expectedHashBytes), txSearchResult.Txs[0].Hash)
	require.Equal(t, tmtypes.Tx(tx), txSearchResult.Txs[0].Tx)
	require.Equal(t, log, txSearchResult.Txs[0].TxResult.Log)
	require.Equal(t, code, txSearchResult.Txs[0].TxResult.Code)
	require.Equal(t, eventType, txSearchResult.Txs[0].TxResult.Events[0].Type)

	mockCli.EXPECT().TxSearch(gomock.AssignableToTypeOf(""), false, gomock.AssignableToTypeOf(0),
		gomock.AssignableToTypeOf(0), gomock.AssignableToTypeOf("")).Return(nil, errors.New("default error"))
	_, err = mockCli.Tendermint().QueryTxsByEvents(queryStr, 1, 30)
	require.Error(t, err)

	badQueryStr := fmt.Sprintf("message.sender%s", addr)
	_, err = mockCli.Tendermint().QueryTxsByEvents(badQueryStr, 1, 30)
	require.Error(t, err)

	badQueryStr = fmt.Sprintf("message.sender==%s", addr)
	_, err = mockCli.Tendermint().QueryTxsByEvents(badQueryStr, 1, 30)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryTxsByEvents(queryStr, -1, 30)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryTxsByEvents(queryStr, 1, -30)
	require.Error(t, err)

	_, err = mockCli.Tendermint().QueryTxsByEvents("", 1, 30)
	require.Error(t, err)
}
