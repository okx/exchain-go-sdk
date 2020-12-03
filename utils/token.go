package utils

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/token/types"
	tokentypes "github.com/okex/okexchain/x/token/types"
	"strings"
)

// ParseTransfersStr parses the whole multi-send info string into TransferUnit
// Example:
// `addr1 1okt
// 	addr2 2okt`
func ParseTransfersStr(str string) ([]types.TransferUnit, error) {
	strs := strings.Split(strings.TrimSpace(str), "\n")
	transLen := len(strs)
	transfers := make([]types.TransferUnit, transLen)

	for i := 0; i < transLen; i++ {
		s := strings.Split(strs[i], " ")
		if len(s) != 2 {
			return nil, errors.New("invalid text to parse")
		}
		addrStr, coinStr := s[0], s[1]

		to, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return nil, err
		}

		coins, err := sdk.ParseDecCoins(coinStr)
		if err != nil {
			return nil, err
		}

		transfers[i] = tokentypes.TransferUnit{
			To:    to,
			Coins: coins,
		}
	}

	return transfers, nil
}
