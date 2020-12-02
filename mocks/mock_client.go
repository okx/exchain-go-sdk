package mocks

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/exposed"
	auth "github.com/okex/okexchain-go-sdk/module/auth/types"
	backend "github.com/okex/okexchain-go-sdk/module/backend/types"
	dex "github.com/okex/okexchain-go-sdk/module/dex/types"
	distribution "github.com/okex/okexchain-go-sdk/module/distribution/types"
	farm "github.com/okex/okexchain-go-sdk/module/farm/types"
	governance "github.com/okex/okexchain-go-sdk/module/governance/types"
	order "github.com/okex/okexchain-go-sdk/module/order/types"
	slashing "github.com/okex/okexchain-go-sdk/module/slashing/types"
	staking "github.com/okex/okexchain-go-sdk/module/staking/types"
	tendermint "github.com/okex/okexchain-go-sdk/module/tendermint/types"
	token "github.com/okex/okexchain-go-sdk/module/token/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	evmtypes "github.com/okex/okexchain/app/types"
	orderkeeper "github.com/okex/okexchain/x/order/keeper"
	stakingtypes "github.com/okex/okexchain/x/staking/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"testing"
	"time"
)

// MockClient - structure of the mock client for gosdk testing
type MockClient struct {
	t *testing.T
	*gosdktypes.MockBaseClient
	config  gosdktypes.ClientConfig
	cdc     *codec.Codec
	modules map[string]gosdktypes.Module
}

// NewMockClient creates a new instance of MockClient
func NewMockClient(t *testing.T, ctrl *gomock.Controller, config gosdktypes.ClientConfig) MockClient {
	cdc := gosdktypes.NewCodec()
	pMockClient := &MockClient{
		t:              t,
		MockBaseClient: gosdktypes.NewMockBaseClient(ctrl),
		config:         config,
		cdc:            cdc,
		modules:        make(map[string]gosdktypes.Module),
	}

	return *pMockClient
}

// RegisterModule registers the specific module for MockClient
func (mc *MockClient) RegisterModule(mods ...gosdktypes.Module) {
	for _, mod := range mods {
		moduleName := mod.Name()
		if _, ok := mc.modules[moduleName]; ok {
			panic(fmt.Sprintf("duplicated module: %s", moduleName))
		}
		// register codec by each module
		mod.RegisterCodec(mc.cdc)
		mc.modules[moduleName] = mod
	}
	gosdktypes.RegisterBasicCodec(mc.cdc)
	mc.cdc.Seal()
}

// GetConfig returns the client config
func (mc *MockClient) GetConfig() gosdktypes.ClientConfig {
	return mc.config
}

// GetCodec returns the client codec
func (mc *MockClient) GetCodec() gosdktypes.SDKCodec {
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
func (mc *MockClient) Farm() exposed.Farm {
	return mc.modules[farm.ModuleName].(exposed.Farm)
}
func (mc *MockClient) Governance() exposed.Governance {
	return mc.modules[governance.ModuleName].(exposed.Governance)
}
func (mc *MockClient) Order() exposed.Order {
	return mc.modules[order.ModuleName].(exposed.Order)
}
func (mc *MockClient) Slashing() exposed.Slashing {
	return mc.modules[slashing.ModuleName].(exposed.Slashing)
}
func (mc *MockClient) Staking() exposed.Staking {
	return mc.modules[staking.ModuleName].(exposed.Staking)
}
func (mc *MockClient) Tendermint() exposed.Tendermint {
	return mc.modules[tendermint.ModuleName].(exposed.Tendermint)
}
func (mc *MockClient) Token() exposed.Token {
	return mc.modules[token.ModuleName].(exposed.Token)
}

// BuildAccountBytes generates the account bytes for test
func (mc *MockClient) BuildAccountBytes(accAddrStr, accPubkeyStr, codeHash, coinsStr string, accNum, seqNum uint64) []byte {
	accAddr, err := sdk.AccAddressFromBech32(accAddrStr)
	require.NoError(mc.t, err)
	accPubkey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, accPubkeyStr)
	require.NoError(mc.t, err)
	coins, err := sdk.ParseDecCoins(coinsStr)
	require.NoError(mc.t, err)

	account := evmtypes.EthAccount{
		BaseAccount: &authtypes.BaseAccount{
			Address:       accAddr,
			Coins:         coins,
			PubKey:        accPubkey,
			AccountNumber: accNum,
			Sequence:      seqNum,
		},
		CodeHash: []byte(codeHash),
	}

	bytes, err := mc.cdc.MarshalJSON(account)
	require.NoError(mc.t, err)

	return bytes
}

