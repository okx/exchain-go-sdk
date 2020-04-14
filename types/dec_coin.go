package types

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	reDecAmt    = `[[:digit:]]*\.?[[:digit:]]+`
	reSpc       = `[[:space:]]*`
	reDnmString = `[a-z][a-z0-9]{0,9}(\-[a-z0-9]{3})?`
	reDecCoin   = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reDecAmt, reSpc, reDnmString))
)

type DecCoin struct {
	Denom  string `json:"denom"`
	Amount Dec    `json:"amount"`
}

// NewDecCoinFromDec returns a dec coin from a Dec number and string denom
func NewDecCoinFromDec(denom string, amount Dec) DecCoin {
	mustValidateDenom(denom)

	if amount.LT(ZeroDec()) {
		panic(fmt.Sprintf("negative decimal coin amount: %v\n", amount))
	}

	return DecCoin{
		Denom:  denom,
		Amount: amount,
	}
}

// IsPositive returns true if coin amount is positive
func (coin DecCoin) IsPositive() bool {
	return coin.Amount.IsPositive()
}

// IsNegative returns true if the coin amount is negative and false otherwise
func (coin DecCoin) IsNegative() bool {
	return coin.Amount.Sign() == -1
}

type DecCoins []DecCoin

//nolint
func (coins DecCoins) Len() int           { return len(coins) }
func (coins DecCoins) Less(i, j int) bool { return coins[i].Denom < coins[j].Denom }
func (coins DecCoins) Swap(i, j int)      { coins[i], coins[j] = coins[j], coins[i] }

// Sort is a helper function to sort the set of decimal coins in-place.
func (coins DecCoins) Sort() DecCoins {
	sort.Sort(coins)
	return coins
}

// IsValid asserts the DecCoins are sorted, have positive amount and Denom does not contain upper case characters
func (coins DecCoins) IsValid() bool {
	switch len(coins) {
	case 0:
		return true

	case 1:
		if err := validateDenom(coins[0].Denom); err != nil {
			return false
		}
		return coins[0].IsPositive()

	default:
		// check single coin case
		if !(DecCoins{coins[0]}).IsValid() {
			return false
		}

		lowDenom := coins[0].Denom
		for _, coin := range coins[1:] {
			if strings.ToLower(coin.Denom) != coin.Denom {
				return false
			}
			if coin.Denom <= lowDenom {
				return false
			}
			if !coin.IsPositive() {
				return false
			}

			// we compare each coin against the last denom
			lowDenom = coin.Denom
		}

		return true
	}
}

// IsAllPositive returns true if there is at least one coin and all currencies
func (coins DecCoins) IsAllPositive() bool {
	if len(coins) == 0 {
		return false
	}

	for _, coin := range coins {
		if !coin.IsPositive() {
			return false
		}
	}

	return true
}