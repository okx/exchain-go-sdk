package backend

import (
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBackendClient_QueryCandles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewBackendClient(mockCli.MockBaseClient))



}
