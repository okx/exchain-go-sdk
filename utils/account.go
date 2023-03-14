package utils

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/bartekn/go-bip39"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/okex/exchain-go-sdk/types/tx"
	"github.com/okx/okbchain/app/crypto/ethsecp256k1"
	"github.com/okx/okbchain/app/crypto/hd"
	exchain "github.com/okx/okbchain/app/types"
	"github.com/okx/okbchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	tmamino "github.com/okx/okbchain/libs/tendermint/crypto/encoding/amino"
)

const (
	defaultName         = "alice"
	defaultPassWd       = "12345678"
	mnemonicEntropySize = 128
	defaultCointype     = 60
)

func init() {
	tmamino.RegisterKeyType(ethsecp256k1.PubKey{}, ethsecp256k1.PubKeyName)
	tmamino.RegisterKeyType(ethsecp256k1.PrivKey{}, ethsecp256k1.PrivKeyName)
	// set the address prefixes
	config := sdk.GetConfig()
	exchain.SetBech32Prefixes(config)
	config.SetCoinType(defaultCointype)
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

	log.Printf("New mnemonic: \"%s\". Be sure to remember that!\n", mnemo)

	return
}

// GeneratePrivateKeyFromMnemo converts mnemonic to private key
func GeneratePrivateKeyFromMnemo(mnemonic string) (privKey string, err error) {
	if len(mnemonic) == 0 {
		return privKey, errors.New("failed. no mnemonic input")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return privKey, errors.New("failed. mnemonic is invalid")
	}

	hdPath := keys.CreateHDPath(0, 0).String()
	if _, err = tx.Kb.CreateAccount(defaultName, mnemonic, "", defaultPassWd, hdPath, hd.EthSecp256k1); err != nil {
		return
	}

	priv, err := tx.Kb.ExportPrivateKeyObject(defaultName, defaultPassWd)
	if err != nil {
		return
	}

	privBytes, ok := priv.(ethsecp256k1.PrivKey)
	if !ok {
		return privKey, fmt.Errorf("invalid private key type, must be Ethereum key: %T", privKey)
	}

	return strings.ToUpper(hexutil.Encode(ethcrypto.FromECDSA(privBytes.ToECDSA()))[2:]), err
}

func GenerateEthPrivateKeyFromMnemo(mnemonic string) (privKey ethsecp256k1.PrivKey, err error) {
	if len(mnemonic) == 0 {
		return privKey, errors.New("failed. no mnemonic input")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return privKey, errors.New("failed. mnemonic is invalid")
	}

	hdPath := keys.CreateHDPath(0, 0).String()
	if _, err = tx.Kb.CreateAccount(defaultName, mnemonic, "", defaultPassWd, hdPath, hd.EthSecp256k1); err != nil {
		return
	}

	priv, err := tx.Kb.ExportPrivateKeyObject(defaultName, defaultPassWd)
	if err != nil {
		return
	}

	privBytes, ok := priv.(ethsecp256k1.PrivKey)
	if !ok {
		return privKey, fmt.Errorf("invalid private key type, must be Ethereum key: %T", privKey)
	}

	return privBytes, nil
}
