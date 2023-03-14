package governance

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/okex/exchain-go-sdk/mocks"
	"github.com/okex/exchain-go-sdk/module/auth"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain-go-sdk/utils"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	textProposalFilePath               = "./text_proposal.json"
	paramsChangeProposalFilePath       = "./param_change_proposal.json"
	delistProposalFilePath             = "./delist_proposal.json"
	communityPoolSpendProposalFilePath = "./community_pool_spend_proposal.json"
	manageWhiteListProposalFilePath    = "./manage_white_list_proposal.json"
	badProposalFilePath                = "./bad_proposal.json"

	textProposalJSON               = `{"title":"Text Proposal","description":"text proposal description","proposal_type":"Text","deposit":"100okt"}`
	paramsChangeProposalJSON       = `{"title":"Param Change Proposal","description":"param change proposal description","changes":[{"subspace":"staking","key":"MaxValidators","value":105}],"deposit":[{"denom":"okt","amount":"100"}],"height":"1024"}`
	delistProposalJSON             = `{"title":"Delist Proposal","description":"delist proposal description","base_asset":"btc-000","quote_asset":"okt","deposit":[{"denom":"okt","amount":"100"}]}`
	communityPoolSpendProposalJSON = `{"title":"Community Pool Spend Proposal","description":"community pool spend description","recipient":"ex1qj5c07sm6jetjz8f509qtrxgh4psxkv3ddyq7u","amount":[{"denom":"okt","amount":"10.24"}],"deposit":[{"denom":"okt","amount":"100"}]}`
	manageWhiteListProposalJSON    = `{"title":"Manage White List Proposal","description":"manage white list description","pool_name":"pool1","is_added":true,"deposit":[{"denom":"tokt","amount":"100"}]}`
)

