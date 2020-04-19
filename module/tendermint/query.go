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

// QueryBlockResults gets the abci result of the block on a specific height
func (tc tendermintClient) QueryBlockResults(height int64) (blockResults types.BlockResults, err error) {
	pTmBlockResults, err := tc.BlockResults(&height)
	if err != nil {
		return
	}

	return utils.ParseBlockResults(pTmBlockResults), err
}

// QueryCommitResult gets the commit info of the block on a specific height
func (tc tendermintClient) QueryCommitResult(height int64) (commitResult types.ResultCommit, err error) {
	pTmCommitResult, err := tc.Commit(&height)
	if err != nil {
		return
	}

	return utils.ParseCommitResult(pTmCommitResult), err
}
