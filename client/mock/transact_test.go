package mock

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/okex/okchain-go-sdk/client"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
	types2 "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"testing"
)

const (
	name   = "louisliu1024"
	passWd = "12345678"
	// send's mnemonic
	mnemonic = "cupboard hover grass method tourist fee clerk adapt prefer word card pretty"
	// target address
	toaddr = "okchain1e0wwl9thwjclh5vg6rmqhs739hwu0p0d6zg40q"
)

var accAddress, _ = types.AccAddressFromBech32("okchain1h643csgkdr8yzey22tekxwzjlxgn87sm0ust5u")
var accPubKey, _ = client.GetAccPubKeyBech32("okchainpub1addwnpepqwqz3e53lu43pf9zhpxyxrwl9zpzrt3l4y2fsj9dqq947p604z6wzs7u070")
var accounts = types.BaseAccount{
	Address: accAddress,
	Coins: types.Coins{
		types.Coin{
			Denom:  "okt",
			Amount: types.NewInt(100000),
		},
		types.Coin{
			Denom:  "acoin",
			Amount: types.NewInt(2000),
		},
	},
	PubKey:        accPubKey,
	AccountNumber: 2917,
	Sequence:      17,
}

func TestSend(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountBytes, err := cli.GetCdc().MarshalBinaryBare(accounts)
	var queryAccountResult ctypes.ResultABCIQuery
	queryAccountResult.Response.Code = 0
	queryAccountResult.Response.Value = accountBytes

	txResult := ctypes.ResultBroadcastTxCommit{
		CheckTx: types2.ResponseCheckTx{
			Code:                 0,
			Data:                 nil,
			Log:                  "",
			Info:                 "",
			GasWanted:            0,
			GasUsed:              0,
			Tags:                 nil,
			Codespace:            "",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
		DeliverTx: types2.ResponseDeliverTx{
			Code:                 0,
			Data:                 nil,
			Log:                  "",
			Info:                 "",
			GasWanted:            0,
			GasUsed:              0,
			Tags:                 nil,
			Codespace:            "",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
		Hash:   nil,
		Height: 1,
	}

	mockClient := NewMockClient(ctrl)
	gomock.InOrder(
		mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&queryAccountResult, nil),
		mockClient.EXPECT().BroadcastTxCommit(gomock.Any()).Return(&txResult, nil),
	)
	cli.SetClient(mockClient)

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assert.Equal(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assert.Equal(t, err, nil)
	res, err := cli.Send(fromInfo, passWd, toaddr, "1tokt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assert.Equal(t, err, nil)
	fmt.Println(res)
}

func TestNewOrder(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountBytes, err := cli.GetCdc().MarshalBinaryBare(accounts)
	var queryAccountResult ctypes.ResultABCIQuery
	queryAccountResult.Response.Code = 0
	queryAccountResult.Response.Value = accountBytes

	txResult := ctypes.ResultBroadcastTxCommit{
		CheckTx: types2.ResponseCheckTx{
			Code:                 0,
			Data:                 nil,
			Log:                  "",
			Info:                 "",
			GasWanted:            0,
			GasUsed:              0,
			Tags:                 nil,
			Codespace:            "",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
		DeliverTx: types2.ResponseDeliverTx{
			Code:                 0,
			Data:                 nil,
			Log:                  "",
			Info:                 "",
			GasWanted:            0,
			GasUsed:              0,
			Tags:                 nil,
			Codespace:            "",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
		Hash:   nil,
		Height: 1,
	}

	mockClient := NewMockClient(ctrl)
	gomock.InOrder(
		mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&queryAccountResult, nil),
		mockClient.EXPECT().BroadcastTxCommit(gomock.Any()).Return(&txResult, nil),
	)
	cli.SetClient(mockClient)

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assert.Equal(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assert.Equal(t, err, nil)
	res, err := cli.NewOrder(fromInfo, passWd, "xxb_okt", "BUY", "11.2", "1.23", "buy other token", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assert.Equal(t, err, nil)
	fmt.Println(res)
	//fmt.Println("orderId:", client.GetOrderIdFromResponse(&res))
}

func TestCancelOrder(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accountBytes, err := cli.GetCdc().MarshalBinaryBare(accounts)
	var queryAccountResult ctypes.ResultABCIQuery
	queryAccountResult.Response.Code = 0
	queryAccountResult.Response.Value = accountBytes

	txResult := ctypes.ResultBroadcastTxCommit{
		CheckTx: types2.ResponseCheckTx{
			Code:                 0,
			Data:                 nil,
			Log:                  "",
			Info:                 "",
			GasWanted:            0,
			GasUsed:              0,
			Tags:                 nil,
			Codespace:            "",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
		DeliverTx: types2.ResponseDeliverTx{
			Code:                 0,
			Data:                 nil,
			Log:                  "",
			Info:                 "",
			GasWanted:            0,
			GasUsed:              0,
			Tags:                 nil,
			Codespace:            "",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
		Hash:   nil,
		Height: 1,
	}

	mockClient := NewMockClient(ctrl)
	gomock.InOrder(
		mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&queryAccountResult, nil),
		mockClient.EXPECT().BroadcastTxCommit(gomock.Any()).Return(&txResult, nil),
	)
	cli.SetClient(mockClient)

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assert.Equal(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assert.Equal(t, err, nil)
	res, err := cli.CancelOrder(fromInfo, passWd, "ID0000004307-1", "cancel order", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assert.Equal(t, err, nil)
	fmt.Println(res)
}
