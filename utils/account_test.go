package utils

import (
	"fmt"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	info, mnemo, err := CreateAccount("alice", "12345678")
	assertNotEqual(t, err, nil)
	assertEqual(t,info,nil)
	assertEqual(t,mnemo,"")

	fmt.Println(info.GetAddress().String())
	fmt.Println(info.GetName())
	fmt.Println(info.GetPubKey())
	fmt.Println(mnemo)
}

func assertNotEqual(t *testing.T, err, b interface{}) {
	if err != b {
		t.Errorf("test failed: %s", err)
	}
}

func assertEqual(t *testing.T, err, b interface{}) {
	if err == b {
		t.Errorf("test failed: %s", err)
	}
}
