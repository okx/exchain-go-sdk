package okclient

import (
	"github.com/ok-chain/ok-gosdk/common/transactParams"
	"github.com/ok-chain/ok-gosdk/types"
)

func (okCli *OKClient) Send(tp transactParams.TransferParams) (resp types.TxResponse, err error) {
	if ok, err := tp.IsValid(); !ok {
		return resp, err
	}



	return types.TxResponse{}, nil
}
