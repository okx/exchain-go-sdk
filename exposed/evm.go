package exposed

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// Evm shows the expected behavior for inner farm client
type Evm interface {
	gosdktypes.Module
	EvmTx
}

// EvmTx shows the expected tx behavior for inner evm client
type EvmTx interface {
	SendTx(fromInfo keys.Info, passWd, toAddrStr, amountStr, payloadStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}
