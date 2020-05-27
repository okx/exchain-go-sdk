package governance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/auth"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

const (
	addr      = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okchainpub1addwnpepqgzuks5c07kfce85e0t0x8qkuvvxu874965ruafn6svhjrhswt0lgdj85lv"
	mnemonic  = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	memo      = "my memo"
)

const (
	textProposalFilePath    = "./text_proposal.json"
	badProposalFilePath = "./bad_proposal.json"
	textProposalJSON        = `{"title": "Text Proposal", "description": "text proposal description","proposal_type": "Text","deposit": "100okt"}`
)

func TestGovClient_SubmitTextProposal(t *testing.T) {
	// build the text proposal JSON file
	err := ioutil.WriteFile(textProposalFilePath, []byte(textProposalJSON), 0644)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, nil)

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
