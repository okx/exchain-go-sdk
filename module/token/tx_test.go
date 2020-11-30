package token

import (
	"bytes"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/auth"
	"github.com/okex/okexchain-go-sdk/module/token/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTokenClient_Send(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
		accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Token().Send(fromInfo, passWd, recAddr, "10.24okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	res, err = mockCli.Token().Send(fromInfo, passWd, recAddr[1:], "10.24okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Token().Send(fromInfo, "", recAddr, "10.24okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Token().Send(fromInfo, passWd, recAddr, "10.24", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	badBech32Addr := fmt.Sprintf("%s1", recAddr[:len(recAddr)-1])
	_, err = mockCli.Token().Send(fromInfo, passWd, badBech32Addr, "10.24okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)
}

func TestTokenClient_MultiSend(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	transfersStr := "okexchain1qeh2fz0a4t78ylesd4cyd2mwt5wcfnfj98ev0u 1024okt,2048btc\nokexchain1zfy46z8tqlu0cjmuj5t9dple276der4rtq5tu3 20.48okt"
	transfers, err := utils.ParseTransfersStr(transfersStr)
	require.NoError(t, err)
	res, err := mockCli.Token().MultiSend(fromInfo, passWd, transfers, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	var emptyTransfer types.TransferUnit
	badTransfers := []types.TransferUnit{emptyTransfer}

	_, err = mockCli.Token().MultiSend(fromInfo, passWd, badTransfers, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Token().MultiSend(fromInfo, "", transfers, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestTokenClient_Issue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Token().Issue(fromInfo, passWd, "default original symbol", "default whole name",
		"default total supply", "default token description", memo, true, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Token().Issue(fromInfo, passWd, "", "default whole name",
		"default total supply", "default token description", memo, true, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Token().Issue(fromInfo, passWd, "default original symbol", "default whole name",
		"default total supply", "", memo, true, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	// build a invalid long token description
	var buffer bytes.Buffer
	for i := 0; i < 257; i++ {
		_, _ = buffer.WriteString("a")
	}
	longDesc := buffer.String()

	_, err = mockCli.Token().Issue(fromInfo, passWd, "default original symbol", "default whole name",
		"default total supply", longDesc, memo, true, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Token().Issue(fromInfo, passWd, "default original symbol", "",
		"default total supply", "default token description", memo, true, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Token().Issue(fromInfo, "", "default original symbol", "default whole name",
		"default total supply", "default token description", memo, true, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)
}

func TestTokenClient_Mint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Token().Mint(fromInfo, passWd, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	res, err = mockCli.Token().Mint(fromInfo, passWd, "1024.1024", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	res, err = mockCli.Token().Mint(fromInfo, "", "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	res, err = mockCli.Token().Mint(fromInfo, passWd, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

}

func TestTokenClient_Burn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Token().Burn(fromInfo, passWd, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	res, err = mockCli.Token().Burn(fromInfo, passWd, "1024.1024", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	res, err = mockCli.Token().Burn(fromInfo, "", "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	res, err = mockCli.Token().Burn(fromInfo, passWd, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

}

func TestTokenClient_Edit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewTokenClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Token().Edit(fromInfo, passWd, "okt", "new description", "new whole name",
		memo, true, true, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	res, err = mockCli.Token().Edit(fromInfo, passWd, "", "new description", "new whole name",
		memo, true, true, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	res, err = mockCli.Token().Edit(fromInfo, passWd, "okt", "new description", ".new whole name",
		memo, true, true, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// build a invalid long token description
	var buffer bytes.Buffer
	for i := 0; i < 257; i++ {
		_, _ = buffer.WriteString("a")
	}
	longDesc := buffer.String()
	res, err = mockCli.Token().Edit(fromInfo, passWd, "okt", longDesc, "new whole name",
		memo, true, true, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	res, err = mockCli.Token().Edit(fromInfo, "", "okt", "new description", "new whole name",
		memo, true, true, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	res, err = mockCli.Token().Edit(fromInfo, passWd, "okt", "new description", "new whole name",
		memo, true, true, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

}
