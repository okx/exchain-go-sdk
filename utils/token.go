package utils

import (
	"encoding/json"
	"errors"
	"fmt"
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

func ParseCoins(coinsStr string) (coins types.Coins, err error) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return nil, nil
	}

	coinStrs := strings.Split(coinsStr, ",")
	for _, coinStr := range coinStrs {
		coin, err := ParseCoin(coinStr)
		if err != nil {
			return nil, err
		}
		coins = append(coins, coin)
	}

	// Sort coins for determinism.
	coins.Sort()

	// Validate coins before returning.
	if !coins.IsValid() {
		return nil, fmt.Errorf("parseCoins invalid: %#v", coins)
	}

	return coins, nil
}

func ParseCoin(coinStr string) (coin types.Coin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reDecCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		return types.Coin{}, fmt.Errorf("invalid coin expression: %s", coinStr)
	}

	denomStr, amountStr := matches[2], matches[1]

	amount, err := types.NewDecFromStr(amountStr)
	if err != nil {
		return types.Coin{}, fmt.Errorf("failed to parse coin amount %s: %s", amountStr, err.Error())
	}

	if err := validateDenom(denomStr); err != nil {
		return types.Coin{}, fmt.Errorf("invalid denom cannot contain upper case characters or spaces: %s", err)
	}

	coin = types.NewCoin(denomStr, types.NewIntFromBigInt(amount.Int))

	return coin, nil
}

func StrToTransfers(str string) (transfers []types.TransferUnit, err error) {
	var transfer []types.Transfer
	err = json.Unmarshal([]byte(str), &transfer)
	if err != nil {
		return transfers, err
	}

	for _, trans := range transfer {
		var t types.TransferUnit
		to, err := types.AccAddressFromBech32(trans.To)
		if err != nil {
			return transfers, err
		}
		t.To = to
		t.Coins = AmountToCoins(trans.Amount)
		transfers = append(transfers, t)
	}
	return transfers, nil
}

func AmountToCoins(amount string) types.Coins {
	var res types.Coins
	res, _ = ParseCoins(amount)
	return res
}

func validateDenom(denom string) error {
	if !ReDnm.MatchString(denom) {
		return errors.New("illegal characters")
	}
	return nil
}
