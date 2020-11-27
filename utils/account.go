package utils

import (
	"errors"
	"fmt"
	"github.com/bartekn/go-bip39"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/okex/okexchain-go-sdk/types/tx"
	"log"
	"strings"
)

const (
	mnemonicEntropySize = 128
)

// CreateAccount creates a random key info with the given name and password
func CreateAccount(name, passWd string) (info keys.Info, mnemo string, err error) {
	if len(name) == 0 {
		name = "alice"
		log.Println("Default name : \"OKer\"")
	}

	if len(passWd) == 0 {
		passWd = "12345678"
		log.Println("Default passWd : \"12345678\"")
	}

	// TODO
	mnemo, err = GenerateMnemonic()
	if err != nil {
		return
	}

	hdPath := keys.CreateHDPath(0, 0).String()
	info, err = tx.Kb.CreateAccount(name, mnemo, "", passWd, hdPath, keys.Secp256k1)
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
		name = "alice"
		log.Println("Default name : \"alice\"")
	}

	if len(passWd) == 0 {
		passWd = "12345678"
		log.Println("Default passWd : \"12345678\"")
	}

	if strings.Contains(mnemonic, " ") && !bip39.IsMnemonicValid(mnemonic) {
		return info, mnemo, errors.New("failed. mnemonic is invalid")
	}

	hdPath := keys.CreateHDPath(0, 0).String()
	info, err = tx.Kb.CreateAccount(name, mnemonic, "", passWd, hdPath, keys.Secp256k1)
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
	return
	//derivedPrivSlice, err := hex.DecodeString(privateKey)
	//if err != nil {
	//	return
	//}
	//derivedPriv, err := sliceToArray(derivedPrivSlice)
	//if err != nil {
	//	return
	//}
	//priv := secp256k1.PrivKeySecp256k1(derivedPriv)
	//
	//privateKeyArmor := mintkey.EncryptArmorPrivKey(priv, passWd)
	//return keys.NewLocalInfo(name, priv.PubKey(), privateKeyArmor), err
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
