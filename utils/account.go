package utils

import (
	"errors"
	"fmt"
	"github.com/bartekn/go-bip39"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/tx"
	"github.com/okex/okexchain/app/crypto/hd"
	okexchain "github.com/okex/okexchain/app/types"
	"log"
	"strings"
)

const (
	defaultName         = "alice"
	defaultPassWd       = "12345678"
	mnemonicEntropySize = 128
)

func init() {
	// set the address prefixes
	config := sdk.GetConfig()
	okexchain.SetBech32Prefixes(config)
	okexchain.SetBip44CoinType(config)
}

// CreateAccount creates a random key info with the given name and password
func CreateAccount(name, passWd string) (info keys.Info, mnemo string, err error) {
	if len(name) == 0 {
		name = defaultName
		log.Printf("Default name: \"%s\"\n", name)
	}

	if len(passWd) == 0 {
		passWd = defaultPassWd
		log.Printf("Default password: \"%s\"\n", passWd)
	}

	mnemo, err = GenerateMnemonic()
	if err != nil {
		return
	}

	hdPath := keys.CreateHDPath(0, 0).String()
	info, err = tx.Kb.CreateAccount(name, mnemo, "", passWd, hdPath, hd.EthSecp256k1)
	if err != nil {
		return info, mnemo, fmt.Errorf("failed. Kb.CreateAccount err : %s", err.Error())
	}

	return
}

// CreateAccountWithMnemo creates the key info with the given mnemonic, name and password
func CreateAccountWithMnemo(mnemonic, name, passWd string) (info keys.Info, mnemo string, err error) {
	if len(mnemonic) == 0 {
		return info, mnemo, errors.New("failed. no mnemonic input")
	}

	if len(name) == 0 {
		name = defaultName
		log.Printf("Default name: \"%s\"\n", name)
	}

	if len(passWd) == 0 {
		passWd = defaultPassWd
		log.Printf("Default password: \"%s\"\n", passWd)
	}

	if strings.Contains(mnemonic, " ") && !bip39.IsMnemonicValid(mnemonic) {
		return info, mnemo, errors.New("failed. mnemonic is invalid")
	}

	hdPath := keys.CreateHDPath(0, 0).String()
	info, err = tx.Kb.CreateAccount(name, mnemonic, "", passWd, hdPath, hd.EthSecp256k1)
	if err != nil {
		return info, mnemonic, fmt.Errorf("failed. Kb.CreateAccount err : %s", err.Error())
	}

	return info, mnemonic, err
}

// CreateAccountWithPrivateKey creates the key info with the given privateKey string, name and password
func CreateAccountWithPrivateKey(privateKey, name, passWd string) (info keys.Info, err error) {
	if len(privateKey) == 0 {
		return info, errors.New("failed. empty privateKey")
	}

	if len(name) == 0 {
		name = "alice"
		log.Println("Default name : \"alice\"")
	}

	if len(passWd) == 0 {
		passWd = "12345678"
		log.Println("Default passWd : \"12345678\"")
	}

	hdPath := keys.CreateHDPath(0, 0).String()
	info, err = tx.Kb.CreateAccount(name, privateKey, "", passWd, hdPath, hd.EthSecp256k1)
	if err != nil {
		return info, fmt.Errorf("failed. Kb.CreateAccount err : %s", err.Error())
	}

	return
}

// GenerateMnemonic creates a random mnemonic
func GenerateMnemonic() (mnemo string, err error) {
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return mnemo, fmt.Errorf("failed. bip39.NewEntropy err : %s", err.Error())
	}

	mnemo, err = bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return mnemo, fmt.Errorf("failed. bip39.NewMnemonic err : %s", err.Error())
	}

	return
}
