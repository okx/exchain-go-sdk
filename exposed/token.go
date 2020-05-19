package exposed

import (
	"github.com/okex/okchain-go-sdk/module/token/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
)

// Token shows the expected behavior for inner token client
type Token interface {
	sdk.Module
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
}

// TokenQuery shows the expected query behavior for inner token client
type TokenQuery interface {
	QueryTokenInfo(ownerAddr, symbol string) ([]types.Token, error)
	QueryAccountTokensInfo(addrStr string) (types.AccountTokensInfo, error)
	QueryAccountTokenInfo(addrStr, symbol string) (types.AccountTokensInfo, error)
}
