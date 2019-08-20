package types

import (
	abci "github.com/ok-chain/ok-gosdk/types/abci"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/crypto/merkle"


)

type Tx []byte

type Txs []Tx


// One usage is indexing transaction results.
type TxResult struct {
	Height int64                  `json:"height"`
	Index  uint32                 `json:"index"`
	Tx     Tx                     `json:"tx"`
	Result abci.ResponseDeliverTx `json:"result"`
}

type TxProof struct {
	RootHash cmn.HexBytes
	Data     Tx
	Proof    merkle.SimpleProof
}
