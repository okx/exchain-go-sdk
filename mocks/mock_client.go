package mocks

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/exposed"
	auth "github.com/okex/okchain-go-sdk/module/auth/types"
	backend "github.com/okex/okchain-go-sdk/module/backend/types"
	dex "github.com/okex/okchain-go-sdk/module/dex/types"
	order "github.com/okex/okchain-go-sdk/module/order/types"
	slashing "github.com/okex/okchain-go-sdk/module/slashing/types"
	staking "github.com/okex/okchain-go-sdk/module/staking/types"
	tendermint "github.com/okex/okchain-go-sdk/module/tendermint/types"
	token "github.com/okex/okchain-go-sdk/module/token/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"
	cmn "github.com/tendermint/tendermint/libs/common"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmstate "github.com/tendermint/tendermint/state"
	tmtypes "github.com/tendermint/tendermint/types"

	"testing"
	"time"
)

// MockClient - structure of the mock client for gosdk testing
type MockClient struct {
	t *testing.T
	*sdk.MockBaseClient
	config  sdk.ClientConfig
	cdc     sdk.SDKCodec
	modules map[string]sdk.Module
}

// NewMockClient creates a new instance of MockClient
func NewMockClient(t *testing.T, ctrl *gomock.Controller, config sdk.ClientConfig) MockClient {
	cdc := sdk.NewCodec()
	pMockClient := &MockClient{
		t:              t,
		MockBaseClient: sdk.NewMockBaseClient(ctrl),
		config:         config,
		cdc:            cdc,
		modules:        make(map[string]sdk.Module),
	}

	return *pMockClient
}

// RegisterModule registers the specific module for MockClient
func (mc *MockClient) RegisterModule(mods ...sdk.Module) {
	for _, mod := range mods {
		moduleName := mod.Name()
		if _, ok := mc.modules[moduleName]; ok {
			panic(fmt.Sprintf("duplicated module: %s", moduleName))
		}
		// register codec by each module
		mod.RegisterCodec(mc.cdc)
		mc.modules[moduleName] = mod
	}
	sdk.RegisterBasicCodec(mc.cdc)
	mc.cdc.Seal()
}

// GetConfig returns the client config
func (mc *MockClient) GetConfig() sdk.ClientConfig {
	return mc.config
}

// GetCodec returns the client codec
func (mc *MockClient) GetCodec() sdk.SDKCodec {
	return mc.cdc
}

// nolint
func (mc *MockClient) Auth() exposed.Auth {
	return mc.modules[auth.ModuleName].(exposed.Auth)
}
func (mc *MockClient) Backend() exposed.Backend {
	return mc.modules[backend.ModuleName].(exposed.Backend)
}
func (mc *MockClient) Dex() exposed.Dex {
	return mc.modules[dex.ModuleName].(exposed.Dex)
}
func (mc *MockClient) Order() exposed.Order {
	return mc.modules[order.ModuleName].(exposed.Order)
}
func (mc *MockClient) Staking() exposed.Staking {
	return mc.modules[staking.ModuleName].(exposed.Staking)
}
func (mc *MockClient) Slashing() exposed.Slashing {
	return mc.modules[slashing.ModuleName].(exposed.Slashing)
}
func (mc *MockClient) Token() exposed.Token {
	return mc.modules[token.ModuleName].(exposed.Token)
}
func (mc *MockClient) Tendermint() exposed.Tendermint {
	return mc.modules[tendermint.ModuleName].(exposed.Tendermint)
}

// BuildAccountBytes generates the account bytes for test
func (mc *MockClient) BuildAccountBytes(accAddrStr, accPubkeyStr, coinsStr string, accNum, seqNum uint64) []byte {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	require.NoError(mc.t, err)
	accPubkey, err := sdk.GetAccPubKeyBech32(accPubkeyStr)
	require.NoError(mc.t, err)
	coins, err := sdk.ParseDecCoins(coinsStr)
	require.NoError(mc.t, err)
	account := auth.BaseAccount{
		Address:       accAddr,
		Coins:         coins,
		PubKey:        accPubkey,
		AccountNumber: accNum,
		Sequence:      seqNum,
	}

	bytes, err := mc.cdc.MarshalBinaryBare(account)
	require.NoError(mc.t, err)

	return bytes
}

