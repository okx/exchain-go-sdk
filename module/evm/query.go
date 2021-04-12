package evm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/okex/exchain-go-sdk/module/evm/types"
	"github.com/okex/exchain-go-sdk/utils"
	evmtypes "github.com/okex/exchain/x/evm/types"
)

// QueryCode gets the contract code from ExChain
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

// QueryStorageAt gets the encoded bytes with a specific key from the storage
func (ec evmClient) QueryStorageAt(contractAddrStr, keyStr string) (resStorage types.QueryResStorage, err error) {
	if !strings.HasPrefix(contractAddrStr, "0x") {
		contractAddrStr = fmt.Sprintf("0x%s", contractAddrStr)
	}

	key := utils.FormatKeyToHash(keyStr)
	path := fmt.Sprintf("custom/%s/storage/%s/%s", evmtypes.RouterKey, common.HexToAddress(contractAddrStr).Hex(), key)
	res, _, err := ec.Query(path, nil)
	if err != nil {
		return resStorage, utils.ErrClientQuery(err.Error())
	}

	if err = ec.GetCodec().UnmarshalJSON(res, &resStorage); err != nil {
		return resStorage, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
