package utils

import (
	"fmt"
	"github.com/ok-chain/gosdk/common/libs/pkg/errors"
	"github.com/ok-chain/gosdk/crypto/go-bip39"
	"github.com/ok-chain/gosdk/crypto/keys"
)

const mnemonicEntropySize = 128

var (
	Kb keys.Keybase
)

func init() {
	Kb = keys.NewInMemory()
}

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

	info, err := Kb.CreateAccount(name, mnemo, "", passWd, 0, 0)
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

	info, err := Kb.CreateAccount(name, mnemo, "", passWd, 0, 0)
	if err != nil {
		return nil, "", fmt.Errorf("Kb.CreateAccount err : %s", err.Error())
	}

	return info, mnemo, nil
}


