package utils

import (
	//"testing"
	//"time"
	//
	//sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/stretchr/testify/require"
	//abci "github.com/tendermint/tendermint/abci/types"
	//"github.com/tendermint/tendermint/libs/common"
	//cmn "github.com/tendermint/tendermint/libs/common"
	//ctypes "github.com/tendermint/tendermint/rpc/core/types"
	//tmstate "github.com/tendermint/tendermint/state"
	//tmtypes "github.com/tendermint/tendermint/types"
)
//
//var testCdc sdk.SDKCodec
//
//func init() {
//	testCdc = sdk.NewCodec()
//	sdk.RegisterBasicCodec(testCdc)
//	testCdc.RegisterConcrete(TestMsg{}, "testing/TestMsg")
//	testCdc.Seal()
//}
//
//// mock msg for the testing
//var _ sdk.Msg = (*TestMsg)(nil)
//
//type TestMsg struct {
//	sdk.AccAddress
//}
//
//func (TestMsg) Route() string { return "" }
//func (TestMsg) Type() string  { return "" }
//func (TestMsg) ValidateBasic() sdk.Error {
//	return nil
//}
//func (TestMsg) GetSigners() []sdk.AccAddress {
//	return nil
//}
//func (msg TestMsg) GetSignBytes() []byte {
//	return sdk.MustSortJSON(testCdc.MustMarshalJSON(msg))
//}
//
//func TestParseBlock(t *testing.T) {
//	// data preparation
//	height, blockTime := int64(1024), time.Now()
//	appHash, blockIDHash := cmn.HexBytes("default app hash"), cmn.HexBytes("default block ID hash")
//	chainID := "default chainID"
//	addr, err := sdk.AccAddressFromBech32(accAddr1)
//	require.NoError(t, err)
//	feeCoins, err := sdk.ParseDecCoins("1024okt,2.048btc")
//	require.NoError(t, err)
//	stdFee := sdk.NewStdFee(20000, feeCoins)
//
//	stdTx := sdk.StdTx{
//		Msgs: []sdk.Msg{
//			TestMsg{
//				addr,
//			},
//		},
//		Fee:  stdFee,
//		Memo: defaultMemo,
//	}
//
//	pTmBlock := &tmtypes.Block{
//		Header: tmtypes.Header{
//			ChainID: chainID,
//			Height:  height,
//			Time:    blockTime,
//			AppHash: appHash,
//		},
//		Data: tmtypes.Data{
//			Txs: tmtypes.Txs{
//				testCdc.MustMarshalBinaryLengthPrefixed(stdTx),
//			},
//		},
//		Evidence: tmtypes.EvidenceData{},
//		LastCommit: &tmtypes.Commit{
//			BlockID: tmtypes.BlockID{
//				Hash: blockIDHash,
//			},
//		},
//	}
//
//	block, err := ParseBlock(testCdc, pTmBlock)
//	require.NoError(t, err)
//	require.Equal(t, chainID, block.ChainID)
//	require.Equal(t, height, block.Height)
//	require.Equal(t, appHash, block.AppHash)
//	require.Equal(t, blockIDHash, block.LastCommit.BlockID.Hash)
//	require.Equal(t, 1, len(block.Txs))
//	require.Equal(t, stdTx, block.Txs[0])
//	require.True(t, blockTime.Equal(block.Time))
//
//	// bad stdTx bytes
//	pTmBlock.Txs[0] = []byte("common string but not stdTx amino bytes")
//	_, err = ParseBlock(testCdc, pTmBlock)
//	require.Error(t, err)
//}
//
//func TestParseBlockResults(t *testing.T) {
//	// data preparation
//	power, height := int64(1000), int64(1024)
//	pubkeyType, eventType, info := "default pubkey type", "default event type", "default info"
//	kvPairKey := []byte("default kv pair key")
//	maxBytes, maxGas, maxAge := int64(1024), int64(20000), int64(2048)
//
//	pTmBlockResults := &ctypes.ResultBlockResults{
//		Height: height,
//		Results: &tmstate.ABCIResponses{
//			DeliverTx: []*abci.ResponseDeliverTx{
//				{
//					Info: info,
//				},
//			},
//			BeginBlock: &abci.ResponseBeginBlock{
//				Events: []abci.Event{
//					{
//						Type: eventType,
//						Attributes: []common.KVPair{
//							{
//								Key: kvPairKey,
//							},
//						},
//					},
//				},
//			},
//			EndBlock: &abci.ResponseEndBlock{
//				ValidatorUpdates: []abci.ValidatorUpdate{
//					{
//						PubKey: abci.PubKey{
//							Type: pubkeyType,
//						},
//						Power: power,
//					},
//				},
//				ConsensusParamUpdates: &abci.ConsensusParams{
//					Block: &abci.BlockParams{
//						MaxBytes: maxBytes,
//						MaxGas:   maxGas,
//					},
//					Evidence: &abci.EvidenceParams{
//						MaxAge: maxAge,
//					},
//					Validator: &abci.ValidatorParams{
//						PubKeyTypes: []string{pubkeyType},
//					},
//				},
//			},
//		},
//	}
//
//	blockResults := ParseBlockResults(pTmBlockResults)
//	require.Equal(t, height, blockResults.Height)
//	require.Equal(t, 1, len(blockResults.Results.DeliverTx))
//	require.Equal(t, info, blockResults.Results.DeliverTx[0].Info)
//	require.Equal(t, 1, len(blockResults.Results.BeginBlock.Events))
//	require.Equal(t, eventType, blockResults.Results.BeginBlock.Events[0].Type)
//	require.Equal(t, kvPairKey, blockResults.Results.BeginBlock.Events[0].Attributes[0].Key)
//	require.Equal(t, pubkeyType, blockResults.Results.EndBlock.ValidatorUpdates[0].PubKey.Type)
//	require.Equal(t, power, blockResults.Results.EndBlock.ValidatorUpdates[0].Power)
//	require.Equal(t, maxBytes, blockResults.Results.EndBlock.ConsensusParamUpdates.Block.MaxBytes)
//	require.Equal(t, maxGas, blockResults.Results.EndBlock.ConsensusParamUpdates.Block.MaxGas)
//	require.Equal(t, maxAge, blockResults.Results.EndBlock.ConsensusParamUpdates.Evidence.MaxAge)
//	require.Equal(t, 1, len(blockResults.Results.EndBlock.ConsensusParamUpdates.Validator.PubKeyTypes))
//	require.Equal(t, pubkeyType, blockResults.Results.EndBlock.ConsensusParamUpdates.Validator.PubKeyTypes[0])
//
//	// consensus params check: empty pointer
//	pTmBlockResults.Results.EndBlock.ConsensusParamUpdates = nil
//	blockResults = ParseBlockResults(pTmBlockResults)
//	require.NotNil(t, blockResults.Results.EndBlock.ConsensusParamUpdates)
//}
//
//func TestParseCommitResult(t *testing.T) {
//	// data preparation
//	height, blockTime := int64(1024), time.Now()
//	appHash, blockIDHash := cmn.HexBytes("default app hash"), cmn.HexBytes("default block ID hash")
//	chainID := "default chainID"
//
//	pTmCommitResult := &ctypes.ResultCommit{
//		CanonicalCommit: true,
//		SignedHeader: tmtypes.SignedHeader{
//			Header: &tmtypes.Header{
//				ChainID: chainID,
//				Height:  height,
//				Time:    blockTime,
//				AppHash: appHash,
//			},
//			Commit: &tmtypes.Commit{
//				BlockID: tmtypes.BlockID{
//					Hash: blockIDHash,
//				},
//			},
//		},
//	}
//
//	commitResult := ParseCommitResult(pTmCommitResult)
//	require.Equal(t, true, commitResult.CanonicalCommit)
//	require.Equal(t, chainID, commitResult.ChainID)
//	require.Equal(t, appHash, commitResult.AppHash)
//	require.Equal(t, height, commitResult.Header.Height)
//	require.Equal(t, blockIDHash, commitResult.Commit.BlockID.Hash)
//	require.True(t, blockTime.Equal(commitResult.Time))
//}
//
//func TestParseValidatorsResult(t *testing.T) {
//	// data preparation
//	height, votingPower, proposerPriority := int64(1024), int64(2048), int64(-1024)
//	consPubkey, err := sdk.GetConsPubKeyBech32(valConsPK)
//	require.NoError(t, err)
//
//	pTmValsResult := &ctypes.ResultValidators{
//		BlockHeight: height,
//		Validators: []*tmtypes.Validator{
//			{
//				PubKey:           consPubkey,
//				VotingPower:      votingPower,
//				ProposerPriority: proposerPriority,
//			},
//		},
//	}
//
//	valsResult := ParseValidatorsResult(pTmValsResult)
//	require.Equal(t, height, valsResult.BlockHeight)
//	require.Equal(t, 1, len(valsResult.Validators))
//	require.Equal(t, proposerPriority, valsResult.Validators[0].ProposerPriority)
//	require.Equal(t, votingPower, valsResult.Validators[0].VotingPower)
//	require.Equal(t, consPubkey, valsResult.Validators[0].PubKey)
//}
//
//func TestParseTxResult(t *testing.T) {
//	// data preparation
//	txHash, tx := []byte("default tx hash"), []byte("default tx")
//	height, code := int64(1024), uint32(0)
//	log, eventType := "default log", "default event type"
//
//	pTmTxResult := &ctypes.ResultTx{
//		Hash:   txHash,
//		Height: height,
//		Tx:     tx,
//		TxResult: abci.ResponseDeliverTx{
//			Code: code,
//			Log:  log,
//			Events: []abci.Event{
//				{
//					Type: eventType,
//				},
//			},
//		},
//	}
//
//	txResult := ParseTxResult(pTmTxResult)
//	require.Equal(t, height, txResult.Height)
//	require.Equal(t, cmn.HexBytes(txHash), txResult.Hash)
//	require.Equal(t, tmtypes.Tx(tx), txResult.Tx)
//	require.Equal(t, log, txResult.TxResult.Log)
//	require.Equal(t, code, txResult.TxResult.Code)
//	require.Equal(t, 1, len(txResult.TxResult.Events))
//	require.Equal(t, eventType, txResult.TxResult.Events[0].Type)
//}
//
//func TestParseTxsResult(t *testing.T) {
//	// data preparation
//	txHash, tx := []byte("default tx hash"), []byte("default tx")
//	height, code, totalCount := int64(1024), uint32(0), 1
//	log, eventType := "default log", "default event type"
//
//	pTmTxsResult := &ctypes.ResultTxSearch{
//		TotalCount: totalCount,
//		Txs: []*ctypes.ResultTx{
//			{
//				Hash:   txHash,
//				Height: height,
//				Tx:     tx,
//				TxResult: abci.ResponseDeliverTx{
//					Code: code,
//					Log:  log,
//					Events: []abci.Event{
//						{
//							Type: eventType,
//						},
//					},
//				},
//			},
//		},
//	}
//
//	txSearchResult := ParseTxsResult(pTmTxsResult)
//	require.Equal(t, totalCount, txSearchResult.TotalCount)
//	require.Equal(t, 1, len(txSearchResult.Txs))
//	require.Equal(t, height, txSearchResult.Txs[0].Height)
//	require.Equal(t, cmn.HexBytes(txHash), txSearchResult.Txs[0].Hash)
//	require.Equal(t, tmtypes.Tx(tx), txSearchResult.Txs[0].Tx)
//	require.Equal(t, log, txSearchResult.Txs[0].TxResult.Log)
//	require.Equal(t, code, txSearchResult.Txs[0].TxResult.Code)
//	require.Equal(t, eventType, txSearchResult.Txs[0].TxResult.Events[0].Type)
//}
