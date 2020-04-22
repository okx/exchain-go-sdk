package dex

import (
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

const (
	unsignedPath           = "./unsignedTx.json"
	signedPath             = "./signedTx.json"
	expectedUnsignedTxJSON = `{"type":"cosmos-sdk/StdTx","value":{"msg":[{"type":"okchain/dex/MsgTransferTradingPairOwnership","value":{"from_address":"okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz","to_address":"okchain1wux20ku36ntgtxpgm7my9863xy3fqs0xgh66d7","product":"btc-000_okt","to_signature":{"pub_key":null,"signature":null}}}],"fee":{"amount":[{"denom":"okt","amount":"0.01000000"}],"gas":"200000"},"signatures":null,"memo":"my memo"}}`
	expectedSignedTxJSON   = `{"type":"cosmos-sdk/StdTx","value":{"msg":[{"type":"okchain/dex/MsgTransferTradingPairOwnership","value":{"from_address":"okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz","to_address":"okchain1wux20ku36ntgtxpgm7my9863xy3fqs0xgh66d7","product":"btc-000_okt","to_signature":{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"AgNf7LwXZdAuTZs7XAY7lbdVaOvQlpGVyc2TV0QrSgj+"},"signature":"/3yLlZ96fDy97CRC/6kvooNJveVcANUTu6dbAEaaiipwKrpeUeNFoMFsBgxvM/+dWf2PckTezHRLKwk/ExQ+gA=="}}}],"fee":{"amount":[{"denom":"okt","amount":"0.01000000"}],"gas":"200000"},"signatures":null,"memo":"my memo"}}`
)

func TestDexClient_GenerateUnsignedTransferOwnershipTx(t *testing.T) {
	// make dex client
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, nil, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient))
	cdc := mockCli.GetCodec()
	dexClient := NewDexClient(module.NewBaseClient(cdc, &config))

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
	stdTx, err := utils.GetStdTxFromFile(cdc, unsignedPath)
	require.NoError(t, err)

	var expectedStdTx sdk.StdTx
	cdc.MustUnmarshalJSON([]byte(expectedUnsignedTxJSON), &expectedStdTx)
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
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, nil, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient))
	cdc := mockCli.GetCodec()
	dexClient := NewDexClient(module.NewBaseClient(cdc, &config))

	// multisign
	recInfo, _, err := utils.CreateAccountWithMnemo(recMnemonic, name, passWd)
	require.NoError(t, err)
	err = dexClient.MultiSign(recInfo, passWd, unsignedPath, signedPath)
	require.NoError(t, err)

	err = dexClient.MultiSign(recInfo, passWd[1:], unsignedPath, signedPath)
	require.Error(t, err)

	err = dexClient.MultiSign(recInfo, passWd, unsignedPath, signedPath[1:])
	require.Error(t, err)

	// read back to check
	stdTx, err := utils.GetStdTxFromFile(cdc, signedPath)
	require.NoError(t, err)

	var expectedStdTx sdk.StdTx
	cdc.MustUnmarshalJSON([]byte(expectedSignedTxJSON), &expectedStdTx)
	require.Equal(t, expectedStdTx, stdTx)

	// remove the temporary files: unsignedTx.json and signedTx.json
	err = os.Remove(unsignedPath)
	require.NoError(t, err)
	err = os.Remove(signedPath)
	require.NoError(t, err)

}
