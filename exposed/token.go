package exposed

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/token/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// Token shows the expected behavior for inner token client
type Token interface {
	gosdktypes.Module
	TokenTx
	TokenQuery
}

// TokenTx shows the expected tx behavior for inner token client
type TokenTx interface {
	Send(fromInfo keys.Info, passWd, toAddrStr, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	MultiSend(fromInfo keys.Info, passWd string, transfers []types.TransferUnit, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	Issue(fromInfo keys.Info, passWd, orgSymbol, wholeName, totalSupply, tokenDesc, memo string, mintable bool, accNum,
		seqNum uint64) (sdk.TxResponse, error)
	Mint(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Burn(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Edit(fromInfo keys.Info, passWd, symbol, description, wholeName, memo string, isDescEdit, isWholeNameEdit bool, accNum,
		seqNum uint64) (sdk.TxResponse, error)
}

// TokenQuery shows the expected query behavior for inner token client
type TokenQuery interface {
	QueryTokenInfo(ownerAddr, symbol string) ([]types.Token, error)
	QueryAccountTokensInfo(addrStr string) (types.AccountTokensInfo, error)
	QueryAccountTokenInfo(addrStr, symbol string) (types.AccountTokensInfo, error)
}
