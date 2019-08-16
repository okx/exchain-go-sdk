package transactParams

import (
	"errors"
	"github.com/ok-chain/ok-gosdk/crypto/keys"
	"strings"
)

type TransferParams struct {
	fromInfo keys.Info
	PassWd   string
	ToAddr   string
	CoinsStr string
	Memo     string
}

func NewTransferParams(Info keys.Info, passWd, toAddr, coinsStr, memo string) TransferParams {
	return TransferParams{
		Info,
		passWd,
		toAddr,
		coinsStr,
		memo,
	}
}

func (tp *TransferParams) IsValid() (bool, error) {
	if tp.fromInfo == nil {
		return false, errors.New("input invalid name")
	}
	if len(tp.PassWd) == 0 {
		return false, errors.New("no password input")
	}
	if len(tp.ToAddr) != 46 || strings.HasPrefix(tp.ToAddr, "okchain") {
		return false, errors.New("input invalid receiver address")
	}
	return true, nil
}
