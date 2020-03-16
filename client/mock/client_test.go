package mock

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/okex/okchain-go-sdk/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"testing"
)

func TestNewClient(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountAssertStr := `{"address":"okchain10q0rk5qnyag7wfvvt7rtphlw589m7frsmyq4ya","currencies":[{"symbol":"acoin","available":"10000000.00000000","freeze":"0","locked":"0"},{"symbol":"bcoin","available":"10000000.00000000","freeze":"0","locked":"0"}]}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(accountAssertStr)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	resp, err := cli.GetTokensInfoByAddr(addr)
	assert.Equal(t, err, nil)
	fmt.Println(resp)
}
