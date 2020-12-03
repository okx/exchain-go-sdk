package tendermint

import (
	"errors"
	"fmt"
	"github.com/okex/okexchain-go-sdk/module/tendermint/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	"github.com/okex/okexchain-go-sdk/utils"
	tmtypes "github.com/tendermint/tendermint/types"
	"strings"
)

// QueryBlock gets the block info of a specific height
// query the latest block with height 0 input
func (tc tendermintClient) QueryBlock(height int64) (pBlock *types.Block, err error) {
	if height < 0 {
		return pBlock, errors.New("failed. negative height is not available")
	}

	var pHeight *int64
	if height > 0 {
		pHeight = &height
	}

	pTmBlockResult, err := tc.Block(pHeight)
	if err != nil {
		return
	}

	return pTmBlockResult.Block, err
}

// QueryBlockResults gets the abci result of the block on a specific height
// query the latest block with height 0 input
func (tc tendermintClient) QueryBlockResults(height int64) (blockResults *types.ResultBlockResults, err error) {
	if height < 0 {
		return blockResults, errors.New("failed. negative height is not available")
	}

	var pHeight *int64
	if height > 0 {
		pHeight = &height
	}

	return tc.BlockResults(pHeight)
}

// QueryCommitResult gets the commit info of the block on a specific height
func (tc tendermintClient) QueryCommitResult(height int64) (commitResult types.ResultCommit, err error) {
	pTmCommitResult, err := tc.Commit(&height)
	if err != nil {
		return
	}

	return utils.ParseCommitResult(pTmCommitResult), err
}

// QueryValidatorsResult gets the validators info on a specific height
func (tc tendermintClient) QueryValidatorsResult(height int64) (valsResult types.ResultValidators, err error) {
	//pTmValsResult, err := tc.Validators(&height)
	// TODO
	pTmValsResult, err := tc.Validators(&height, 1, 0)
	if err != nil {
		return
	}

	return utils.ParseValidatorsResult(pTmValsResult), err
}

// QueryTxResult gets the detail info of a tx with its tx hash
func (tc tendermintClient) QueryTxResult(txHash []byte, prove bool) (txResult types.ResultTx, err error) {
	pTmTxResult, err := tc.Tx(txHash, prove)
	if err != nil {
		return
	}

	return utils.ParseTxResult(pTmTxResult), err
}

// QueryTxsResult gets txs result by a specific searching string
// NOTE: QueryTxsResult assumes the node telling truth
func (tc tendermintClient) QueryTxsResult(searchStr string, page, perPage int) (txsResult types.ResultTxs,
	err error) {

	tmEventStrs, err := parseSearchingStr(searchStr)
	if err != nil {
		return
	}

	if err = params.CheckQueryTxResultParams(tmEventStrs, page, perPage); err != nil {
		return
	}

	queryStr := strings.Join(tmEventStrs, " AND ")
	// TODO
	pTmTxsResult, err := tc.TxSearch(queryStr, false, page, perPage, "")
	if err != nil {
		return
	}

	return utils.ParseTxsResult(pTmTxsResult), err
}

func parseSearchingStr(searchStr string) (tmEventStrs []string, err error) {
	var events []string
	searchStr = strings.TrimSpace(searchStr)
	if strings.Contains(searchStr, "&") {
		events = strings.Split(searchStr, "&")
	} else {
		events = append(events, searchStr)
	}

	for _, event := range events {
		if !strings.Contains(event, "=") || strings.Count(event, "=") > 1 {
			return tmEventStrs, fmt.Errorf("failed. event %s should be of the format: %s", event, types.EventFormat)
		}

		tokens := strings.Split(event, "=")
		if tokens[0] == tmtypes.TxHeightKey {
			event = fmt.Sprintf("%s=%s", tokens[0], tokens[1])
		} else {
			event = fmt.Sprintf("%s='%s'", tokens[0], tokens[1])
		}

		tmEventStrs = append(tmEventStrs, event)
	}

	return
}
