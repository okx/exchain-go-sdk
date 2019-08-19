package types

import (
	"encoding/json"
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

func (coin Coin) IsZero() bool {
	return coin.Amount.IsZero()
}

func (coin Coin) IsPositive() bool {
	return coin.Amount.Sign() == 1
}

func (coin Coin) String() string {
	dec := NewDecFromIntWithPrec(coin.Amount, Precision)
	return fmt.Sprintf("%s%v", dec, coin.Denom)
}
// MarshalJSON marshals the coin
func (coin Coin) MarshalJSON() ([]byte, error) {
	type Alias Coin
	return json.Marshal(&struct {
		Denom  string `json:"denom"`
		Amount Dec    `json:"amount"`
	}{
		coin.Denom,
		NewDecFromIntWithPrec(coin.Amount, Precision),
	})
}

func (coin *Coin) UnmarshalJSON(data []byte) error {
	c := &struct {
		Denom  string `json:"denom"`
		Amount Dec    `json:"amount"`
	}{}
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	coin.Denom = c.Denom
	coin.Amount = NewIntFromBigInt(c.Amount.Int)
	return nil
}
type Coins []Coin

// NewCoins constructs a new coin set.
func NewCoins(coins ...Coin) Coins {
	// remove zeroes
	newCoins := removeZeroCoins(Coins(coins))
	if len(newCoins) == 0 {
		return Coins{}
	}

	newCoins.Sort()

	// detect duplicate Denoms
	if dupIndex := findDup(newCoins); dupIndex != -1 {
		panic(fmt.Errorf("find duplicate denom: %s", newCoins[dupIndex]))
	}

	if !newCoins.IsValid() {
		panic(fmt.Errorf("invalid coin set: %s", newCoins))
	}

	return newCoins
}

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

func (coins Coins) IsAllPositive() bool {
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

func (coins Coins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}
	return out[:len(out)-1]
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

func removeZeroCoins(coins Coins) Coins {
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

func findDup(coins Coins) int {
	if len(coins) <= 1 {
		return -1
	}

	prevDenom := coins[0]
	for i := 1; i < len(coins); i++ {
		if coins[i] == prevDenom {
			return i
		}
	}

	return -1
}
