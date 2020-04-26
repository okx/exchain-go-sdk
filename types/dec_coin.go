package types

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	// Denominations can be 3 ~ 16 characters long
	reDnmString = `[a-z][a-z0-9]{0,9}(\-[a-z0-9]{3})?`
	reDecAmt    = `[[:digit:]]*\.?[[:digit:]]+`
	reSpc       = `[[:space:]]*`
	reDecCoin   = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reDecAmt, reSpc, reDnmString))
	reDnm       = regexp.MustCompile(fmt.Sprintf(`^%s$`, reDnmString))
)

// DecCoin defines a coin which can have additional decimal points
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

// IsZero returns if the DecCoin amount is zero
func (coin DecCoin) IsZero() bool {
	return coin.Amount.IsZero()
}

// Add adds amounts of two decimal coins with same denom
func (coin DecCoin) Add(coinB DecCoin) DecCoin {
	if coin.Denom != coinB.Denom {
		panic(fmt.Sprintf("coin denom different: %v %v\n", coin.Denom, coinB.Denom))
	}
	return DecCoin{coin.Denom, coin.Amount.Add(coinB.Amount)}
}

// DecCoins defines a slice of coins with decimal values
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

// IsZero returns whether all coins are zero
func (coins DecCoins) IsZero() bool {
	for _, coin := range coins {
		if !coin.Amount.IsZero() {
			return false
		}
	}
	return true
}

// Add adds two sets of DecCoins
// NOTE: Add operates under the invariant that coins are sorted by
// denominations.
// CONTRACT: Add will never return Coins where one Coin has a non-positive
// amount. In otherwords, IsValid will always return true.
func (coins DecCoins) Add(coinsB DecCoins) DecCoins {
	return coins.safeAdd(coinsB)
}

func (coins DecCoins) safeAdd(coinsB DecCoins) DecCoins {
	sum := ([]DecCoin)(nil)
	indexA, indexB := 0, 0
	lenA, lenB := len(coins), len(coinsB)

	for {
		if indexA == lenA {
			if indexB == lenB {
				// return nil coins if both sets are empty
				return sum
			}

			// return set B (excluding zero coins) if set A is empty
			return append(sum, removeZeroDecCoins(coinsB[indexB:])...)
		} else if indexB == lenB {
			// return set A (excluding zero coins) if set B is empty
			return append(sum, removeZeroDecCoins(coins[indexA:])...)
		}

		coinA, coinB := coins[indexA], coinsB[indexB]

		switch strings.Compare(coinA.Denom, coinB.Denom) {
		case -1: // coin A denom < coin B denom
			if !coinA.IsZero() {
				sum = append(sum, coinA)
			}

			indexA++

		case 0: // coin A denom == coin B denom
			res := coinA.Add(coinB)
			if !res.IsZero() {
				sum = append(sum, res)
			}

			indexA++
			indexB++

		case 1: // coin A denom > coin B denom
			if !coinB.IsZero() {
				sum = append(sum, coinB)
			}

			indexB++
		}
	}
}

// NewDecCoins creates a new instance of DecCoins
func NewDecCoins(coins ...DecCoin) DecCoins {
	// remove zeroes
	newCoins := removeZeroDecCoins(coins)
	if len(newCoins) == 0 {
		return DecCoins{}
	}

	newCoins.Sort()

	// detect duplicate Denoms
	if dupIndex := findDup(newCoins); dupIndex != -1 {
		panic(fmt.Errorf("find duplicate denom: %s", newCoins[dupIndex]))
	}

	if !newCoins.IsValid() {
		panic(fmt.Errorf("invalid dec coin set: %s", newCoins))
	}

	return newCoins
}

// ParseDecCoins parses DecCoins from string
func ParseDecCoins(coinsStr string) (DecCoins, error) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return nil, nil
	}

	coinStrs := strings.Split(coinsStr, ",")
	coins := make(DecCoins, len(coinStrs))
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
func ParseDecCoin(coinStr string) (coin DecCoin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reDecCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		return coin, fmt.Errorf("invalid decimal coin expression: %s", coinStr)
	}

	amountStr, denomStr := matches[1], matches[2]

	amount, err := NewDecFromStr(amountStr)
	if err != nil {
		return coin, fmt.Errorf("failed to parse decimal coin amount: %s, %s", amountStr, err.Error())
	}

	if err := validateDenom(denomStr); err != nil {
		return coin, fmt.Errorf("invalid denom cannot contain upper case characters or spaces: %s", err)
	}

	return NewDecCoinFromDec(denomStr, amount), nil
}

func findDup(coins DecCoins) int {
	if len(coins) <= 1 {
		return -1
	}

	prevDenom := coins[0].Denom
	for i := 1; i < len(coins); i++ {
		if coins[i].Denom == prevDenom {
			return i
		}
		prevDenom = coins[i].Denom
	}

	return -1
}

func removeZeroDecCoins(coins DecCoins) DecCoins {
	i, l := 0, len(coins)
	for i < l {
		if coins[i].IsZero() {
			// remove coin
			coins = append(coins[:i], coins[i+1:]...)
			l--
		} else {
			i++
		}
	}

	return coins[:i]
}

func mustValidateDenom(denom string) {
	if err := validateDenom(denom); err != nil {
		panic(err)
	}
}

func validateDenom(denom string) error {
	if !reDnm.MatchString(denom) {
		return fmt.Errorf("invalid denom: %s", denom)
	}
	return nil
}
