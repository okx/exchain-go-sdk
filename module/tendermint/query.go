package tendermint

import (
	"github.com/okex/okchain-go-sdk/module/tendermint/types"
	"github.com/okex/okchain-go-sdk/utils"
)

// QueryBlock gets the block info of a specific height
func (tb tendermintClient) QueryBlock(height int64) (block types.Block, err error) {
	pBlockResult, err := tb.Block(&height)
	if err != nil {
		return
	}

	return utils.ParseBlock(tb.GetCodec(),pBlockResult.Block)
}
