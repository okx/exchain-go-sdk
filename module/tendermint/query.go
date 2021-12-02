package tendermint

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/okex/exchain-go-sdk/module/tendermint/types"
	"github.com/okex/exchain-go-sdk/types/params"
	tmtypes "github.com/okex/exchain/libs/tendermint/types"
)

// QueryBlock gets the block info of a specific height
// query the latest block with height 0 input
func (tc tendermintClient) QueryBlock(height int64) (pBlock *types.Block, err error) {
	if err = params.CheckQueryHeightParams(height); err != nil {
		return pBlock, err
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
func (tc tendermintClient) QueryBlockResults(height int64) (pBlockResults *types.ResultBlockResults, err error) {
	if err = params.CheckQueryHeightParams(height); err != nil {
		return pBlockResults, err
	}

	var pHeight *int64
	if height > 0 {
		pHeight = &height
	}

	return tc.BlockResults(pHeight)
}

// QueryCommitResult gets the commit info of the block on a specific height
// query the latest block with height 0 input
func (tc tendermintClient) QueryCommitResult(height int64) (pCommitResult *types.ResultCommit, err error) {
	if err = params.CheckQueryHeightParams(height); err != nil {
		return pCommitResult, err
	}

	var pHeight *int64
	if height > 0 {
		pHeight = &height
	}

	return tc.Commit(pHeight)
}

// QueryValidatorsResult gets the validators info on a specific height
// query the latest block with height 0 input
func (tc tendermintClient) QueryValidatorsResult(height int64) (pValsResult *types.ResultValidators, err error) {
	if err = params.CheckQueryHeightParams(height); err != nil {
		return pValsResult, err
	}

	var pHeight *int64
	if height > 0 {
		pHeight = &height
	}

	return tc.Validators(pHeight, 1, 0)
}

// QueryTxResult gets the detail info of a tx with its tx hash
func (tc tendermintClient) QueryTxResult(hashHexStr string, prove bool) (pResultTx *types.ResultTx, err error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return
	}

	return tc.Tx(hash, prove)
}

// QueryTxsByEvents gets txs result by a group of specific searching string
// NOTE: it assumes the node to query a truth teller
func (tc tendermintClient) QueryTxsByEvents(eventsStr string, page, limit int) (pResultTxSearch *types.ResultTxSearch, err error) {
	// parse the eventsStr
	var events []string
	if strings.Contains(eventsStr, "&") {
		events = strings.Split(eventsStr, "&")
	} else {
		events = append(events, eventsStr)
	}

	// build tm events
	var tmEvents []string
	for _, event := range events {
		if !strings.Contains(event, "=") {
			return pResultTxSearch, fmt.Errorf("invalid event; event %s should be of the format: %s", event, types.EventFormat)
		} else if strings.Count(event, "=") > 1 {
			return pResultTxSearch, fmt.Errorf("invalid event; event %s should be of the format: %s", event, types.EventFormat)
		}

		tokens := strings.Split(event, "=")
		if tokens[0] == tmtypes.TxHeightKey {
			event = fmt.Sprintf("%s=%s", tokens[0], tokens[1])
		} else {
			event = fmt.Sprintf("%s='%s'", tokens[0], tokens[1])
		}

		tmEvents = append(tmEvents, event)
	}

	if len(tmEvents) == 0 {
		return pResultTxSearch, errors.New("must declare at least one event to search")
	}

	if page <= 0 {
		return pResultTxSearch, errors.New("page must greater than 0")
	}

	if limit <= 0 {
		return pResultTxSearch, errors.New("limit must greater than 0")
	}

	// XXX: implement ANY
	query := strings.Join(tmEvents, " AND ")
	// assumes the node to query a truth teller
	return tc.TxSearch(query, false, page, limit, "")
}
