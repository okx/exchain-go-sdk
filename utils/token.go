package utils

import (
	"errors"
	"fmt"
	tokentypes "github.com/okex/okchain-go-sdk/module/token/types"
	"github.com/okex/okchain-go-sdk/types"
	"regexp"
	"strings"
)

var (
	reDecAmt    = `[[:digit:]]*\.?[[:digit:]]+`
	reSpc       = `[[:space:]]*`
	reDnmString = `[a-z][a-z0-9]{0,5}(\-[a-z0-9]{3})?`
	reDecCoin   = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reDecAmt, reSpc, reDnmString))
	ReDnm       = regexp.MustCompile(fmt.Sprintf(`^%s$`, reDnmString))
)

func ParseDecCoins(coinsStr string) (types.DecCoins, error) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return nil, nil
	}

	coinStrs := strings.Split(coinsStr, ",")
	coins := make(types.DecCoins, len(coinStrs))
	for i, coinStr := range coinStrs {
		coin, err := ParseDecCoin(coinStr)
		if err != nil {
			return nil, err
		}

		coins[i] = coin
	}

	// sort coins for determinism
	coins.Sort()

	// validate coins before returning
	if !coins.IsValid() {
		return nil, fmt.Errorf("parsed decimal coins are invalid: %#v", coins)
	}

	return coins, nil
}

// ParseDecCoin parses a decimal coin from a string, returning an error if invalid
// An empty string is considered invalid
func ParseDecCoin(coinStr string) (coin types.DecCoin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reDecCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		return coin, fmt.Errorf("invalid decimal coin expression: %s", coinStr)
	}

	amountStr, denomStr := matches[1], matches[2]

	amount, err := types.NewDecFromStr(amountStr)
	if err != nil {
		return coin, fmt.Errorf("failed to parse decimal coin amount: %s, %s", amountStr, err.Error())
	}

	if err := validateDenom(denomStr); err != nil {
		return coin, fmt.Errorf("invalid denom cannot contain upper case characters or spaces: %s", err)
	}

	return types.NewDecCoinFromDec(denomStr, amount), nil
}

// ParseTransfersStr parses the whole multi-send info string into TransferUnit
// Example:
// `addr1 1okt
// 	addr2 2okt`
func ParseTransfersStr(str string) ([]tokentypes.TransferUnit, error) {
	strs := strings.Split(strings.TrimSpace(str), "\n")
	transLen := len(strs)
	transfers := make([]tokentypes.TransferUnit, transLen)

	for i := 0; i < transLen; i++ {
		s := strings.Split(strs[i], " ")
		if len(s) != 2 {
			return nil, errors.New("invalid text to parse")
		}
		addrStr, coinStr := s[0], s[1]

		to, err := types.AccAddressFromBech32(addrStr)
		if err != nil {
			return nil, err
		}

		coins, err := ParseDecCoins(coinStr)
		if err != nil {
			return nil, err
		}

		transfers[i] = tokentypes.NewTransferUnit(to, coins)
	}

	return transfers, nil
}

func validateDenom(denom string) error {
	if !ReDnm.MatchString(denom) {
		return errors.New("illegal characters")
	}
	return nil
}
