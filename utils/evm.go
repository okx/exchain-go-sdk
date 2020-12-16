package utils

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

// ToCosmosAddress converts string address of cosmos and ethereum style to cosmos address
func ToCosmosAddress(addrStr string) (toAddr sdk.AccAddress, err error) {
	if strings.HasPrefix(addrStr, sdk.GetConfig().GetBech32AccountAddrPrefix()) {
		toAddr, err = sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return toAddr, fmt.Errorf("failed. invalid bech32 formatted address: %s", err)
		}
		return
	}

	// strip 0x prefix if exists
	addrStr = strings.TrimPrefix(addrStr, "0x")
	return sdk.AccAddressFromHex(addrStr)
}

// GetEthAddressStrFromCosmosAddr gets the string of eth address from a cosmos acc addr
func GetEthAddressStrFromCosmosAddr(accAddr sdk.AccAddress) string {
	return common.BytesToAddress(accAddr.Bytes()).Hex()
}

// Uint256 gets the available arg for payload Build
func Uint256(n int) *big.Int {
	return big.NewInt(int64(n))
}

// EthAddress gets the available arg for payload Build
func EthAddress(ethAddrStr string) common.Address {
	return common.HexToAddress(ethAddrStr)
}

// PayloadBuilder - structure of a useful tool to build payload
type PayloadBuilder struct {
	innerABI abi.ABI
	binData  []byte
}

// NewPayloadBuilder creates a new instance of PayloadBuilder
func NewPayloadBuilder(binStr, abiJSON string) (payloadBuilder PayloadBuilder, err error) {
	innerABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return
	}

	if !strings.HasPrefix(binStr, "0x") {
		binStr = fmt.Sprintf("0x%s", binStr)
	}

	binData, err := hexutil.Decode(binStr)
	if err != nil {
		return
	}

	return PayloadBuilder{
		innerABI: innerABI,
		binData:  binData,
	}, err
}

// Build gets the payload data for tx in evm module
func (pb *PayloadBuilder) Build(methodName string, args ...interface{}) (payload []byte, err error) {
	paramsData, err := pb.innerABI.Pack(methodName, args...)
	if err != nil {
		return
	}

	if len(methodName) == 0 {
		lenBinData := len(pb.binData)
		payload = make([]byte, lenBinData+len(paramsData))
		copy(payload[:lenBinData], pb.binData)
		copy(payload[lenBinData:], paramsData)
	} else {
		payload = paramsData
	}

	return
}
