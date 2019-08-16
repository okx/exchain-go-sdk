package types

import (
	"errors"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
	"time"
)


type Account interface {
	GetAddress() AccAddress
	SetAddress(AccAddress) error // errors if already set.

	GetPubKey() crypto.PubKey // can return nil.
	SetPubKey(crypto.PubKey) error

	GetAccountNumber() uint64
	SetAccountNumber(uint64) error

	GetSequence() uint64
	SetSequence(uint64) error

	GetCoins() Coins
	SetCoins(Coins) error

	// Calculates the amount of coins that can be sent to other accounts given
	// the current time.
	SpendableCoins(blockTime time.Time) Coins

	// Ensure that account implements stringer
	String() string
}

// implementation of Account
type BaseAccount struct {
	Address       AccAddress    `json:"address"`
	Coins         Coins         `json:"coins"`
	PubKey        crypto.PubKey `json:"public_key"`
	AccountNumber uint64        `json:"account_number"`
	Sequence      uint64        `json:"sequence"`
}

func (acc BaseAccount) String() string {
	var pubkey string

	if acc.PubKey != nil {
		pubkey = MustBech32ifyAccPub(acc.PubKey)
	}

	return fmt.Sprintf(`Account:
  Address:       %s
  Pubkey:        %s
  Coins:         %s
  AccountNumber: %d
  Sequence:      %d`,
		acc.Address, pubkey, acc.Coins, acc.AccountNumber, acc.Sequence,
	)
}

func (acc BaseAccount) GetAddress() AccAddress {
	return acc.Address
}

func (acc *BaseAccount) SetAddress(addr AccAddress) error {
	if len(acc.Address) != 0 {
		return errors.New("cannot override BaseAccount address")
	}
	acc.Address = addr
	return nil
}

func (acc BaseAccount) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

func (acc *BaseAccount) SetPubKey(pubKey crypto.PubKey) error {
	acc.PubKey = pubKey
	return nil
}

func (acc *BaseAccount) GetCoins() Coins {
	return acc.Coins
}

func (acc *BaseAccount) SetCoins(coins Coins) error {
	acc.Coins = coins
	return nil
}

func (acc *BaseAccount) GetAccountNumber() uint64 {
	return acc.AccountNumber
}

func (acc *BaseAccount) SetAccountNumber(accNumber uint64) error {
	acc.AccountNumber = accNumber
	return nil
}

func (acc *BaseAccount) GetSequence() uint64 {
	return acc.Sequence
}

func (acc *BaseAccount) SetSequence(seq uint64) error {
	acc.Sequence = seq
	return nil
}

func (acc *BaseAccount) SpendableCoins(_ time.Time) Coins {
	return acc.GetCoins()
}

