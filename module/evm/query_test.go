package evm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/okex/exchain-go-sdk/mocks"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain-go-sdk/utils"
	evmtypes "github.com/okx/okbchain/x/evm/types"
	"github.com/stretchr/testify/require"
)

const (
	addr              = "ex1qj5c07sm6jetjz8f509qtrxgh4psxkv3ddyq7u"
	name              = "alice"
	passWd            = "12345678"
	accPubkey         = "expub17weu6qepqtfc6zq8dukwc3lhlhx7th2csfjw0g3cqnqvanh7z9c2nhkr8mn5z9uq4q6"
	mnemonic          = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	memo              = "my memo"
	recAddr           = "ex1alrwch5sd3wm3np4njz7l754xtnng6cf4z9s5v"
	recAddrEth        = "0x2c4d395A628e68aD818a9938998a5C4723BBEa83"
	defaultPayloadStr = "0x0123456789abcdef"
	badPayloadStr     = "0x0123456789abcdefg"

	contractAddr = "0x9aD84c8630E0282F78e5479B46E64E17779e3Cfb"
	codeContent  = "default code content"
	storageValue = "default storage value"
	keyStr       = "defaultKey"
)

func TestEvmClient_QueryCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewEvmClient(mockCli.MockBaseClient))

	expectedRet := mockCli.BuildQueryResCode(codeContent)
	expectedCdc := mockCli.GetCodec()
	expectedPath := fmt.Sprintf("custom/%s/code/%s", evmtypes.RouterKey, common.HexToAddress(contractAddr).Hex())

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(3)
	mockCli.EXPECT().Query(expectedPath, nil).Return(expectedRet, int64(1024), nil).Times(2)

	resCode, err := mockCli.Evm().QueryCode(contractAddr)
	require.NoError(t, err)
	require.True(t, strings.EqualFold(codeContent, string(resCode.Code)))

	resCode, err = mockCli.Evm().QueryCode(contractAddr[2:])
	require.NoError(t, err)
	require.True(t, strings.EqualFold(codeContent, string(resCode.Code)))

	_, err = mockCli.Evm().QueryCode(contractAddr[1:])
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, nil).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Evm().QueryCode(contractAddr)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, nil).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Evm().QueryCode(contractAddr)
	require.Error(t, err)
}

func TestEvmClient_QueryStorageAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewEvmClient(mockCli.MockBaseClient))

	expectedRet := mockCli.BuildQueryResStorage(storageValue)
	expectedCdc := mockCli.GetCodec()
	expectedPath := fmt.Sprintf("custom/%s/storage/%s/%s", evmtypes.RouterKey, common.HexToAddress(contractAddr).Hex(),
		utils.FormatKeyToHash(keyStr))

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(3)
	mockCli.EXPECT().Query(expectedPath, nil).Return(expectedRet, int64(1024), nil).Times(2)

	resStorage, err := mockCli.Evm().QueryStorageAt(contractAddr, keyStr)
	require.NoError(t, err)
	require.True(t, strings.EqualFold(storageValue, string(resStorage.Value)))

	resStorage, err = mockCli.Evm().QueryStorageAt(contractAddr[2:], keyStr)
	require.NoError(t, err)
	require.True(t, strings.EqualFold(storageValue, string(resStorage.Value)))

	mockCli.EXPECT().Query(expectedPath, nil).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Evm().QueryStorageAt(contractAddr, keyStr)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, nil).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Evm().QueryStorageAt(contractAddr, keyStr)
	require.Error(t, err)
}
