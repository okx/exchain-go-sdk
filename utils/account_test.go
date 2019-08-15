package utils

import (
	"fmt"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	info, mnemo, err := CreateAccount("alice", "12345678")
	assertEqual(t, err, nil)
	fmt.Println(info.GetAddress().String())
	fmt.Println(info.GetName())
	fmt.Println(info.GetPubKey())
	fmt.Println(mnemo)
}

func assertEqual(t *testing.T, err, b interface{}) {
	if err != b {
		t.Errorf("test failed: %s", err)
	}
}
