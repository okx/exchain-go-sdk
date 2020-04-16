package auth

import (
	"errors"
	"github.com/okex/okchain-go-sdk/module/auth/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
)

// QueryAccount gets the account info
func (ac authClient) QueryAccount(accAddrStr string) (account types.Account, err error) {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	if err != nil {
		return account, errors.New("failed. accAddress converted from Bech32 error")
	}

	res, err := ac.Query(types.AccountInfoPath, types.AddressStoreKey(accAddr))
	if err != nil {
		return account, utils.ErrClientQuery(err.Error())
	}

	if res == nil {
		return account, errors.New("failed. your account has no record on the chain")
	}

	if err = ac.GetCodec().UnmarshalBinaryBare(res, &account); err != nil {
		return
	}

	return
}
