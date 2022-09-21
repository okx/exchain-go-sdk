package exposed

import (
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	"github.com/okex/exchain/x/wasm/types"
)
import gosdktypes "github.com/okex/exchain-go-sdk/types"

// Wasm shows the expected behavior for inner wasm client
type Wasm interface {
	gosdktypes.Module
	wasmTx
	wasmQuery
}

type wasmTx interface {
}

type wasmQuery interface {
	// QueryListCode query all wasm bytecode on the chain
	QueryListCode(pageReq *query.PageRequest) (*types.QueryCodesResponse, error)

	// QueryListContract query all bytecode on the chain for given code id
	QueryListContract(codeID uint64, pageReq *query.PageRequest) (*types.QueryContractsByCodeResponse, error)

	// QueryCode query wasm bytecode for given code id
	QueryCode(codeID uint64) (*types.QueryCodeResponse, error)

	// QueryCodeInfo query metadata of code for given code id
	QueryCodeInfo(codeID uint64) (*types.CodeInfoResponse, error)

	// QueryContractInfo query metadata of a contract given its address
	QueryContractInfo(address string) (*types.QueryContractInfoResponse, error)

	// QueryContractHistory query the code history for a contract given its address
	QueryContractHistory(address string, pageReq *query.PageRequest) (*types.QueryContractHistoryResponse, error)

	// QueryContractStateAll query all internal state of a contract given its address
	QueryContractStateAll(address string, pageReq *query.PageRequest) (*types.QueryAllContractStateResponse, error)

	// QueryContractStateRaw query internal state for key of a contract given its address
	QueryContractStateRaw(address string, queryData string) (*types.QueryRawContractStateResponse, error)

	// QueryContractStateSmart query contract with given address with query data
	QueryContractStateSmart(address string, queryData string) (*types.QuerySmartContractStateResponse, error)

	// QueryListPinnedCode query all pinned code ids
	QueryListPinnedCode(pageReq *query.PageRequest) (*types.QueryPinnedCodesResponse, error)
}
