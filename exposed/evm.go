package exposed

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/evm/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// Evm shows the expected behavior for inner farm client
type Evm interface {
	gosdktypes.Module
	EvmTx
	EvmQuery
}

// EvmTx shows the expected tx behavior for inner evm client
type EvmTx interface {
	SendTx(fromInfo keys.Info, passWd, toAddrStr, amountStr, payloadStr, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	CreateContract(fromInfo keys.Info, passWd, amountStr, payloadStr, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, string, error)
	SendTxEthereum(privHex, toAddrStr, amountStr, payloadStr string, seqNum uint64) (sdk.TxResponse, error)
}

// EvmQuery shows the expected query behavior for inner evm client
type EvmQuery interface {
	QueryCode(contractAddrStr string) (types.QueryResCode, error)
	QueryStorageAt(contractAddrStr, keyHexStr string) (types.QueryResStorage, error)
}
