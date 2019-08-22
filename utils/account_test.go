package utils

import (
	"fmt"
	"testing"
)

const (
	name     = "alice"
	passWd   = "12345678"
	mnemonic = "sustain hole urban away boy core lazy brick wait drive tiger tell"
)

func TestCreateAccount(t *testing.T) {
	info, mnemo, err := CreateAccount("", "")
	assertNotEqual(t, err, nil)
	assertEqual(t, info, nil)
	assertEqual(t, mnemo, "")

	fmt.Println(info.GetAddress().String())
	fmt.Println(info.GetName())
	fmt.Println(info.GetPubKey())
	fmt.Println(mnemo)
}

func TestCreateAccountWithMnemo(t *testing.T) {
	info, mnemo, err := CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	assertEqual(t, info, nil)
	assertEqual(t, mnemo, "")

	fmt.Println(info.GetAddress().String())
	fmt.Println(info.GetName())
	fmt.Println(info.GetPubKey())
	fmt.Println(mnemo)
}

func TestCreateAccountWithPrivateKey(t *testing.T) {
	privateKeyStr, err := GeneratePrivateKeyFromMnemo(mnemonic)
	assertNotEqual(t, err, nil)
	fmt.Println(privateKeyStr)
	info, err := CreateAccountWithPrivateKey(privateKeyStr, name, passWd)
	assertNotEqual(t, err, nil)
	fmt.Println(info.GetAddress().String())
	fmt.Println(info.GetName())
	fmt.Println(info.GetPubKey())
}

func TestGenerateMnemonic(t *testing.T) {
	mnemo, err := GenerateMnemonic()
	assertNotEqual(t, err, nil)
	fmt.Println(mnemo)
}

func TestGeneratePrivateKeyFromMnemo(t *testing.T) {
	privateKey, err := GeneratePrivateKeyFromMnemo(mnemonic)
	assertNotEqual(t, err, nil)
	fmt.Println(privateKey)
}

func assertNotEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("test failed: %s", a)
	}
}

func assertEqual(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("test failed: %s", a)
	}
}
