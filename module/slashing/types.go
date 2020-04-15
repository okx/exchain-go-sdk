package slashing

import "github.com/okex/okchain-go-sdk/types"

// const
const (
	ModuleName = "slashing"
)

var (
	msgCdc = types.NewCodec()
)

func init() {
	registerCodec(msgCdc)
}
