package tendermint

import (
	"github.com/okex/okchain-go-sdk/module/tendermint/types"
	"github.com/okex/okchain-go-sdk/utils"
)

// QueryBlock gets the block info of a specific height
func (tc tendermintClient) QueryBlock(height int64) (block types.Block, err error) {
	pTmBlockResult, err := tc.Block(&height)
	if err != nil {
		return
	}

	return utils.ParseBlock(tc.GetCodec(), pTmBlockResult.Block)
}

func (tc tendermintClient) QueryBlockResults(height int64) (blockResults types.BlockResults, err error) {
	pTmBlockResults, err := tc.BlockResults(&height)
	if err != nil {
		return
	}

	return utils.ParseBlockResults(pTmBlockResults), err
}
