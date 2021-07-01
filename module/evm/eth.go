package evm

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// BalanceAt returns the wei balance of the given account.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (ec evmClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return ec.ec.BalanceAt(ctx, account, blockNumber)
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (ec evmClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return ec.ec.TransactionReceipt(ctx, txHash)
}

// ChainId retrieves the current chain ID for transaction replay protection.
func (ec evmClient) ChainID(ctx context.Context) (*big.Int, error) {
	return ec.ec.ChainID(ctx)
}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec evmClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return ec.ec.SendTransaction(ctx, tx)
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (ec evmClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return ec.ec.PendingNonceAt(ctx, account)
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (ec evmClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return ec.ec.NonceAt(ctx, account, blockNumber)
}

// PendingCodeAt returns the contract code of the given account in the pending state.
func (ec evmClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return ec.ec.PendingCodeAt(ctx, account)
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (ec evmClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return ec.ec.CodeAt(ctx, account, blockNumber)
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (ec evmClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return ec.ec.EstimateGas(ctx, msg)
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (ec evmClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return ec.ec.SuggestGasPrice(ctx)
}

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
func (ec evmClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return ec.ec.CallContract(ctx, msg, blockNumber)
}

// BlockByNumber returns a block from the current canonical chain. If number is nil, the
// latest known block is returned.
//
// Note that loading full blocks requires two requests. Use HeaderByNumber
// if you don't need all transactions or uncle headers.
func (ec evmClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return ec.ec.BlockByNumber(ctx, number)
}

// FilterLogs executes a filter query.
func (ec evmClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return ec.ec.FilterLogs(ctx, q)
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (ec evmClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return ec.ec.SubscribeFilterLogs(ctx, q, ch)
}
