package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/crypto/encoding/codec"
	"github.com/okex/okchain-go-sdk/types"
)

const (
	transactionsInfoPath  = "custom/backend/txs"

)




func (cli *OKChainClient) GetTransactionsInfo(addr string, type_, start, end, page, perPage int) ([]types.Transaction, error) {
	perPageTmp, err := checkParamsGetTransactionsInfo(addr, type_, start, end, page, perPage)
	if err != nil {
		return nil, err
	}

	params := params.NewQueryTxListParams(addr, int64(type_), int64(start), int64(end), page, perPageTmp)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryTxListParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(transactionsInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var transactionsInfo []types.Transaction
	if err = codec.UnmarshalListResponse(res, &transactionsInfo); err != nil {
		return nil, fmt.Errorf("transactions Info list unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return transactionsInfo, nil
}