// BuildTokenPairsBytes generates the token pairs bytes for test
func (mc *MockClient) BuildTokenPairsBytes(baseAssetSymbol1, baseAssetSymbol2, quoteAssetSymbol string, initPrice,
	minQuantity sdk.Dec, maxPriceDigit, maxQuantityDigit, blockHeight1, blockHeight2 int64, ID1, ID2 uint64, delisting bool,
	owner sdk.AccAddress, deposits sdk.DecCoin) []byte {

	var tokenPairs []dex.TokenPair

	tokenPairs = append(tokenPairs, dex.TokenPair{
		BaseAssetSymbol:  baseAssetSymbol1,
		QuoteAssetSymbol: quoteAssetSymbol,
		InitPrice:        initPrice,
		MaxPriceDigit:    maxPriceDigit,
		MaxQuantityDigit: maxQuantityDigit,
		MinQuantity:      minQuantity,
		ID:               ID1,
		Delisting:        delisting,
		Owner:            owner,
		Deposits:         deposits,
		BlockHeight:      blockHeight1,
	})

	tokenPairs = append(tokenPairs, dex.TokenPair{
		BaseAssetSymbol:  baseAssetSymbol2,
		QuoteAssetSymbol: quoteAssetSymbol,
		InitPrice:        initPrice,
		MaxPriceDigit:    maxPriceDigit,
		MaxQuantityDigit: maxQuantityDigit,
		MinQuantity:      minQuantity,
		ID:               ID2,
		Delisting:        delisting,
		Owner:            owner,
		Deposits:         deposits,
		BlockHeight:      blockHeight2,
	})

	return mc.cdc.MustMarshalJSON(tokenPairs)
}

// BuildOrderDetailBytes generates the order detail bytes for test
func (mc *MockClient) BuildOrderDetailBytes(txHash, orderID, extraInfo, product, side string, status, timestamp,
	orderExpireBlocks int64, sender sdk.AccAddress, price, quantity, filledAvgPrice, remainQuantity, remainLocked sdk.Dec,
	feePerBlock sdk.DecCoin) []byte {
	orderDetail := order.OrderDetail{
		TxHash:            txHash,
		OrderID:           orderID,
		Sender:            sender,
		Product:           product,
		Side:              side,
		Price:             price,
		Quantity:          quantity,
		Status:            status,
		FilledAvgPrice:    filledAvgPrice,
		RemainQuantity:    remainQuantity,
		RemainLocked:      remainLocked,
		Timestamp:         timestamp,
		OrderExpireBlocks: orderExpireBlocks,
		FeePerBlock:       feePerBlock,
		ExtraInfo:         extraInfo,
	}

	return mc.cdc.MustMarshalJSON(orderDetail)
}

// BuildBookResBytes generates the book result bytes for test
func (mc *MockClient) BuildBookResBytes(askPrice, askQuantity, bidPrice, bidQuantity string) []byte {
	var bookRes order.BookRes
	bookRes.Asks = append(bookRes.Asks, order.BookResItem{
		Price:    askPrice,
		Quantity: askQuantity,
	})

	bookRes.Bids = append(bookRes.Bids, order.BookResItem{
		Price:    bidPrice,
		Quantity: bidQuantity,
	})

	return mc.cdc.MustMarshalJSON(bookRes)
}

// BuildAccountTokensInfoBytes generates the account tokens info bytes for test
func (mc *MockClient) BuildAccountTokensInfoBytes(addrStr, symbol, available, freeze, locked string) []byte {
	accTokensInfo := token.AccountTokensInfo{
		Address: addrStr,
	}

	accTokensInfo.Currencies = append(accTokensInfo.Currencies, token.CoinInfo{
		Symbol:    symbol,
		Available: available,
		Freeze:    freeze,
		Locked:    locked,
	})

	return mc.cdc.MustMarshalJSON(accTokensInfo)
}

// BuildTokenInfoBytes generates the token info bytes for test
func (mc *MockClient) BuildTokenInfoBytes(description, symbol, originalSymbol, wholeName string, originalTotalSupply,
	totalSupply sdk.Dec, owner sdk.AccAddress, mintable, isSlice bool) []byte {
	tokenInfo := token.Token{
		Description:         description,
		Symbol:              symbol,
		OriginalSymbol:      originalSymbol,
		WholeName:           wholeName,
		OriginalTotalSupply: originalTotalSupply,
		TotalSupply:         totalSupply,
		Owner:               owner,
		Mintable:            mintable,
	}

	if isSlice {
		return mc.cdc.MustMarshalJSON([]token.Token{tokenInfo})
	}

	return mc.cdc.MustMarshalJSON(tokenInfo)
}

// BuildValidatorsBytes generates the validator bytes for test
func (mc *MockClient) BuildValidatorBytes(valAddr sdk.ValAddress, consPubKey, moniker, identity, website, details string,
	status byte, delegatorShares, minSelfDelegation sdk.Dec, unbondingHeight int64, unbondingCompletionTime time.Time,
	jailed bool) []byte {
	consPK, err := sdk.GetConsPubKeyBech32(consPubKey)
	require.NoError(mc.t, err)
	val := staking.ValidatorInner{
		OperatorAddress: valAddr,
		ConsPubKey:      consPK,
		Jailed:          jailed,
		Status:          status,
		DelegatorShares: delegatorShares,
		Description: staking.Description{
			Moniker:  moniker,
			Identity: identity,
			Website:  website,
			Details:  details,
		},
		UnbondingHeight:         unbondingHeight,
		UnbondingCompletionTime: unbondingCompletionTime,
		MinSelfDelegation:       minSelfDelegation,
	}

	return mc.cdc.MustMarshalBinaryLengthPrefixed(val)

}

