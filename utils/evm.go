package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
)

// ToCosmosAddress converts string address of cosmos and ethereum style to cosmos address
func ToCosmosAddress(addrStr string) (accAddr sdk.AccAddress, err error) {
	if strings.HasPrefix(addrStr, sdk.GetConfig().GetBech32AccountAddrPrefix()) {
		accAddr, err = sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return accAddr, fmt.Errorf("failed. invalid bech32 formatted address: %s", err)
		}
		return
	}

	addrStr = strings.TrimPrefix(addrStr, "0x")
	return sdk.AccAddressFromHex(addrStr)
}

// ToHexAddress converts string address of cosmos and ethereum style to ethereum address
func ToHexAddress(addrStr string) (ethAddr ethcmn.Address, err error) {
	if strings.HasPrefix(addrStr, sdk.GetConfig().GetBech32AccountAddrPrefix()) {
		accAddr, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return ethAddr, fmt.Errorf("failed. invalid bech32 formatted address: %s", err)
		}

		return ethcmn.BytesToAddress(accAddr.Bytes()), err
	}

	// hex string
	if !ethcmn.IsHexAddress(addrStr) {
		return ethAddr, fmt.Errorf("failed. invalid hex address: %s", addrStr)
	}

	return ethcmn.HexToAddress(addrStr), err
}

// FormatKeyToHash converts the key string to hash
func FormatKeyToHash(keyStr string) string {
	if !strings.HasPrefix(keyStr, "0x") {
		keyStr = fmt.Sprintf("0x%s", keyStr)
	}

	ethkey := ethcmn.HexToHash(keyStr)
	return ethkey.Hex()
}

// Uint256 gets the available arg for payload Build
func Uint256(i *big.Int) *big.Int {
	return i
}

// EthAddress gets the available arg for payload Build
func EthAddress(ethAddrStr string) ethcmn.Address {
	return ethcmn.HexToAddress(ethAddrStr)
}

// EthAddresses gets the available arg for payload Build
func EthAddresses(ethAddrsStr []string) []ethcmn.Address {
	var ethAddrs []ethcmn.Address
	for _, ethAddrStr := range ethAddrsStr {
		ethAddrs = append(ethAddrs, ethcmn.HexToAddress(ethAddrStr))
	}

	return ethAddrs
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
