package evm

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	evmtypes "github.com/okex/exchain/x/evm/types"
)

// SendTxEthereum sends an ethereum tx
func (ec evmClient) SendTxEthereum(priv *ecdsa.PrivateKey, nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (
	resp sdk.TxResponse, err error) {

	ethMsg := evmtypes.NewMsgEthereumTx(
		nonce,
		&to,
		amount,
		gasLimit,
		gasPrice,
		data,
	)

	config := ec.GetConfig()
	if err = ethMsg.Sign(config.ChainIDBigInt, priv); err != nil {
		return
	}

	bytes, err := ec.GetCodec().MarshalBinaryLengthPrefixed(ethMsg)
	if err != nil {
		return resp, fmt.Errorf("failed. encoded MsgEthereumTx error: %s", err)
	}

	return ec.Broadcast(bytes, ec.GetConfig().BroadcastMode)
}

// CreateContractEthereum generates an ethereum tx to deploy a smart contract
func (ec evmClient) CreateContractEthereum(priv *ecdsa.PrivateKey, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (
	resp sdk.TxResponse, err error) {

	ethMsg := evmtypes.NewMsgEthereumTxContract(
		nonce,
		amount,
		gasLimit,
		gasPrice,
		data,
	)

	config := ec.GetConfig()
	if err = ethMsg.Sign(config.ChainIDBigInt, priv); err != nil {
		return
	}

	bytes, err := ec.GetCodec().MarshalBinaryLengthPrefixed(ethMsg)
	if err != nil {
		return resp, fmt.Errorf("failed. encoded MsgEthereumTx error: %s", err)
	}

	return ec.Broadcast(bytes, ec.GetConfig().BroadcastMode)
}
