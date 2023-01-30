package exposed

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
)

type Icamauth interface {
	gosdktypes.Module
	SubmitTx(fromInfo keys.Info, passWd, connectionID,
		memo string, data []byte, accNum, seqNum uint64) (resp sdk.TxResponse, err error)
}
