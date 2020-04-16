package params

import (
	"encoding/json"
	"fmt"
	"github.com/okex/okchain-go-sdk/types"
	"time"
)

const (
	DefaultBookSize = 200
	DefaultPage     = 1
	DefaultPerPage  = 50

	StatusDepositPeriod ProposalStatus = 0x01
	StatusVotingPeriod  ProposalStatus = 0x02
	StatusPassed        ProposalStatus = 0x03
	StatusRejected      ProposalStatus = 0x04
)

// QueryAccTokenParams - structure of params to query a specific token in an account
type QueryAccTokenParams struct {
	Symbol string `json:"symbol"`
	Show   string `json:"show"`
}

// NewQueryAccTokenParams creates a new instance of QueryAccTokenParams
func NewQueryAccTokenParams(symbol, show string) QueryAccTokenParams {
	return QueryAccTokenParams{
		symbol,
		show,
	}
}

type QueryDepthBookParams struct {
	Product string
	Size    int
}

func NewQueryDepthBookParams(product string, size int) QueryDepthBookParams {
	if size == 0 {
		size = DefaultBookSize
	}
	return QueryDepthBookParams{
		Product: product,
		Size:    size,
	}
}

type QueryKlinesParams struct {
	Product     string
	Granularity int
	Size        int
}

func NewQueryKlinesParams(product string, granularity, size int) QueryKlinesParams {
	return QueryKlinesParams{
		product,
		granularity,
		size,
	}
}

type QueryTickerParams struct {
	Product string `json:"product"`
	Count   int    `json:"count"`
	Sort    bool   `json:"sort"`
}

func NewQueryTickerParams(product string, count int, sort bool) QueryTickerParams {
	return QueryTickerParams{
		product,
		count,
		sort,
	}
}

// QueryMatchParams - structure of params to query tx record
type QueryMatchParams struct {
	Product string
	Start   int64
	End     int64
	Page    int
	PerPage int
}

// NewQueryMatchParams creates a new instance of QueryMatchParams
func NewQueryMatchParams(product string, start, end int64, page, perPage int) QueryMatchParams {
	if page == 0 && perPage == 0 {
		page = DefaultPage
		perPage = DefaultPerPage
	}
	return QueryMatchParams{
		Product: product,
		Start:   start,
		End:     end,
		Page:    page,
		PerPage: perPage,
	}
}

type QueryOrderListParams struct {
	Address    string
	Product    string
	Page       int
	PerPage    int
	Start      int64
	End        int64
	Side       string
	HideNoFill bool
}

// creates a new instance of NewQueryOrderListParams
func NewQueryOrderListParams(addr, product, side string, page, perPage int, start, end int64,
	hideNoFill bool) QueryOrderListParams {
	if page == 0 && perPage == 0 {
		page = DefaultPage
		perPage = DefaultPerPage
	}
	if start == 0 && end == 0 {
		end = time.Now().Unix()
	}
	return QueryOrderListParams{
		Address:    addr,
		Product:    product,
		Page:       page,
		PerPage:    perPage,
		Start:      start,
		End:        end,
		Side:       side,
		HideNoFill: hideNoFill,
	}
}

type QueryDealsParams struct {
	Address string
	Product string
	Start   int64
	End     int64
	Page    int
	PerPage int
	Side    string
}

func NewQueryDealsParams(addr, product string, start, end int64, page, perPage int, side string) QueryDealsParams {
	if page == 0 && perPage == 0 {
		page = DefaultPage
		perPage = DefaultPerPage
	}
	return QueryDealsParams{
		Address: addr,
		Product: product,
		Start:   start,
		End:     end,
		Page:    page,
		PerPage: perPage,
		Side:    side,
	}
}

type QueryTxListParams struct {
	Address   string
	TxType    int64
	StartTime int64
	EndTime   int64
	Page      int
	PerPage   int
}

// creates a new instance of NewQueryOrderListParams
func NewQueryTxListParams(addr string, txType, startTime, endTime int64, page, perPage int) QueryTxListParams {
	if page == 0 && perPage == 0 {
		page = DefaultPage
		perPage = DefaultPerPage
	}
	return QueryTxListParams{
		Address:   addr,
		TxType:    txType,
		StartTime: startTime,
		EndTime:   endTime,
		Page:      page,
		PerPage:   perPage,
	}
}

type QueryProposalsParams struct {
	Voter          types.AccAddress
	Depositor      types.AccAddress
	ProposalStatus ProposalStatus
	Limit          uint64
}

type ProposalStatus byte

func (status ProposalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.String())
}

func (status ProposalStatus) String() string {
	switch status {
	case StatusDepositPeriod:
		return "DepositPeriod"
	case StatusVotingPeriod:
		return "VotingPeriod"
	case StatusPassed:
		return "Passed"
	case StatusRejected:
		return "Rejected"
	default:
		return ""
	}
}

func NewQueryProposalsParams(status ProposalStatus, limit uint64, voter, depositor types.AccAddress) QueryProposalsParams {
	return QueryProposalsParams{
		Voter:          voter,
		Depositor:      depositor,
		ProposalStatus: status,
		Limit:          limit,
	}
}

type QueryProposalParams struct {
	ProposalID uint64
}

// creates a new instance of QueryProposalParams
func NewQueryProposalParams(proposalID uint64) QueryProposalParams {
	return QueryProposalParams{
		ProposalID: proposalID,
	}
}

// QueryDelegatorParams defines query params of delegator info
type QueryDelegatorParams struct {
	DelegatorAddr types.AccAddress
}

// NewQueryDelegatorParams creates a new instance of QueryDelegatorParams
func NewQueryDelegatorParams(delegatorAddr types.AccAddress) QueryDelegatorParams {
	return QueryDelegatorParams{
		DelegatorAddr: delegatorAddr,
	}
}

// QueryDexInfoParams defines query params of dex info
type QueryDexInfoParams struct {
	Owner   string
	Page    int
	PerPage int
}

// NewQueryDexInfoParams creates query params of dex info
func NewQueryDexInfoParams(owner string, page, perPage int) (QueryDexInfoParams, error) {
	if len(owner) == 0 {
		owner = ""
	} else {
		_, err := types.AccAddressFromBech32(owner)
		if err != nil {
			return QueryDexInfoParams{}, fmt.Errorf("invalid address：%s", owner)
		}
	}

	if page <= 0 {
		return QueryDexInfoParams{}, fmt.Errorf("invalid page：%d", page)
	}
	if perPage <= 0 {
		return QueryDexInfoParams{}, fmt.Errorf("invalid per-page：%d", perPage)
	}
	return QueryDexInfoParams{
		Owner:   owner,
		Page:    page,
		PerPage: perPage,
	}, nil
}
