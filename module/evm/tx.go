package evm

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/okex/okexchain-go-sdk/utils"
	apptypes "github.com/okex/okexchain/app/types"
	evmtypes "github.com/okex/okexchain/x/evm/types"
	"strconv"
	"strings"
)

func (ec evmClient) SendTx(fromInfo keys.Info, passWd, toAddrStr, amountStr, payloadStr, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	toAddr, err := utils.ToCosmosAddress(toAddrStr)
	if err != nil {
		return
	}

	amount, err := strconv.ParseInt(amountStr, 0, 64)
	if err != nil {
		return
	}

	var data []byte
	if len(payloadStr) != 0 {
		if !strings.HasPrefix(payloadStr, "0x") {
			payloadStr = fmt.Sprintf("0x%s", payloadStr)
		}

		data, err = hexutil.Decode(payloadStr)
		if err != nil {
			return
		}
	}

	msg := evmtypes.NewMsgEthermint(
		seqNum,
		&toAddr,
		sdk.NewInt(amount),
		ec.GetConfig().Gas,
		sdk.NewInt(apptypes.DefaultGasPrice),
		data,
		fromInfo.GetAddress(),
	)
	return ec.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}
