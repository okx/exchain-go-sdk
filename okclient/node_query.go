package okclient

import (
	abci "github.com/ok-chain/gosdk/types/abci"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

const (
	abciTokenPairPath = "/custom/token/tokenpair"
)

func (okCli *OKClient) QueryABCIInfo() (abci.ResponseInfo, error) {
	resp, err := okCli.cli.ABCIInfo()
	if err != nil {
		return abci.ResponseInfo{}, err
	}
	return resp.Response, nil
}

func (okCli *OKClient) QueryConsenusState() ([]byte, error) {
	resp, err := okCli.cli.ConsensusState()
	if err != nil {
		return nil, err
	}
	return resp.RoundState, nil
}

func (okCli *OKClient) QueryDumpConsenusState() ([]byte, error) {
	resp, err := okCli.cli.DumpConsensusState()
	if err != nil {
		return nil, err
	}
	return resp.RoundState, nil
}

func (okCli *OKClient) QueryNetInfo() (*ctypes.ResultNetInfo, error) {
	resp, err := okCli.cli.NetInfo()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryGenesisFile() (*types.GenesisDoc, error) {
	resp, err := okCli.cli.Genesis()
	if err != nil {
		return nil, err
	}
	return resp.Genesis, nil
}

func (okCli *OKClient) QueryHealthInfo() (*ctypes.ResultHealth, error) {
	resp, err := okCli.cli.Health()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryUnconfirmedTxsNum(limit int) (*ctypes.ResultUnconfirmedTxs, error) {
	resp, err := okCli.cli.UnconfirmedTxs(limit)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryStateInfo() (*ctypes.ResultStatus, error) {
	resp, err := okCli.cli.Status()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryABCITokenpair() (abci.ResponseQuery, error) {
	resp, err := okCli.cli.ABCIQuery(abciTokenPairPath, nil)
	if err != nil {
		return abci.ResponseQuery{}, err
	}
	return resp.Response, nil
}

func (okCli *OKClient) QueryBlock(height *int64) (*ctypes.ResultBlock, error) {
	resp, err := okCli.cli.Block(height)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryBlockResults(height *int64) (*ctypes.ResultBlockResults, error) {
	resp, err := okCli.cli.BlockResults(height)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryBlockchainInfo(minHeight, maxHeight int64) (*ctypes.ResultBlockchainInfo, error) {
	resp, err := okCli.cli.BlockchainInfo(minHeight, maxHeight)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryCommit(height *int64) (*ctypes.ResultCommit, error) {
	resp, err := okCli.cli.Commit(height)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryTx(txHash []byte, prove bool) (*ctypes.ResultTx, error) {
	resp, err := okCli.cli.Tx(txHash, prove)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (okCli *OKClient) QueryTxOnHeight(queryStr string, prove bool, page, perPage int) (*ctypes.ResultTxSearch, error) {
	resp, err := okCli.cli.TxSearch(queryStr, prove, page, perPage)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
