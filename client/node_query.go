package client

import (
	sdktypes "github.com/ok-chain/gosdk/types"
	abci "github.com/ok-chain/gosdk/types/abci"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

const (
	abciTokenPairPath = "/custom/token/tokenpair"
)

func (cli *OKChainClient) QueryABCIInfo() (abci.ResponseInfo, error) {
	resp, err := cli.cli.ABCIInfo()
	if err != nil {
		return abci.ResponseInfo{}, err
	}
	return resp.Response, nil
}

func (cli *OKChainClient) QueryConsenusState() ([]byte, error) {
	resp, err := cli.cli.ConsensusState()
	if err != nil {
		return nil, err
	}
	return resp.RoundState, nil
}

func (cli *OKChainClient) QueryDumpConsenusState() ([]byte, error) {
	resp, err := cli.cli.DumpConsensusState()
	if err != nil {
		return nil, err
	}
	return resp.RoundState, nil
}

func (cli *OKChainClient) QueryNetInfo() (*ctypes.ResultNetInfo, error) {
	resp, err := cli.cli.NetInfo()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryGenesisFile() (*types.GenesisDoc, error) {
	resp, err := cli.cli.Genesis()
	if err != nil {
		return nil, err
	}
	return resp.Genesis, nil
}

func (cli *OKChainClient) QueryHealthInfo() (*ctypes.ResultHealth, error) {
	resp, err := cli.cli.Health()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryUnconfirmedTxsNum(limit int) (*ctypes.ResultUnconfirmedTxs, error) {
	resp, err := cli.cli.UnconfirmedTxs(limit)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryStateInfo() (*ctypes.ResultStatus, error) {
	resp, err := cli.cli.Status()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryABCITokenpair() (abci.ResponseQuery, error) {
	resp, err := cli.cli.ABCIQuery(abciTokenPairPath, nil)
	if err != nil {
		return abci.ResponseQuery{}, err
	}
	return resp.Response, nil
}

func (cli *OKChainClient) QueryBlock(height *int64) (*ctypes.ResultBlock, error) {
	resp, err := cli.cli.Block(height)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryBlockResults(height *int64) (*ctypes.ResultBlockResults, error) {
	resp, err := cli.cli.BlockResults(height)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryBlockchainInfo(minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
	resp, err := cli.cli.BlockchainInfo(minHeight, maxHeight)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryCommit(height *int64) (*ctypes.ResultCommit, error) {
	resp, err := cli.cli.Commit(height)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryTx(txHash []byte, prove bool) (*ctypes.ResultTx, error) {
	resp, err := cli.cli.Tx(txHash, prove)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryTxOnHeight(queryStr string, prove bool, page, perPage int) (*ctypes.ResultTxSearch, error) {
	resp, err := cli.cli.TxSearch(queryStr, prove, page, perPage)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *OKChainClient) QueryCurrentValidators() (sdktypes.ResultValidatorsOutput, error) {
	resp, err := cli.cli.Validators(nil)
	if err != nil {
		return sdktypes.ResultValidatorsOutput{}, err
	}

	outputValidatorsRes, err := sdktypes.NewResultValidatorsOutput(resp)
	if err != nil {
		return sdktypes.ResultValidatorsOutput{}, err
	}

	return outputValidatorsRes, nil
}