// BuildTokenPairsResponseBytes generates the response of token pairs bytes for test
func (mc *MockClient) BuildTokenPairsResponseBytes(baseAssetSymbol1, baseAssetSymbol2, quoteAssetSymbol string, initPrice,
	minQuantity sdk.Dec, maxPriceDigit, maxQuantityDigit, blockHeight1, blockHeight2 int64, ID1, ID2 uint64, delisting bool,
	owner sdk.AccAddress, deposits sdk.DecCoin) []byte {
	tokenPairs := []dex.TokenPair{
		{
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
		},
		{
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
		},
	}

	response := dex.ListResponse{
		Data: dex.ListDataRes{
			Data: tokenPairs,
		},
	}

	res, err := json.Marshal(response)
	require.NoError(mc.t, err)
	return res
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
	var bookRes orderkeeper.BookRes
	bookRes.Asks = append(bookRes.Asks, orderkeeper.BookResItem{
		Price:    askPrice,
		Quantity: askQuantity,
	})

	bookRes.Bids = append(bookRes.Bids, orderkeeper.BookResItem{
		Price:    bidPrice,
		Quantity: bidQuantity,
	})

	return mc.cdc.MustMarshalJSON(bookRes)
}

// BuildTokenInfoBytes generates the token info bytes for test
func (mc *MockClient) BuildTokenInfoBytes(description, symbol, originalSymbol, wholeName string, originalTotalSupply,
	totalSupply sdk.Dec, owner sdk.AccAddress, mintable, isSlice bool, tokenType int) []byte {
	tokenInfo := token.TokenResp{
		Description:         description,
		Symbol:              symbol,
		OriginalSymbol:      originalSymbol,
		WholeName:           wholeName,
		OriginalTotalSupply: originalTotalSupply,
		Type:                tokenType,
		Owner:               owner,
		Mintable:            mintable,
		TotalSupply:         totalSupply,
	}

	if isSlice {
		return mc.cdc.MustMarshalJSON([]token.TokenResp{tokenInfo})
	}

	return mc.cdc.MustMarshalJSON(tokenInfo)
}

// BuildValidatorsBytes generates the validator bytes for test
func (mc *MockClient) BuildValidatorsBytes(valAddr sdk.ValAddress, consPubKey, moniker, identity, website, details string,
	status byte, delegatorShares, minSelfDelegation sdk.Dec, unbondingHeight int64, unbondingCompletionTime time.Time,
	jailed, isSlice bool) []byte {
	consPK, err := stakingtypes.GetConsPubKeyBech32(consPubKey)
	require.NoError(mc.t, err)
	val := stakingtypes.Validator{
		OperatorAddress: valAddr,
		ConsPubKey:      consPK,
		Jailed:          jailed,
		Status:          sdk.BondStatus(status),
		DelegatorShares: delegatorShares,
		Description: stakingtypes.Description{
			Moniker:  moniker,
			Identity: identity,
			Website:  website,
			Details:  details,
		},
		UnbondingHeight:         unbondingHeight,
		UnbondingCompletionTime: unbondingCompletionTime,
		MinSelfDelegation:       minSelfDelegation,
	}

	if isSlice {
		return mc.cdc.MustMarshalJSON([]stakingtypes.Validator{val})
	}

	return mc.cdc.MustMarshalJSON(val)
}

// BuildDelegatorBytes generates the delegator bytes for test
func (mc *MockClient) BuildDelegatorBytes(delAddr, proxyAddr sdk.AccAddress, valAddrs []sdk.ValAddress, shares, tokens,
	totalDelegatedTokens sdk.Dec, isProxy bool) []byte {
	delegator := stakingtypes.Delegator{
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
	undelegation := stakingtypes.UndelegationInfo{
		DelegatorAddress: delAddr,
		Quantity:         quantity,
		CompletionTime:   completionTime,
	}

	return mc.cdc.MustMarshalJSON(undelegation)
}

// GetRawResultBlockPointer generates the raw tendermint block result pointer for test
func (mc *MockClient) GetRawResultBlockPointer(chainID string, height int64, time time.Time, appHash,
	blockIDHash tmbytes.HexBytes) *ctypes.ResultBlock {
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
		// TODO
		//Results: &tmstate.ABCIResponses{
		//	BeginBlock: &abci.ResponseBeginBlock{
		//		Events: []abci.Event{
		//			{
		//				Type: eventType,
		//				Attributes: []common.KVPair{
		//					{
		//						Key: kvPairKey,
		//					},
		//				},
		//			},
		//		},
		//	},
		//	EndBlock: &abci.ResponseEndBlock{
		//		ValidatorUpdates: []abci.ValidatorUpdate{
		//			{
		//				PubKey: abci.PubKey{
		//					Type: pkType,
		//				},
		//				Power: power,
		//			},
		//		},
		//	},
		//},
	}
}

