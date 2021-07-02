package evm

import (
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	evmtypes "github.com/okex/exchain/x/evm/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// GetTxHash calculates the tx hash
func (ec evmClient) GetTxHash(signedTx *ethcore.Transaction) (txHash ethcmn.Hash, err error) {
	var tx evmtypes.MsgEthereumTx
	if err = rlp.DecodeBytes(ethcore.Transactions{signedTx}.GetRlp(0), &tx); err != nil {
		return
	}

	txBytes, err := authcli.GetTxEncoder(ec.GetCodec())(tx)
	if err != nil {
		return
	}

	txHash = ethcmn.BytesToHash(tmhash.Sum(txBytes))
	return
}
