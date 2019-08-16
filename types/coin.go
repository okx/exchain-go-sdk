package types

import (
	"fmt"
	"sort"
	"strings"
)

type Coin struct {
	Denom string `json:"denom"`

	// To allow the use of unsigned integers (see: #1273) a larger refactor will
	// need to be made. So we use signed integers for now with safety measures in
	// place preventing negative values being used.
	Amount Int `json:"amount"`
}

func NewCoin(denom string, amount Int) Coin {
	mustValidateDenom(denom)

	if amount.LT(ZeroInt()) {
		panic(fmt.Errorf("negative coin amount: %v", amount))
	}

	return Coin{
		Denom:  denom,
		Amount: amount,
	}
}

func (coin Coin) IsPositive() bool {
	return coin.Amount.Sign() == 1
}

type Coins []Coin


// Sort interface
func (coins Coins) Len() int           { return len(coins) }
func (coins Coins) Less(i, j int) bool { return coins[i].Denom < coins[j].Denom }
func (coins Coins) Swap(i, j int)      { coins[i], coins[j] = coins[j], coins[i] }

var _ sort.Interface = Coins{}
// Sort is a helper function to sort the set of coins inplace
func (coins Coins) Sort() Coins {
	sort.Sort(coins)
	return coins
}

// IsValid asserts the Coins are sorted, have positive amount,
// and Denom does not contain upper case characters.
func (coins Coins) IsValid() bool {
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
		if !(Coins{coins[0]}).IsValid() {
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

func validateDenom(denom string) error {
	//if !reDnm.MatchString(denom) {
	//	return errors.New("illegal characters")
	//}
	return nil
}

func mustValidateDenom(denom string) {
	if err := validateDenom(denom); err != nil {
		panic(err)
	}
}