package okclient

import (
	"fmt"
	//"fmt"
	"github.com/ok-chain/ok-gosdk/utils"

	"testing"
)

const (
	name     = "alice"
	passWd   = "12345678"
	mnemonic = "sustain hole urban away boy core lazy brick wait drive tiger tell"
	addr1    = "okchain1dycww54mz20sfakx7hqtkf2ghdlx6tjry977gy"
	addr2    = "okchain1n0njw83czuk2c8v03fh64jd2u3sxqhdhckmdy6"
)

func TestSend(t *testing.T) {
	okCli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := okCli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := okCli.Send(fromInfo, passWd, addr1, "10.24okb", "I love OK", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestNewOrder(t *testing.T) {
	okCli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := okCli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := okCli.NewOrder(fromInfo, passWd, "xxb_okb", "BUY", "11.1", "1.23", "I love OK", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
	fmt.Println("orderId:", res.Tags[1].Value)
}

func TestCancelOrder(t *testing.T) {
	okCli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := okCli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := okCli.CancelOrder(fromInfo, passWd, "ID0000177104-1", "I love OK", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}