// GetRawCommitResultPointer generates the raw tendermint commit result pointer for test
func (mc *MockClient) GetRawCommitResultPointer(canonicalCommit bool, chainID string, height int64, time time.Time, appHash,
	blockIDHash tmbytes.HexBytes) *ctypes.ResultCommit {
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
func (mc *MockClient) GetRawTxResultPointer(hash tmbytes.HexBytes, height int64, code uint32, log, eventType string,
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
func (mc *MockClient) GetRawResultTxSearchPointer(totalCount int, hash tmbytes.HexBytes, height int64, code uint32, log,
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

// BuildFarmPoolsBytes generates the farm pools bytes for test
func (mc *MockClient) BuildFarmPoolsBytes(poolName1, poolName2, ownerAddrStr, tokenSymbol string, height int64, amountDec sdk.Dec) []byte {
	ownerAddr, err := sdk.AccAddressFromBech32(ownerAddrStr)
	require.NoError(mc.t, err)

	testDecCoin := sdk.NewDecCoinFromDec(tokenSymbol, amountDec)
	farmPools := []farm.FarmPool{
		{
			Owner:            ownerAddr,
			Name:             poolName1,
			MinLockAmount:    testDecCoin,
			DepositAmount:    testDecCoin,
			TotalValueLocked: testDecCoin,
			YieldedTokenInfos: farm.YieldedTokenInfos{
				{
					RemainingAmount:         sdk.NewDecCoinFromDec(tokenSymbol, amountDec),
					StartBlockHeightToYield: height,
					AmountYieldedPerBlock:   amountDec,
				},
			},
			TotalAccumulatedRewards: sdk.SysCoins{testDecCoin},
		},
		{
			Owner:            ownerAddr,
			Name:             poolName2,
			MinLockAmount:    testDecCoin,
			DepositAmount:    testDecCoin,
			TotalValueLocked: testDecCoin,
			YieldedTokenInfos: farm.YieldedTokenInfos{
				{
					RemainingAmount:         sdk.NewDecCoinFromDec(tokenSymbol, amountDec),
					StartBlockHeightToYield: height,
					AmountYieldedPerBlock:   amountDec,
				},
			},
			TotalAccumulatedRewards: sdk.SysCoins{testDecCoin},
		},
	}

	return mc.cdc.MustMarshalJSON(farmPools)
}

// BuildFarmPoolBytes generates the farm pool bytes for test
func (mc *MockClient) BuildFarmPoolBytes(poolName, ownerAddrStr, tokenSymbol string, height int64, amountDec sdk.Dec) []byte {
	ownerAddr, err := sdk.AccAddressFromBech32(ownerAddrStr)
	require.NoError(mc.t, err)

	testDecCoin := sdk.NewDecCoinFromDec(tokenSymbol, amountDec)
	farmPool := farm.FarmPool{
		Owner:            ownerAddr,
		Name:             poolName,
		MinLockAmount:    testDecCoin,
		DepositAmount:    testDecCoin,
		TotalValueLocked: testDecCoin,
		YieldedTokenInfos: farm.YieldedTokenInfos{
			{
				RemainingAmount:         sdk.NewDecCoinFromDec(tokenSymbol, amountDec),
				StartBlockHeightToYield: height,
				AmountYieldedPerBlock:   amountDec,
			},
		},
		TotalAccumulatedRewards: sdk.SysCoins{testDecCoin},
	}

	return mc.cdc.MustMarshalJSON(farmPool)
}

// BuildFarmPoolNameList generates the farm pool name list bytes for test
func (mc *MockClient) BuildFarmPoolNameListBytes(poolName ...string) []byte {
	return mc.cdc.MustMarshalJSON(poolName)
}
