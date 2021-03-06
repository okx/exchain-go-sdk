package evm

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/okex/okexchain-go-sdk/exposed"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// Web3Proxy returns the client with exposed.Web3Proxy's behaviour
func (ec evmClient) Web3Proxy() exposed.Web3Proxy {
	return gosdktypes.Module(ec).(exposed.Web3Proxy)
}

// BlockNumber returns the current block number as method "eth_blockNumber"
func (ec evmClient) BlockNumber() (hexutil.Uint64, error) {
	resBlockchainInfo, err := ec.BlockchainInfo(0, 0)
	if err != nil {
		return hexutil.Uint64(0), err
	}

	blockNumber := resBlockchainInfo.LastHeight
	if blockNumber > 0 {
		// decrease blockNumber to make sure every block has been executed in local
		blockNumber--
	}

	return hexutil.Uint64(blockNumber), nil
}
