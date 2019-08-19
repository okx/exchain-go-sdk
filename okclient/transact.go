package okclient

import (
	"fmt"
	"github.com/ok-chain/ok-gosdk/common/libs/pkg/errors"
	"github.com/ok-chain/ok-gosdk/common/transactParams"
	"github.com/ok-chain/ok-gosdk/crypto/keys"
	"github.com/ok-chain/ok-gosdk/types"
	"github.com/ok-chain/ok-gosdk/types/msg"
	"github.com/ok-chain/ok-gosdk/types/tx"
	"github.com/ok-chain/ok-gosdk/utils"
)

// broadcast mode
const (
	BroadcastBlock = "block"
	BroadcastSync  = "sync"
	BroadcastAsync = "async"
)

func (okCli *OKClient) Send(fromInfo keys.Info, passWd, toAddr, coinsStr, memo string, accNum, seqNum uint64) (resp types.TxResponse, err error) {
	if !transactParams.IsValidSendParams(fromInfo, passWd, toAddr) {
		return types.TxResponse{}, errors.New("err : params input to send are invalid")
	}

	to, err := types.AccAddressFromBech32(toAddr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Address [%s] error: %s", toAddr, err)
	}

	coins, err := utils.ParseCoins(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := msg.NewMsgTokenSend(fromInfo.GetAddress(), to, coins)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return okCli.broadcast(stdBytes, BroadcastBlock)
}

func (okCli *OKClient) NewOrder(fromInfo keys.Info, passWd, product, side, price, quantity, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if !transactParams.IsValidNewOrderParams(fromInfo, passWd, product, side, price, quantity, memo, ) {
		return types.TxResponse{}, errors.New("err : params input to pend a order are invalid")
	}
	msg := msg.NewMsgNewOrder(fromInfo.GetAddress(), product, side, price, quantity)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return okCli.broadcast(stdBytes, BroadcastBlock)

}
