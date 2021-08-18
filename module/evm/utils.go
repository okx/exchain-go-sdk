package evm

import (
	"fmt"
	"os"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcore "github.com/ethereum/go-ethereum/core/types"
	"github.com/okex/exchain-go-sdk/utils"
	apptypes "github.com/okex/exchain/app/types"
	evmtypes "github.com/okex/exchain/x/evm/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// GenerateUnsignedEvmTx
func (ec evmClient) GenerateUnsignedEvmTx(targetPath, fromAddrHex, toAddrHex, amountStr, payloadStr, memo string,
	nonce uint64) error {
	fromAddr, err := utils.ToCosmosAddress(fromAddrHex)
	if err != nil {
		return err
	}

	toAddr, err := utils.ToCosmosAddress(toAddrHex)
	if err != nil {
		return err
	}

	amount := sdk.ZeroDec()
	if len(amountStr) != 0 {
		amount, err = sdk.NewDecFromStr(amountStr)
		if err != nil {
			return err
		}
	}

	var data []byte
	if len(payloadStr) != 0 {
		if !strings.HasPrefix(payloadStr, "0x") {
			payloadStr = fmt.Sprintf("0x%s", payloadStr)
		}

		data, err = hexutil.Decode(payloadStr)
		if err != nil {
			return err
		}
	}

	config := ec.GetConfig()
	msg := evmtypes.NewMsgEthermint(
		nonce,
		&toAddr,
		sdk.NewIntFromBigInt(amount.Int),
		config.Gas,
		sdk.NewInt(apptypes.DefaultGasPrice),
		data,
		fromAddr,
	)

	unsignedStdTx := authtypes.NewStdTx(
		[]sdk.Msg{msg},
		authtypes.NewStdFee(config.Gas, config.Fees),
		nil,
		memo,
	)

	jsonBytes, err := ec.GetCodec().MarshalJSONIndent(unsignedStdTx, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))
	f, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(jsonBytes)
	return err
}

// GetTxHash calculates the tx hash
func (ec evmClient) GetTxHash(signedTx *ethcore.Transaction) (txHash ethcmn.Hash, err error) {
	v, r, s := signedTx.RawSignatureValues()
	tx := evmtypes.MsgEthereumTx{
		Data: evmtypes.TxData{
			AccountNonce: signedTx.Nonce(),
			Price:        signedTx.GasPrice(),
			GasLimit:     signedTx.Gas(),
			Recipient:    signedTx.To(),
			Amount:       signedTx.Value(),
			Payload:      signedTx.Data(),
			V:            v,
			R:            r,
			S:            s,
		},
	}

	txBytes, err := authcli.GetTxEncoder(ec.GetCodec())(tx)
	if err != nil {
		return
	}

	txHash = ethcmn.BytesToHash(tmhash.Sum(txBytes))
	return
}
