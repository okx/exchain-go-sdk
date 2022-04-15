package mocks

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/okex/exchain-go-sdk/exposed"
	auth "github.com/okex/exchain-go-sdk/module/auth/types"
	dex "github.com/okex/exchain-go-sdk/module/dex/types"
	distribution "github.com/okex/exchain-go-sdk/module/distribution/types"
	evm "github.com/okex/exchain-go-sdk/module/evm/types"
	farm "github.com/okex/exchain-go-sdk/module/farm/types"
	governance "github.com/okex/exchain-go-sdk/module/governance/types"
	order "github.com/okex/exchain-go-sdk/module/order/types"
	slashing "github.com/okex/exchain-go-sdk/module/slashing/types"
	staking "github.com/okex/exchain-go-sdk/module/staking/types"
	tendermint "github.com/okex/exchain-go-sdk/module/tendermint/types"
	token "github.com/okex/exchain-go-sdk/module/token/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	evmtypes "github.com/okex/exchain/app/types"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	authtypes "github.com/okex/exchain/libs/cosmos-sdk/x/auth/types"
	abci "github.com/okex/exchain/libs/tendermint/abci/types"
	"github.com/okex/exchain/libs/tendermint/crypto"
	tmbytes "github.com/okex/exchain/libs/tendermint/libs/bytes"
	"github.com/okex/exchain/libs/tendermint/libs/kv"
	ctypes "github.com/okex/exchain/libs/tendermint/rpc/core/types"
	tmtypes "github.com/okex/exchain/libs/tendermint/types"
	govtypes "github.com/okex/exchain/x/gov/types"
	orderkeeper "github.com/okex/exchain/x/order/keeper"
	stakingtypes "github.com/okex/exchain/x/staking/types"
	"github.com/stretchr/testify/require"
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
func (mc *MockClient) GetCodec() *codec.Codec {
	return mc.cdc
}

// nolint
func (mc *MockClient) Auth() exposed.Auth {
	return mc.modules[auth.ModuleName].(exposed.Auth)
}
func (mc *MockClient) Dex() exposed.Dex {
	return mc.modules[dex.ModuleName].(exposed.Dex)
}
func (mc *MockClient) Distribution() exposed.Distribution {
	return mc.modules[distribution.ModuleName].(exposed.Distribution)
}
func (mc *MockClient) Evm() exposed.Evm {
	return mc.modules[evm.ModuleName].(exposed.Evm)
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
	mockEvents := []abci.Event{
		{
			Type: eventType,
			Attributes: []kv.Pair{
				{
					Key: kvPairKey,
				},
			},
		},
	}
	return &ctypes.ResultBlockResults{
		Height: height,
		TxsResults: []*abci.ResponseDeliverTx{
			{
				Events: mockEvents,
			},
		},
		BeginBlockEvents: mockEvents,
		EndBlockEvents:   mockEvents,
		ValidatorUpdates: []abci.ValidatorUpdate{
			{
				PubKey: abci.PubKey{
					Type: pkType,
				},
				Power: power,
			},
		},
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
func (mc *MockClient) GetRawTxResultPointer(height int64, code uint32, log, hashHexStr, eventType string, tx []byte) *ctypes.ResultTx {
	hash, err := hex.DecodeString(hashHexStr)
	require.NoError(mc.t, err)

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
func (mc *MockClient) GetRawResultTxSearchPointer(totalCount int, height int64, code uint32, log, hashHexStr, eventType string,
	tx []byte) *ctypes.ResultTxSearch {
	return &ctypes.ResultTxSearch{
		TotalCount: totalCount,
		Txs: []*ctypes.ResultTx{
			mc.GetRawTxResultPointer(height, code, log, hashHexStr, eventType, tx),
		},
	}
}

// BuildProposalsBytes generates the proposals bytes for test
func (mc *MockClient) BuildProposalsBytes(proposalID uint64, status govtypes.ProposalStatus,
	mockTime time.Time, totalDeposit sdk.DecCoins, mockPower sdk.Dec) []byte {
	proposals := []governance.Proposal{
		{
			Content:         govtypes.TextProposal{},
			ProposalID:      proposalID,
			Status:          status,
			SubmitTime:      mockTime,
			DepositEndTime:  mockTime,
			VotingStartTime: mockTime,
			VotingEndTime:   mockTime,
			TotalDeposit:    totalDeposit,
			FinalTallyResult: govtypes.TallyResult{
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

// BuildAccAddrList generates the account address list bytes for test
func (mc *MockClient) BuildAccAddrListBytes(accAddr ...sdk.AccAddress) []byte {
	return mc.cdc.MustMarshalJSON(accAddr)
}

// BuildLockInfoBytes generates the lock info bytes for test
func (mc *MockClient) BuildLockInfoBytes(accAddr sdk.AccAddress, poolName, tokenSymbol string, amountDec sdk.Dec, height int64,
	referencePeriod uint64) []byte {
	info := farm.LockInfo{
		Owner:            accAddr,
		PoolName:         poolName,
		Amount:           sdk.NewDecCoinFromDec(tokenSymbol, amountDec),
		StartBlockHeight: height,
		ReferencePeriod:  referencePeriod,
	}
	return mc.cdc.MustMarshalJSON(info)
}

// BuildQueryResCode generates query res code bytes for test
func (mc *MockClient) BuildQueryResCode(codeStr string) []byte {
	info := evm.QueryResCode{
		Code: []byte(codeStr),
	}
	return mc.cdc.MustMarshalJSON(info)
}

// BuildQueryResStorage generates query res storage bytes for test
func (mc *MockClient) BuildQueryResStorage(storageStr string) []byte {
	info := evm.QueryResStorage{
		Value: []byte(storageStr),
	}
	return mc.cdc.MustMarshalJSON(info)
}
