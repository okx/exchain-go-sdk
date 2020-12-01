package dex

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"io/ioutil"
	"os"
	"testing"

	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module"
	"github.com/okex/okexchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
)

const (
	unsignedPath           = "./unsignedTx.json"
	signedPath             = "./signedTx.json"
	expectedUnsignedTxJSON = `{"type":"cosmos-sdk/StdTx","value":{"msg":[{"type":"okexchain/dex/MsgTransferTradingPairOwnership","value":{"from_address":"okexchain1kfs5q53jzgzkepqa6ual0z7f97wvxnkamr5vys","to_address":"okexchain16zgvph7qc3n4jvamq0lkv3y37k0hc5pw9hhhrs","product":"btc-000_okt","to_signature":{"pub_key":null,"signature":null}}}],"fee":{"amount":[{"denom":"okt","amount":"0.01000000"}],"gas":"200000"},"signatures":null,"memo":"my memo"}}`
	expectedSignedTxJSON   = `{"type":"cosmos-sdk/StdTx","value":{"msg":[{"type":"okexchain/dex/MsgTransferTradingPairOwnership","value":{"from_address":"okexchain1kfs5q53jzgzkepqa6ual0z7f97wvxnkamr5vys","to_address":"okexchain16zgvph7qc3n4jvamq0lkv3y37k0hc5pw9hhhrs","product":"btc-000_okt","to_signature":{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ap+v94JS5I9tQ8Mu2JmNKyo20EVVdNFpOUqjpxmQCuRS"},"signature":"8J3mYGFb7c4ZxDvOZ/9/wdez4YK4yHO0ERaVNgubYiZ1+AS+vDdRHHrNB4TZ+WvJa/XGGIJpCsw6HyaGK6QXIA=="}}}],"fee":{"amount":[{"denom":"okt","amount":"0.01000000"}],"gas":"200000"},"signatures":null,"memo":"my memo"}}`
)

func TestDexClient_GenerateUnsignedTransferOwnershipTx(t *testing.T) {
	// make dex client
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "0.01okt",
		200000, 0, "")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, nil, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient))
	//TODO
	//cdc := mockCli.GetCodec()
	dexClient := NewDexClient(module.NewBaseClient(nil, &config))

	// generate unsigned transfer ownership tx
	err = dexClient.GenerateUnsignedTransferOwnershipTx(product, addr, recAddr, memo, unsignedPath)
	require.NoError(t, err)

	err = dexClient.GenerateUnsignedTransferOwnershipTx("", addr, recAddr, memo, unsignedPath)
	require.Error(t, err)

	err = dexClient.GenerateUnsignedTransferOwnershipTx(product, addr[1:], recAddr, memo, unsignedPath)
	require.Error(t, err)

	err = dexClient.GenerateUnsignedTransferOwnershipTx(product, addr, recAddr[1:], memo, unsignedPath)
	require.Error(t, err)

	// read back to check
	// TODO
	//stdTx, err := utils.GetStdTxFromFile(cdc, unsignedPath)
	stdTx, err := utils.GetStdTxFromFile(nil, unsignedPath)
	require.NoError(t, err)

	var expectedStdTx authtypes.StdTx
	//cdc.MustUnmarshalJSON([]byte(expectedUnsignedTxJSON), &expectedStdTx)
	require.Equal(t, expectedStdTx, stdTx)

	// remove the temporary file: unsignedTx.json
	err = os.Remove(unsignedPath)
	require.NoError(t, err)
}

func TestDexClient_MultiSign(t *testing.T) {
	// set up unsignedTx.json
	err := ioutil.WriteFile(unsignedPath, []byte(expectedUnsignedTxJSON), 0644)
	require.NoError(t, err)

	// make dex client
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "0.01okt",
		200000, 0, "")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, nil, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient))
	//TODO
	//cdc := mockCli.GetCodec()
	dexClient := NewDexClient(module.NewBaseClient(nil, &config))

	// multisign
	recInfo, _, err := utils.CreateAccountWithMnemo(recMnemonic, name, passWd)
	require.NoError(t, err)
	err = dexClient.MultiSign(recInfo, passWd, unsignedPath, signedPath)
	require.NoError(t, err)

	err = dexClient.MultiSign(recInfo, passWd[1:], unsignedPath, signedPath)
	require.Error(t, err)

	err = dexClient.MultiSign(recInfo, passWd, unsignedPath, signedPath[1:])
	require.Error(t, err)

	err = dexClient.MultiSign(recInfo, "", unsignedPath, signedPath)
	require.Error(t, err)

	err = dexClient.MultiSign(recInfo, passWd, unsignedPath[1:], signedPath)
	require.Error(t, err)

	// read back to check
	//stdTx, err := utils.GetStdTxFromFile(cdc, signedPath)
	require.NoError(t, err)

	//var expectedStdTx authtypes.StdTx
	//cdc.MustUnmarshalJSON([]byte(expectedSignedTxJSON), &expectedStdTx)
	//require.Equal(t, expectedStdTx, stdTx)

	// remove the temporary files: unsignedTx.json and signedTx.json
	err = os.Remove(unsignedPath)
	require.NoError(t, err)
	err = os.Remove(signedPath)
	require.NoError(t, err)
}
