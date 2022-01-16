package evm

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core/types"
	authcli "github.com/okex/exchain/libs/cosmos-sdk/x/auth/client/utils"
	"github.com/okex/exchain/libs/tendermint/crypto/etherhash"
	evmtypes "github.com/okex/exchain/x/evm/types"
)

// GetTxHash calculates the tx hash
func (ec evmClient) GetTxHash(signedTx *ethcore.Transaction) (txHash ethcmn.Hash, err error) {
	v, r, s := signedTx.RawSignatureValues()
	tx := evmtypes.MsgEthereumTx{
		Data: evmtypes.TxData{
			AccountNonce: signedTx.Nonce(),
			Price:        signedTx.GasPrice(),
			GasLimit:     signedTx.Gas(),
			Recipient:    signedTx.To(),
			Amount:       signedTx.Value(),
			Payload:      signedTx.Data(),
			V:            v,
			R:            r,
			S:            s,
		},
	}

	txBytes, err := authcli.GetTxEncoder(ec.GetCodec(), authcli.WithEthereumTx())(&tx)
	if err != nil {
		return
	}

	txHash = ethcmn.BytesToHash(etherhash.Sum(txBytes))
	return
}
