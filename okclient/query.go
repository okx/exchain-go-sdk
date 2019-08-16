package okclient

import (
	"errors"
	"fmt"
	"github.com/ok-chain/ok-gosdk/types"
	"github.com/ok-chain/ok-gosdk/utils"
)

const (
	accountPath = "/store/acc/key"
)

func (okCli *OKClient) GetAccountInfoByAddr(addr string) (types.Account, error) {
	accAddr, err := types.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errors.New("err : AccAddress converted from Bech32 Failed")
	}

	res, err := okCli.query(accountPath, utils.AddressStoreKey(accAddr))
	if err != nil {
		return nil, fmt.Errorf("ok cliemt query error : %s", err.Error())
	}

	var account types.Account
	if err = okCli.cdc.UnmarshalBinaryBare(res, &account); err != nil {
		return nil, fmt.Errorf("err : %s", err.Error())
	}

	return account, nil
}
