package utils

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/module/tendermint/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// ParseBlock converts raw tendermint block type to the one gosdk requires
func ParseBlock(cdc sdk.SDKCodec, pTmBlock *tmtypes.Block) (block types.Block, err error) {
	var stdTxs []sdk.StdTx
	for _, txBytes := range pTmBlock.Txs {
		var stdTx sdk.StdTx
		if err = cdc.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx); err != nil {
			return block, fmt.Errorf("failed. unmarshal tx info from tendermint block query error: %s", err)
		}
		stdTxs = append(stdTxs, stdTx)
	}

	return types.NewBlock(pTmBlock.Header, types.NewData(stdTxs), pTmBlock.Evidence, pTmBlock.LastCommit), err
}
