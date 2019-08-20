package okclient

import (
	abci "github.com/ok-chain/ok-gosdk/types/abci"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
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


