package okclient

import (
	"github.com/ok-chain/ok-gosdk/common/transactParams"
	"github.com/ok-chain/ok-gosdk/utils"

	"testing"
)

const (
	name     = "alice"
	passWd   = "12345678"
	mnemonic = "sustain hole urban away boy core lazy brick wait drive tiger tell"
	addr1    = "okchain1dycww54mz20sfakx7hqtkf2ghdlx6tjry977gy"
	addr2    = "okchain1n0njw83czuk2c8v03fh64jd2u3sxqhdhckmdy6"
	rpcUrl   = "localhost:26657"
)

func TestSend(t *testing.T) {
	okCli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	//fmt.Println(fromInfo.GetAddress().String())
	params := transactParams.NewTransferParams(fromInfo, passWd, addr1, "10.24okb", "I love okb")
	okCli.Send(params)
}
