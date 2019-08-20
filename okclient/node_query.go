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