// BuildDelegatorBytes generates the delegator bytes for test
func (mc *MockClient) BuildDelegatorBytes(delAddr, proxyAddr sdk.AccAddress, valAddrs []sdk.ValAddress, shares, tokens,
	totalDelegatedTokens sdk.Dec, isProxy bool) []byte {
	delegator := staking.Delegator{
		DelegatorAddress:     delAddr,
		ValidatorAddresses:   valAddrs,
		Shares:               shares,
		Tokens:               tokens,
		IsProxy:              isProxy,
		TotalDelegatedTokens: totalDelegatedTokens,
		ProxyAddress:         proxyAddr,
	}

	return mc.cdc.MustMarshalBinaryLengthPrefixed(delegator)
}

// BuildUndelegationBytes generates the undelegation bytes for test
func (mc *MockClient) BuildUndelegationBytes(delAddr sdk.AccAddress, quantity sdk.Dec, completionTime time.Time) []byte {
	undelegation := staking.Undelegation{
		DelegatorAddress: delAddr,
		Quantity:         quantity,
		CompletionTime:   completionTime,
	}

	return mc.cdc.MustMarshalJSON(undelegation)
}

// GetRawResultBlockPointer generates the raw tendermint block result pointer for test
func (mc *MockClient) GetRawResultBlockPointer(chainID string, height int64, time time.Time, appHash,
	blockIDHash cmn.HexBytes) *ctypes.ResultBlock {
	return &ctypes.ResultBlock{
		Block: &tmtypes.Block{
			Header: tmtypes.Header{
				ChainID: chainID,
				Height:  height,
				Time:    time,
				AppHash: appHash,
			},
			Evidence: tmtypes.EvidenceData{},
			LastCommit: &tmtypes.Commit{
				BlockID: tmtypes.BlockID{
					Hash: blockIDHash,
				},
			},
		},
	}
}

// GetRawResultBlockResultsPointer generates the raw tendermint result block results pointer for test
func (mc *MockClient) GetRawResultBlockResultsPointer(power, height int64, pkType, eventType string,
	kvPairKey []byte) *ctypes.ResultBlockResults {
	return &ctypes.ResultBlockResults{
		Height: height,
		Results: &tmstate.ABCIResponses{
			BeginBlock: &abci.ResponseBeginBlock{
				Events: []abci.Event{
					{
						Type: eventType,
						Attributes: []common.KVPair{
							{
								Key: kvPairKey,
							},
						},
					},
				},
			},
			EndBlock: &abci.ResponseEndBlock{
				ValidatorUpdates: []abci.ValidatorUpdate{
					{
						PubKey: abci.PubKey{
							Type: pkType,
						},
						Power: power,
					},
				},
			},
		},
	}
}

// GetRawCommitResultPointer generates the raw tendermint commit result pointer for test
func (mc *MockClient) GetRawCommitResultPointer(canonicalCommit bool, chainID string, height int64, time time.Time, appHash,
	blockIDHash cmn.HexBytes) *ctypes.ResultCommit {
	return &ctypes.ResultCommit{
		CanonicalCommit: canonicalCommit,
		SignedHeader: tmtypes.SignedHeader{
			Header: &tmtypes.Header{
				ChainID: chainID,
				Height:  height,
				Time:    time,
				AppHash: appHash,
			},
			Commit: &tmtypes.Commit{
				BlockID: tmtypes.BlockID{
					Hash: blockIDHash,
				},
			},
		},
	}
}

// GetRawValidatorsResultPointer generates the raw tendermint validators result pointer for test
func (mc *MockClient) GetRawValidatorsResultPointer(height, VotingPower, proposerPriority int64,
	consPubkey crypto.PubKey) *ctypes.ResultValidators {
	return &ctypes.ResultValidators{
		BlockHeight: height,
		Validators: []*tmtypes.Validator{
			{
				PubKey:           consPubkey,
				VotingPower:      VotingPower,
				ProposerPriority: proposerPriority,
			},
		},
	}
}

// GetRawTxResultPointer generates the raw tendermint tx result pointer for test
func (mc *MockClient) GetRawTxResultPointer(hash cmn.HexBytes, height int64, code uint32, log, eventType string,
	tx []byte) *ctypes.ResultTx {
	return &ctypes.ResultTx{
		Hash:   hash,
		Height: height,
		Tx:     tx,
		TxResult: abci.ResponseDeliverTx{
			Code: code,
			Log:  log,
			Events: []abci.Event{
				{
					Type: eventType,
				},
			},
		},
	}
}

// GetRawTxResultPointer generates the raw tendermint tx search result pointer for test
func (mc *MockClient) GetRawResultTxSearchPointer(totalCount int, hash cmn.HexBytes, height int64, code uint32, log,
	eventType string, tx []byte) *ctypes.ResultTxSearch {
	return &ctypes.ResultTxSearch{
		TotalCount: totalCount,
		Txs: []*ctypes.ResultTx{
			mc.GetRawTxResultPointer(hash, height, code, log, eventType, tx),
		},
	}
}
