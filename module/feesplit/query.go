package feesplit

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/okx/okbchain/libs/cosmos-sdk/types/query"
	"github.com/okx/okbchain/x/feesplit/types"
)

func (c feesplitClient) QueryFeesplits(pageReq *query.PageRequest) (*types.QueryFeeSplitsResponse, error) {
	req := &types.QueryFeeSplitsRequest{Pagination: pageReq}
	queryData, err := c.GetCodec().MarshalJSON(req)
	if err != nil {
		return nil, err
	}

	// query store
	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryFeeSplits)
	bz, _, err := c.QueryWithData(route, queryData)

	if err != nil {
		return nil, err
	}

	// unmarshal
	var resp types.QueryFeeSplitsResponse
	c.GetCodec().MustUnmarshalJSON(bz, &resp)

	return &resp, nil
}

func (c feesplitClient) QueryFeeSplit(contractAddress string) (*types.QueryFeeSplitResponse, error) {
	if !common.IsHexAddress(contractAddress) {
		return nil, errors.New("invalid contractAddress")
	}

	req := &types.QueryFeeSplitRequest{ContractAddress: contractAddress}
	data, err := c.GetCodec().MarshalJSON(req)
	if err != nil {
		return nil, err
	}

	// Query store
	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryFeeSplit)
	bz, _, err := c.QueryWithData(route, data)
	if err != nil {
		return nil, err
	}

	var resp types.QueryFeeSplitResponse
	c.GetCodec().MustUnmarshalJSON(bz, &resp)

	return &resp, nil
}

func (c feesplitClient) QueryParams() (*types.QueryParamsResponse, error) {
	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryParameters)
	bz, _, err := c.QueryWithData(route, nil)
	if err != nil {
		return nil, err
	}

	var params types.QueryParamsResponse
	c.GetCodec().MustUnmarshalJSON(bz, &params)

	return &params, nil
}

func (c feesplitClient) QueryDeployerFeeSplits(deployerAddress string, pageReq *query.PageRequest) (*types.QueryDeployerFeeSplitsResponse, error) {
	req := &types.QueryDeployerFeeSplitsRequest{
		DeployerAddress: deployerAddress,
		Pagination:      pageReq,
	}
	data, err := c.GetCodec().MarshalJSON(req)
	if err != nil {
		return nil, err
	}

	// Query store
	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryDeployerFeeSplits)
	bz, _, err := c.QueryWithData(route, data)
	if err != nil {
		return nil, err
	}

	var resp types.QueryDeployerFeeSplitsResponse
	c.GetCodec().MustUnmarshalJSON(bz, &resp)

	return &resp, nil
}

func (c feesplitClient) QueryWithdrawerFeeSplits(withdrawerAddress string, pageReq *query.PageRequest) (*types.QueryWithdrawerFeeSplitsResponse, error) {
	req := &types.QueryWithdrawerFeeSplitsRequest{
		WithdrawerAddress: withdrawerAddress,
		Pagination:        pageReq,
	}
	data, err := c.GetCodec().MarshalJSON(req)
	if err != nil {
		return nil, err
	}

	// Query store
	route := fmt.Sprintf("custom/%s/%s", types.ModuleName, types.QueryWithdrawerFeeSplits)
	bz, _, err := c.QueryWithData(route, data)
	if err != nil {
		return nil, err
	}

	var resp types.QueryWithdrawerFeeSplitsResponse
	c.GetCodec().MustUnmarshalJSON(bz, &resp)

	return &resp, nil
}
