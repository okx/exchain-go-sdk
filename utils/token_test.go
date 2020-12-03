package utils

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	accAddr1  = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	accAddr2  = "okexchain1qeh2fz0a4t78ylesd4cyd2mwt5wcfnfj98ev0u"
	coinsStr1 = "1.024okt"
	coinsStr2 = "2.048btc,2.048okt"
)

func TestParseTransfersStr(t *testing.T) {
	addr1, err := sdk.AccAddressFromBech32(accAddr1)
	require.NoError(t, err)
	addr2, err := sdk.AccAddressFromBech32(accAddr2)
	require.NoError(t, err)
	coins1, err := sdk.ParseDecCoins(coinsStr1)
	require.NoError(t, err)
	coins2, err := sdk.ParseDecCoins(coinsStr2)
	require.NoError(t, err)

	transfersStr := fmt.Sprintf("%s %s\n%s %s", accAddr1, coinsStr1, accAddr2, coinsStr2)
	transferUnits, err := ParseTransfersStr(transfersStr)
	require.NoError(t, err)
	require.Equal(t, 2, len(transferUnits))
	require.Equal(t, addr1, transferUnits[0].To)
	require.Equal(t, coins1, transferUnits[0].Coins)
	require.Equal(t, addr2, transferUnits[1].To)
	require.Equal(t, coins2, transferUnits[1].Coins)

	badTransfersStr := fmt.Sprintf("%s %s\n%s %s %s", accAddr1, coinsStr1, accAddr2, coinsStr2, "4.096eth")
	_, err = ParseTransfersStr(badTransfersStr)
	require.Error(t, err)

	badTransfersStr = fmt.Sprintf("%s %s\n%s %s", accAddr1[1:], coinsStr1, accAddr2, coinsStr2)
	_, err = ParseTransfersStr(badTransfersStr)
	require.Error(t, err)

	badTransfersStr = fmt.Sprintf("%s %s\n%s %s", accAddr1, "1.024", accAddr2, coinsStr2)
	_, err = ParseTransfersStr(badTransfersStr)
	require.Error(t, err)
}
