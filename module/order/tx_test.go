package order

import (
	"testing"
)

func TestOrderClient_NewOrders(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	//require.NoError(t, err)
	//mockCli := mocks.NewMockClient(t, ctrl, config)
	//mockCli.RegisterModule(NewOrderClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))
	//
	//fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	//require.NoError(t, err)
	//
	//accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "1024okt", 1, 2)
	//expectedCdc := mockCli.GetCodec()
	//mockCli.EXPECT().GetCodec().Return(expectedCdc)
	//mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)
	//
	//accInfo, err := mockCli.Auth().QueryAccount(addr)
	//require.NoError(t, err)
	//
	//mockCli.EXPECT().BuildAndBroadcast(
	//	fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
	//	accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	//res, err := mockCli.Dex().Withdraw(fromInfo, passWd, product, "1.024okt", memo,
	//	accInfo.GetAccountNumber(), accInfo.GetSequence())
	//require.NoError(t, err)
	//require.Equal(t, uint32(0), res.Code)
}
