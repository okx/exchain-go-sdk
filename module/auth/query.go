package auth

import (
	"errors"
	"fmt"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth"
	authtypes "github.com/okex/exchain/libs/cosmos-sdk/x/auth/types"
	"github.com/okex/exchain-go-sdk/module/auth/types"
	"github.com/okex/exchain-go-sdk/utils"
)

// QueryAccount gets the account info
func (ac authClient) QueryAccount(accAddrStr string) (account types.Account, err error) {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	if err != nil {
		return account, errors.New("failed. accAddress converted from Bech32 error")
	}

	path := fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount)
	bytes, err := ac.GetCodec().MarshalJSON(authtypes.NewQueryAccountParams(accAddr))
	if err != nil {
		return account, utils.ErrClientQuery(err.Error())
	}

	res, _, err := ac.Query(path, bytes)
	if res == nil {
		return account, errors.New("failed. your account has no record on the chain")
	}

	if err = ac.GetCodec().UnmarshalJSON(res, &account); err != nil {
		return account, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
