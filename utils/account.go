package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/okex/okchain-go-sdk/common/libs/pkg/errors"
	"github.com/okex/okchain-go-sdk/crypto/go-bip39"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/crypto/keys/mintkey"
	"github.com/okex/okchain-go-sdk/types/tx"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

const (
	mnemonicEntropySize     = 128
	bcryptSecurityParameter = 12
	defaultKeyDBName        = "keys"
)

func CreateAccount(name, passWd string) (keys.Info, string, error) {
	if len(name) == 0 {
		name = "OKer"
		fmt.Println("Default name : \"OKer\"")
	}

	if len(passWd) == 0 {
		passWd = "12345678"
		fmt.Println("Default passWd : \"12345678\"")
	}

	var entropySeed []byte
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return nil, "", fmt.Errorf("bip39.NewEntropy err : %s", err.Error())
	}

	mnemo, err := bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return nil, "", fmt.Errorf("bip39.NewMnemonic err : %s", err.Error())
	}

	info, err := tx.Kb.CreateAccount(name, mnemo, "", passWd, 0, 0)
	if err != nil {
		return nil, "", fmt.Errorf("Kb.CreateAccount err : %s", err.Error())
	}

	return info, mnemo, nil

}

func CreateAccountWithMnemo(mnemo, name, passWd string) (keys.Info, string, error) {
	if len(mnemo) == 0 {
		return nil, "", errors.New("err : no mnemo input")
	}

	if len(name) == 0 {
		name = "OKer"
		fmt.Println("Default name : \"OKer\"")
	}

	if len(passWd) == 0 {
		passWd = "12345678"
		fmt.Println("Default passWd : \"12345678\"")
	}

	if !bip39.IsMnemonicValid(mnemo) {
		return nil, "", errors.New("err : mnemonic is not valid")
	}

	info, err := tx.Kb.CreateAccount(name, mnemo, "", passWd, 0, 0)
	if err != nil {
		return nil, "", fmt.Errorf("Kb.CreateAccount err : %s", err.Error())
	}

	return info, mnemo, nil
}

func CreateAccountWithPrivateKey(privateKey, name, passWd string) (keys.Info, error) {
	if len(privateKey) == 0 {
		return nil, errors.New("Empty privateKey")
	}
	derivedPrivSlice, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	derivedPriv, err := slice2Array(derivedPrivSlice)
	if err != nil {
		return nil, err
	}
	priv := secp256k1.PrivKeySecp256k1(derivedPriv)

	privateKeyArmor := mintkey.EncryptArmorPrivKey(priv, passWd)
	return keys.NewLocalInfo(name, priv.PubKey(), privateKeyArmor), nil
}

func GenerateMnemonic() (string, error) {
	var entropySeed []byte
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return "", fmt.Errorf("bip39.NewEntropy err : %s", err.Error())
	}

	mnemo, err := bip39.NewMnemonic(entropySeed[:])
	if err != nil {
		return "", fmt.Errorf("bip39.NewMnemonic err : %s", err.Error())
	}
	return mnemo, nil
}
