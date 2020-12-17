package evm

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	evmtypes "github.com/okex/okexchain/x/evm/types"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const (
	addr              = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	name              = "alice"
	passWd            = "12345678"
	accPubkey         = "okexchainpub17weu6qepq0ph2t3u697qar7rmdtdtqp4744jcprjd2h356zr0yh5vmw38a3my4vqjx5"
	mnemonic          = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	memo              = "my memo"
	recAddr           = "okexchain193xnjknz3e52mqv2nyufnzjugu3mh65rpxdasn"
	recAddrEth        = "0x2c4d395A628e68aD818a9938998a5C4723BBEa83"
	defaultPayloadStr = "0x0123456789abcdef"
	badPayloadStr     = "0x0123456789abcdefg"

	contractAddr = "0x9aD84c8630E0282F78e5479B46E64E17779e3Cfb"
	codeContent  = "default code content"
)

func TestEvmClient_QueryCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
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
