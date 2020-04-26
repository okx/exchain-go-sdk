package utils

import (
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	cmn "github.com/tendermint/tendermint/libs/common"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmstate "github.com/tendermint/tendermint/state"
	tmtypes "github.com/tendermint/tendermint/types"
	"testing"
	"time"
)

var testCdc sdk.SDKCodec

func init() {
	testCdc = sdk.NewCodec()
	sdk.RegisterBasicCodec(testCdc)
	testCdc.RegisterConcrete(TestMsg{}, "testing/TestMsg")
	testCdc.Seal()
}

// mock msg for the testing
var _ sdk.Msg = (*TestMsg)(nil)

type TestMsg struct {
	sdk.AccAddress
}

func (TestMsg) Route() string { return "" }
func (TestMsg) Type() string  { return "" }
func (TestMsg) ValidateBasic() sdk.Error {
	return nil
}
func (TestMsg) GetSigners() []sdk.AccAddress {
	return nil
}
func (msg TestMsg) GetSignBytes() []byte {
	return sdk.MustSortJSON(testCdc.MustMarshalJSON(msg))
}

func TestParseBlock(t *testing.T) {
	// data preparation
	height, blockTime := int64(1024), time.Now()
	appHash, blockIDHash := cmn.HexBytes("default app hash"), cmn.HexBytes("default block ID hash")
	chainID := "default chainID"
	addr, err := sdk.AccAddressFromBech32(accAddr1)
	require.NoError(t, err)
	feeCoins, err := sdk.ParseDecCoins("1024okt,2.048btc")
	require.NoError(t, err)

	stdFee := sdk.NewStdFee(20000, feeCoins)
	stdTx := sdk.StdTx{
		Msgs: []sdk.Msg{
			TestMsg{
				addr,
			},
		},
		Fee:  stdFee,
		Memo: defaultMemo,
	}

	pTmBlock := &tmtypes.Block{
		Header: tmtypes.Header{
			ChainID: chainID,
			Height:  height,
			Time:    blockTime,
			AppHash: appHash,
		},
		Data: tmtypes.Data{
			Txs: tmtypes.Txs{
				testCdc.MustMarshalBinaryLengthPrefixed(stdTx),
			},
		},
		Evidence: tmtypes.EvidenceData{},
		LastCommit: &tmtypes.Commit{
			BlockID: tmtypes.BlockID{
				Hash: blockIDHash,
			},
		},
	}

	block, err := ParseBlock(testCdc, pTmBlock)
	require.NoError(t, err)
	require.Equal(t, chainID, block.ChainID)
	require.Equal(t, height, block.Height)
	require.Equal(t, appHash, block.AppHash)
	require.Equal(t, blockIDHash, block.LastCommit.BlockID.Hash)
	require.Equal(t, 1, len(block.Txs))
	require.Equal(t, stdTx, block.Txs[0])
	require.True(t, blockTime.Equal(block.Time))

	// bad stdTx bytes
	pTmBlock.Txs[0] = []byte("common string but not stdTx amino bytes")
	_, err = ParseBlock(testCdc, pTmBlock)
	require.Error(t, err)
}

func TestParseBlockResults(t *testing.T) {
	// data preparation
	power, height := int64(1000), int64(1024)
	pubkeyType, eventType := "default pubkey type", "default event type"
	kvPairKey := []byte("default kv pair key")

	pTmBlockResults := &ctypes.ResultBlockResults{
		Height: height,
		Results: &tmstate.ABCIResponses{
			BeginBlock: &abci.ResponseBeginBlock{
				Events: []abci.Event{
					{
						Type: eventType,
						Attributes: []common.KVPair{
							{
								Key: kvPairKey,
							},
						},
					},
				},
			},
			EndBlock: &abci.ResponseEndBlock{
				ValidatorUpdates: []abci.ValidatorUpdate{
					{
						PubKey: abci.PubKey{
							Type: pubkeyType,
						},
						Power: power,
					},
				},
			},
		},
	}

	blockResults := ParseBlockResults(pTmBlockResults)
	require.Equal(t, height, blockResults.Height)
	require.Equal(t, 1, len(blockResults.Results.BeginBlock.Events))
	require.Equal(t, eventType, blockResults.Results.BeginBlock.Events[0].Type)
	require.Equal(t, kvPairKey, blockResults.Results.BeginBlock.Events[0].Attributes[0].Key)
	require.Equal(t, pubkeyType, blockResults.Results.EndBlock.ValidatorUpdates[0].PubKey.Type)
	require.Equal(t, power, blockResults.Results.EndBlock.ValidatorUpdates[0].Power)
}
