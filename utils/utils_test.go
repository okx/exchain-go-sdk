package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	valAddrStr = "okexchainvaloper10q0rk5qnyag7wfvvt7rtphlw589m7frshchly8"
)

func TestParseValAddresses(t *testing.T) {
	valAddrsStr := []string{valAddrStr}
	valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
	require.NoError(t, err)

	valAddrs, err := ParseValAddresses(valAddrsStr)
	require.NoError(t, err)
	require.Equal(t, 1, len(valAddrs))
	require.Equal(t, valAddr, valAddrs[0])

	// bad val address
	valAddrsStr = append(valAddrsStr, valAddrStr[1:])
	_, err = ParseValAddresses(valAddrsStr)
	require.Error(t, err)
}

func TestGeneratePrivateKeyFromMnemo(t *testing.T) {
	priKey, err := GeneratePrivateKeyFromMnemo(defaultMnemonic)
	require.NoError(t, err)
	require.Equal(t, defaultPrivateKey, priKey)

	// bad mnemonic, add one word in it
	_, err = GeneratePrivateKeyFromMnemo(fmt.Sprintf("%s %s", defaultMnemonic, "offer"))
	require.Error(t, err)
}

func TestGetStdTxFromFile(t *testing.T) {
	// data preparation
	addr, err := sdk.AccAddressFromBech32(accAddr1)
	require.NoError(t, err)
	feeCoins, err := sdk.ParseDecCoins("1024okt,2.048btc")
	require.NoError(t, err)
	stdFee := sdk.NewStdFee(20000, feeCoins)

	mockStdTx := sdk.StdTx{
		Msgs: []sdk.Msg{
			TestMsg{
				addr,
			},
		},
		Fee:  stdFee,
		Memo: defaultMemo,
	}

	// write data to file
	filePath := "./test_std_tx.json"
	err = ioutil.WriteFile(filePath, testCdc.MustMarshalJSON(mockStdTx), 0644)
	defer os.Remove(filePath)
	require.NoError(t, err)

	stdTx, err := GetStdTxFromFile(testCdc, filePath)
	require.NoError(t, err)
	require.Equal(t, mockStdTx, stdTx)

	_, err = GetStdTxFromFile(testCdc, filePath[1:])
	require.Error(t, err)

	// bad JSON bytes in file
	badFilePath := "./test_bad_std_tx.json"
	err = ioutil.WriteFile(badFilePath, testCdc.MustMarshalJSON(mockStdTx)[1:], 0644)
	defer os.Remove(badFilePath)
	require.NoError(t, err)
	require.Panics(t, func() {
		_, _ = GetStdTxFromFile(testCdc, badFilePath)
	})

}
