package exposed

import (
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"time"
)

// Auth shows the expected behavior for inner auth client
type Auth interface {
	types.Module
	AuthQuery
}

// AuthQuery shows the expected query behavior for inner auth client
type AuthQuery interface {
	QueryAccount(accAddrStr string) (Account, error)
}

// Account shows expected behavior that an account has
type Account interface {
	GetAddress() types.AccAddress
	SetAddress(types.AccAddress) error

	GetPubKey() crypto.PubKey
	SetPubKey(crypto.PubKey) error

	GetAccountNumber() uint64
	SetAccountNumber(uint64) error

	GetSequence() uint64
	SetSequence(uint64) error

	GetCoins() types.Coins
	SetCoins(types.Coins) error

	SpendableCoins(blockTime time.Time) types.Coins

	String() string
}

// BaseAccount is the struct of account info
type BaseAccount struct {
	Address       types.AccAddress `json:"address"`
	Coins         types.Coins      `json:"coins"`
	PubKey        crypto.PubKey    `json:"public_key"`
	AccountNumber uint64           `json:"account_number"`
	Sequence      uint64           `json:"sequence"`
}

// String returns a human readable string representation of BaseAccount
func (acc BaseAccount) String() string {
	var pubkey string

	if acc.PubKey != nil {
		pubkey = types.MustBech32ifyAccPub(acc.PubKey)
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

// GetAddress gets the acc address of the account
func (acc BaseAccount) GetAddress() types.AccAddress {
	return acc.Address
}

// SetAddress sets the acc address of the account
func (acc *BaseAccount) SetAddress(addr types.AccAddress) error {
	if len(acc.Address) != 0 {
		return errors.New("cannot override BaseAccount address")
	}
	acc.Address = addr
	return nil
}

// GetPubKey gets the raw public key of the account
func (acc BaseAccount) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

// SetPubKey sets the raw public key of the account
func (acc *BaseAccount) SetPubKey(pubKey crypto.PubKey) error {
	acc.PubKey = pubKey
	return nil
}

// GetCoins gets the coins that the account owns
func (acc *BaseAccount) GetCoins() types.Coins {
	return acc.Coins
}

// SetCoins sets the coins that the account owns
func (acc *BaseAccount) SetCoins(coins types.Coins) error {
	acc.Coins = coins
	return nil
}

// GetAccountNumber gets the account number of the account
func (acc *BaseAccount) GetAccountNumber() uint64 {
	return acc.AccountNumber
}

// SetAccountNumber sets the account number of the account
func (acc *BaseAccount) SetAccountNumber(accNumber uint64) error {
	acc.AccountNumber = accNumber
	return nil
}

// GetSequence gets the sequence number of the account
func (acc *BaseAccount) GetSequence() uint64 {
	return acc.Sequence
}

// SetSequence sets the sequence number of the account
func (acc *BaseAccount) SetSequence(seq uint64) error {
	acc.Sequence = seq
	return nil
}

// SpendableCoins gets the spendable coins that the account owns
func (acc *BaseAccount) SpendableCoins(_ time.Time) types.Coins {
	return acc.GetCoins()
}
