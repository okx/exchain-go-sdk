package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/crypto/encoding/codec"
	"github.com/okex/okchain-go-sdk/types"
)

const (
	dealsInfoPath         = "custom/backend/deals"
	transactionsInfoPath  = "custom/backend/txs"

)




func (cli *OKChainClient) GetDealsInfo(addr, product, side string, start, end, page, perPage int) ([]types.Deal, error) {
	perPageTmp, err := checkParamsGetDealsInfo(addr, product, side, start, end, page, perPage)
	if err != nil {
		return nil, err
	}

	params := params.NewQueryDealsParams(addr, product, int64(start), int64(end), page, perPageTmp, side)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryDealsParams failed in json marshal : %s", err.Error())
	}

	res, err := cli.query(dealsInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var dealsInfo []types.Deal

	if err = codec.UnmarshalListResponse(res, &dealsInfo); err != nil {
		return nil, fmt.Errorf("deals Info list unmarshaled failed from BaseResponse : %s", err.Error())
	}

	return dealsInfo, nil
}

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