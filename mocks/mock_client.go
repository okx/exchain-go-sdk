package mocks

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	distribution "github.com/okex/okexchain-go-sdk/module/distribution/types"
	governance "github.com/okex/okexchain-go-sdk/module/governance/types"

	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/exposed"
	auth "github.com/okex/okexchain-go-sdk/module/auth/types"
	backend "github.com/okex/okexchain-go-sdk/module/backend/types"
	dex "github.com/okex/okexchain-go-sdk/module/dex/types"
	farm "github.com/okex/okexchain-go-sdk/module/farm/types"
	order "github.com/okex/okexchain-go-sdk/module/order/types"
	slashing "github.com/okex/okexchain-go-sdk/module/slashing/types"
	staking "github.com/okex/okexchain-go-sdk/module/staking/types"
	tendermint "github.com/okex/okexchain-go-sdk/module/tendermint/types"
	token "github.com/okex/okexchain-go-sdk/module/token/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"
	cmn "github.com/tendermint/tendermint/libs/common"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmstate "github.com/tendermint/tendermint/state"
	tmtypes "github.com/tendermint/tendermint/types"
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
func (mc *MockClient) Distribution() exposed.Distribution {
	return mc.modules[distribution.ModuleName].(exposed.Distribution)
}
func (mc *MockClient) Governance() exposed.Governance {
	return mc.modules[governance.ModuleName].(exposed.Governance)
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
func (mc *MockClient) Farm() exposed.Farm {
	return mc.modules[farm.ModuleName].(exposed.Farm)
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
func (mc *MockClient) GetRawValidatorsResultPointer(height, votingPower, proposerPriority int64,
	consPubkey crypto.PubKey) *ctypes.ResultValidators {
	return &ctypes.ResultValidators{
		BlockHeight: height,
		Validators: []*tmtypes.Validator{
			{
				PubKey:           consPubkey,
				VotingPower:      votingPower,
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

// BuildBackendDealsResultBytes generates the backend deals result bytes for test
func (mc *MockClient) BuildBackendDealsResultBytes(timestamp, height int64, orderID, sender, product, side, fee string, price, quantity float64) []byte {
	listResp := backend.ListResponse{
		Data: backend.ListDataRes{
			Data: []backend.Deal{
				{
					Timestamp:   timestamp,
					BlockHeight: height,
					OrderID:     orderID,
					Sender:      sender,
					Product:     product,
					Side:        side,
					Price:       price,
					Quantity:    quantity,
					Fee:         fee,
				},
			},
		},
	}

	bytes, err := json.Marshal(listResp)
	require.NoError(mc.t, err)
	return bytes
}

// BuildBackendOrdersResultBytes generates the backend orders result bytes for test
func (mc *MockClient) BuildBackendOrdersResultBytes(txHash, orderID, sender, product, side, price, quantity, filledAvgPrice,
	remainQuantity string, status, timestamp int64) []byte {
	listResp := backend.ListResponse{
		Data: backend.ListDataRes{
			Data: []backend.Order{
				{
					TxHash:         txHash,
					OrderID:        orderID,
					Sender:         sender,
					Product:        product,
					Side:           side,
					Price:          price,
					Quantity:       quantity,
					Status:         status,
					FilledAvgPrice: filledAvgPrice,
					RemainQuantity: remainQuantity,
					Timestamp:      timestamp,
				},
			},
		},
	}

	bytes, err := json.Marshal(listResp)
	require.NoError(mc.t, err)
	return bytes
}

// BuildBackendMatchResultBytes generates the backend match result bytes for test
func (mc *MockClient) BuildBackendMatchResultBytes(timestamp, height int64, product string, price, quantity float64) []byte {
	listResp := backend.ListResponse{
		Data: backend.ListDataRes{
			Data: []backend.MatchResult{
				{
					Timestamp:   timestamp,
					BlockHeight: height,
					Product:     product,
					Price:       price,
					Quantity:    quantity,
				},
			},
		},
	}

	bytes, err := json.Marshal(listResp)
	require.NoError(mc.t, err)
	return bytes
}

// BuildBackendMatchResultBytes generates the backend transactions result bytes for test
func (mc *MockClient) BuildBackendTransactionsResultBytes(txHash, accAddr, symbol, quantity, fee string, txType, side,
	timestamp int64) []byte {
	listResp := backend.ListResponse{
		Data: backend.ListDataRes{
			Data: []backend.Transaction{
				{
					TxHash:    txHash,
					Type:      txType,
					Address:   accAddr,
					Symbol:    symbol,
					Side:      side,
					Quantity:  quantity,
					Fee:       fee,
					Timestamp: timestamp,
				},
			},
		},
	}

	bytes, err := json.Marshal(listResp)
	require.NoError(mc.t, err)
	return bytes
}

// BuildBackendCandlesBytes generates the backend candles bytes for test
func (mc *MockClient) BuildBackendCandlesBytes(candles [][]string) []byte {
	baseResp := backend.BaseResponse{
		Data: candles,
	}

	bytes, err := json.Marshal(baseResp)
	require.NoError(mc.t, err)
	return bytes
}

// BuildBackendTickersBytes generates the backend tickers bytes for test
func (mc *MockClient) BuildBackendTickersBytes(symbol, product, timestamp, open, close, high, low, price, volumn,
	change string) []byte {
	baseResp := backend.BaseResponse{
		Data: []backend.Ticker{
			{
				Symbol:    symbol,
				Product:   product,
				Timestamp: timestamp,
				Open:      open,
				Close:     close,
				High:      high,
				Low:       low,
				Price:     price,
				Volume:    volumn,
				Change:    change,
			},
		},
	}

	bytes, err := json.Marshal(baseResp)
	require.NoError(mc.t, err)
	return bytes
}

// BuildProposalsBytes generates the proposals bytes for test
func (mc *MockClient) BuildProposalsBytes(proposalID uint64, status governance.ProposalStatus,
	mockTime time.Time, totalDeposit sdk.DecCoins, mockPower sdk.Dec) []byte {
	proposals := []governance.Proposal{
		{
			Content:         governance.TextProposal{},
			ProposalID:      proposalID,
			Status:          status,
			SubmitTime:      mockTime,
			DepositEndTime:  mockTime,
			VotingStartTime: mockTime,
			VotingEndTime:   mockTime,
			TotalDeposit:    totalDeposit,
			FinalTallyResult: governance.TallyResult{
				TotalPower:      mockPower,
				TotalVotedPower: mockPower,
				Yes:             mockPower,
				Abstain:         mockPower,
				No:              mockPower,
				NoWithVeto:      mockPower,
			},
		},
	}

	return mc.cdc.MustMarshalJSON(proposals)
}

// BuildFarmPoolBytes generates the farm pool bytes for test
func (mc *MockClient) BuildFarmPoolBytes(poolName, ownerAddrStr, tokenSymbol string, height int64, amountDec sdk.Dec) []byte {
	ownerAddr, err := sdk.AccAddressFromBech32(ownerAddrStr)
	require.NoError(mc.t, err)

	farmPool := farm.FarmPool{
		Owner:        ownerAddr,
		Name:         poolName,
		SymbolLocked: tokenSymbol,
		YieldedTokenInfos: farm.YieldedTokenInfos{
			{
				RemainingAmount:         sdk.NewDecCoinFromDec(tokenSymbol, amountDec),
				StartBlockHeightToYield: height,
				AmountYieldedPerBlock:   amountDec,
			},
		},
		DepositAmount:    sdk.NewDecCoinFromDec(tokenSymbol, amountDec),
		TotalValueLocked: sdk.NewDecCoinFromDec(tokenSymbol, amountDec),
		AmountYielded:    sdk.NewDecCoins(sdk.NewDecCoinFromDec(tokenSymbol, amountDec)),
	}

	return mc.cdc.MustMarshalJSON(farmPool)
}

// BuildFarmPoolNameList generates the farm pool name list bytes for test
func (mc *MockClient) BuildFarmPoolNameList(poolName ...string) []byte {
	return mc.cdc.MustMarshalJSON(poolName)
}

// BuildAccAddrList generates the account address list bytes for test
func (mc *MockClient) BuildAccAddrList(accAddr ...sdk.AccAddress) []byte {
	return mc.cdc.MustMarshalJSON(accAddr)
}
