package okclient

import (
	abci "github.com/ok-chain/ok-gosdk/types/abci"
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