func TestGovClient_SubmitTextProposal(t *testing.T) {
	// build the text proposal JSON file
	err := ioutil.WriteFile(textProposalFilePath, []byte(textProposalJSON), 0644)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

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

	res, err := mockCli.Governance().SubmitTextProposal(fromInfo, passWd, textProposalFilePath, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Governance().SubmitTextProposal(fromInfo, passWd, textProposalFilePath[1:], memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	// bad text proposal JSON file
	err = ioutil.WriteFile(badProposalFilePath, []byte(textProposalJSON[1:]), 0644)
	require.NoError(t, err)
	_, err = mockCli.Governance().SubmitTextProposal(fromInfo, passWd, badProposalFilePath, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	badDepositTextProposalJSON := `{"title": "Text Proposal", "description": "text proposal description","proposal_type": "Text","deposit": "100"}`
	err = ioutil.WriteFile(badProposalFilePath, []byte(badDepositTextProposalJSON), 0644)
	require.NoError(t, err)
	_, err = mockCli.Governance().SubmitTextProposal(fromInfo, passWd, badProposalFilePath, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Governance().SubmitTextProposal(fromInfo, passWd, textProposalFilePath, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().SubmitTextProposal(fromInfo, "", textProposalFilePath, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	// remove the temporary files
	err = os.Remove(textProposalFilePath)
	require.NoError(t, err)
	err = os.Remove(badProposalFilePath)
	require.NoError(t, err)
}

func TestGovClient_SubmitParamChangeProposal(t *testing.T) {
	// build the param change proposal JSON file
	err := ioutil.WriteFile(paramsChangeProposalFilePath, []byte(paramsChangeProposalJSON), 0644)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(6)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Governance().SubmitParamsChangeProposal(fromInfo, passWd, paramsChangeProposalFilePath, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Governance().SubmitParamsChangeProposal(fromInfo, passWd, paramsChangeProposalFilePath[1:], memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// bad param change proposal JSON file
	err = ioutil.WriteFile(badProposalFilePath, []byte(paramsChangeProposalJSON[1:]), 0644)
	require.NoError(t, err)
	_, err = mockCli.Governance().SubmitParamsChangeProposal(fromInfo, passWd, badProposalFilePath, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Governance().SubmitParamsChangeProposal(fromInfo, passWd, paramsChangeProposalFilePath, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().SubmitParamsChangeProposal(fromInfo, "", paramsChangeProposalFilePath, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// remove the temporary files
	err = os.Remove(paramsChangeProposalFilePath)
	require.NoError(t, err)
	err = os.Remove(badProposalFilePath)
	require.NoError(t, err)
}

func TestGovClient_SubmitDelistProposal(t *testing.T) {
	// build the delist proposal JSON file
	err := ioutil.WriteFile(delistProposalFilePath, []byte(delistProposalJSON), 0644)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(6)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Governance().SubmitDelistProposal(fromInfo, passWd, delistProposalFilePath, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Governance().SubmitDelistProposal(fromInfo, passWd, delistProposalFilePath[1:], memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// bad delist proposal JSON file
	err = ioutil.WriteFile(badProposalFilePath, []byte(delistProposalJSON[1:]), 0644)
	require.NoError(t, err)
	_, err = mockCli.Governance().SubmitDelistProposal(fromInfo, passWd, badProposalFilePath, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Governance().SubmitDelistProposal(fromInfo, passWd, delistProposalFilePath, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().SubmitDelistProposal(fromInfo, "", delistProposalFilePath, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// remove the temporary files
	err = os.Remove(delistProposalFilePath)
	require.NoError(t, err)
	err = os.Remove(badProposalFilePath)
	require.NoError(t, err)
}

func TestGovClient_SubmitCommunityPoolSpendProposal(t *testing.T) {
	// build the community pool spend proposal JSON file
	err := ioutil.WriteFile(communityPoolSpendProposalFilePath, []byte(communityPoolSpendProposalJSON), 0644)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(6)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Governance().SubmitCommunityPoolSpendProposal(fromInfo, passWd, communityPoolSpendProposalFilePath,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Governance().SubmitCommunityPoolSpendProposal(fromInfo, passWd, communityPoolSpendProposalFilePath[1:],
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// bad community pool spend proposal JSON file
	err = ioutil.WriteFile(badProposalFilePath, []byte(communityPoolSpendProposalJSON[1:]), 0644)
	require.NoError(t, err)
	_, err = mockCli.Governance().SubmitCommunityPoolSpendProposal(fromInfo, passWd, badProposalFilePath, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Governance().SubmitCommunityPoolSpendProposal(fromInfo, passWd, communityPoolSpendProposalFilePath,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().SubmitCommunityPoolSpendProposal(fromInfo, "", communityPoolSpendProposalFilePath,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// remove the temporary files
	err = os.Remove(communityPoolSpendProposalFilePath)
	require.NoError(t, err)
	err = os.Remove(badProposalFilePath)
	require.NoError(t, err)
}

func TestGovClient_SubmitManageWhiteListProposal(t *testing.T) {
	// build the manage white list proposal JSON file
	err := ioutil.WriteFile(manageWhiteListProposalFilePath, []byte(manageWhiteListProposalJSON), 0644)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(6)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Governance().SubmitManageWhiteListProposal(fromInfo, passWd, manageWhiteListProposalFilePath,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Governance().SubmitManageWhiteListProposal(fromInfo, passWd, manageWhiteListProposalFilePath[1:],
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// bad community pool spend proposal JSON file
	err = ioutil.WriteFile(badProposalFilePath, []byte(manageWhiteListProposalFilePath[1:]), 0644)
	require.NoError(t, err)
	require.Panics(t, func() {
		_, _ = mockCli.Governance().SubmitManageWhiteListProposal(fromInfo, passWd, badProposalFilePath, memo,
			accInfo.GetAccountNumber(), accInfo.GetSequence())
	})

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Governance().SubmitManageWhiteListProposal(fromInfo, passWd, manageWhiteListProposalFilePath,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().SubmitManageWhiteListProposal(fromInfo, "", manageWhiteListProposalFilePath,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	// remove the temporary files
	err = os.Remove(manageWhiteListProposalFilePath)
	require.NoError(t, err)
	err = os.Remove(badProposalFilePath)
	require.NoError(t, err)
}

func TestGovClient_Deposit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

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

	res, err := mockCli.Governance().Deposit(fromInfo, passWd, "100okt", memo, 1,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Governance().Deposit(fromInfo, passWd, "100okt", memo, 0,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().Deposit(fromInfo, passWd, "100", memo, 1,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().Deposit(fromInfo, "", "100okt", memo, 1,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Governance().Deposit(fromInfo, passWd, "100okt", memo, 1,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestGovClient_Vote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

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
		Return(mocks.DefaultMockSuccessTxResponse(), nil).Times(4)

	res, err := mockCli.Governance().Vote(fromInfo, passWd, "yes", memo, 1, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Governance().Vote(fromInfo, passWd, "Abstain", memo, 1, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)

	_, err = mockCli.Governance().Vote(fromInfo, passWd, "no", memo, 1, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)

	_, err = mockCli.Governance().Vote(fromInfo, passWd, "no_with_veto", memo, 1, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)

	// error
	_, err = mockCli.Governance().Vote(fromInfo, passWd, "yes", memo, 0, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().Vote(fromInfo, "", "yes", memo, 1, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Governance().Vote(fromInfo, passWd, "", memo, 1, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Governance().Vote(fromInfo, passWd, "yes", memo, 1, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)
}
