package evm

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/okex/okexchain-go-sdk/module/evm/types"
	"github.com/okex/okexchain-go-sdk/utils"
	evmtypes "github.com/okex/okexchain/x/evm/types"
	"strings"
)

// QueryCode gets the contract code from OKExChain
func (ec evmClient) QueryCode(contractAddrStr string) (resCode types.QueryResCode, err error) {
	if !strings.HasPrefix(contractAddrStr, "0x") {
		contractAddrStr = fmt.Sprintf("0x%s", contractAddrStr)
	}

	if !common.IsHexAddress(contractAddrStr) {
		return resCode, errors.New("failed. invalid contract address")
	}

	path := fmt.Sprintf("custom/%s/code/%s", evmtypes.RouterKey, common.HexToAddress(contractAddrStr).Hex())
	res, _, err := ec.Query(path, nil)
	if err != nil {
		return resCode, utils.ErrClientQuery(err.Error())
	}

	if err = ec.GetCodec().UnmarshalJSON(res, &resCode); err != nil {
		return resCode, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
