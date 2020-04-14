package auth

import (
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/types"
)

// QueryAccount gets the account info
func (ac authClient) QueryAccount(accAddrStr string) (account exposed.Account, err error) {
	accAddr, err := types.AccAddressFromBech32(accAddrStr)
	if err != nil {
		return account, errors.New("failed. accAddress converted from Bech32 error")
	}

	res, err := ac.Query(accountInfoPath, addressStoreKey(accAddr))
	if err != nil {
		return account, fmt.Errorf("failed. ok client query error : %s", err.Error())
	}

	if res == nil {
		return account, errors.New("failed. your account has no record on the chain")
	}

	if err = ac.GetCodec().UnmarshalBinaryBare(res, &account); err != nil {
		return
	}

	return
}
