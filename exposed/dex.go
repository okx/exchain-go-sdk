package exposed

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/dex/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// Dex shows the expected behavior for inner dex client
type Dex interface {
	gosdktypes.Module
	DexTx
	DexQuery
	DexOffline
}

// DexTx shows the expected tx behavior for inner dex client
type DexTx interface {
	List(fromInfo keys.Info, passWd, baseAsset, quoteAsset, initPriceStr, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	Deposit(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Withdraw(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	TransferOwnership(fromInfo keys.Info, passWd, inputPath string, accNum, seqNum uint64) (sdk.TxResponse, error)
	RegisterDexOperator(fromInfo keys.Info, passWd, handleFeeAddrStr, website, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	EditDexOperator(fromInfo keys.Info, passWd, handleFeeAddrStr, website, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}

// DexOffline shows the expected tx behavior offline for inner dex client
type DexOffline interface {
	GenerateUnsignedTransferOwnershipTx(product, fromAddrStr, toAddrStr, memo, outputPath string) error
	MultiSign(fromInfo keys.Info, passWd, inputPath, outputPath string) error
}

// DexQuery shows the expected query behavior for inner dex client
type DexQuery interface {
	QueryProducts(ownerAddr string, page, perPage int) ([]types.TokenPair, error)
}
