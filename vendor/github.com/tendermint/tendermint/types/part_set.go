package types

import (
	cmn "github.com/tendermint/tendermint/libs/common"
)


type PartSetHeader struct {
	Total int          `json:"total"`
	Hash  cmn.HexBytes `json:"hash"`
}
