package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// const
const (
	ModuleName = "tendermint"
)

// Block - structure for the result of block query
type Block struct {
	tmtypes.Header `json:"header"`
	Data           `json:"data"`
	Evidence       tmtypes.EvidenceData `json:"evidence"`
	LastCommit     *tmtypes.Commit      `json:"last_commit"`
}

// NewBlock creates a new instance of Block
func NewBlock(header tmtypes.Header, data Data, evidence tmtypes.EvidenceData, pLastCommit *tmtypes.Commit) Block {
	return Block{
		Header:     header,
		Data:       data,
		Evidence:   evidence,
		LastCommit: pLastCommit,
	}
}

// Data - structure of the stdTxs in a block
type Data struct {
	Txs []sdk.StdTx `json:"txs"`
}

// NewData creates a new instance of Data
func NewData(stdTxs []sdk.StdTx) Data {
	return Data{
		Txs: stdTxs,
	}
}
