package exposed

import (
	"crypto/ecdsa"
	"math/big"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcore "github.com/ethereum/go-ethereum/core/types"
	"github.com/okex/exchain-go-sdk/module/evm/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	rpctypes "github.com/okex/exchain/app/rpc/types"
	"github.com/okex/exchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
)

// Evm shows the expected behavior for inner farm client
type Evm interface {
	gosdktypes.Module
	EvmTx
	EvmQuery
	EvmUtils
	web3Getter
}

// EvmTx shows the expected tx behavior for inner evm client
type EvmTx interface {
	SendTx(fromInfo keys.Info, passWd, toAddrStr, amountStr, payloadStr, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	CreateContract(fromInfo keys.Info, passWd, amountStr, payloadStr, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, string, error)
	SendTxEthereum2(privHex, toAddrStr, amountStr, payloadStr string, gasLimit, seqNum uint64) (sdk.TxResponse, error)
	SendTxEthereum(priv *ecdsa.PrivateKey, nonce uint64, to ethcmn.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (resp sdk.TxResponse, err error)
	CreateContractEthereum(privHex, amountStr, payloadStr string, gasLimit, seqNum uint64) (sdk.TxResponse, error)
}

// EvmQuery shows the expected query behavior for inner evm client
type EvmQuery interface {
	QueryCode(contractAddrStr string) (types.QueryResCode, error)
	QueryStorageAt(contractAddrStr, keyHexStr string) (types.QueryResStorage, error)
}

type EvmUtils interface {
	GetTxHash(signedTx *ethcore.Transaction) (ethcmn.Hash, error)
}

type web3Getter interface {
	Web3Proxy() Web3Proxy
}

// Web3Proxy shows the expected behavior as Web3 without rest server routing
type Web3Proxy interface {
	Web3ProxyQuery
}

// Web3Query shows the expected behavior as web3 query request
type Web3ProxyQuery interface {
	BlockNumberProxy() (hexutil.Uint64, error)
	EstimateGasProxy(args rpctypes.CallArgs) (hexutil.Uint64, error)
}
