package types

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	ProposalTypeText            ProposalKind = 0x01
	ProposalTypeParameterChange ProposalKind = 0x02
	ProposalTypeAppUpgrade      ProposalKind = 0x03
	ProposalTypeDexList         ProposalKind = 0x04

	StatusNil           ProposalStatus = 0x00
	StatusDepositPeriod ProposalStatus = 0x01
	StatusVotingPeriod  ProposalStatus = 0x02
	StatusPassed        ProposalStatus = 0x03
	StatusRejected      ProposalStatus = 0x04
)

type Proposals []Proposal

type Proposal interface {
	GetProposalID() uint64
	SetProposalID(uint64)

	GetTitle() string
	SetTitle(string)

	GetDescription() string
	SetDescription(string)

	GetProposalType() ProposalKind
	SetProposalType(ProposalKind)

	GetStatus() ProposalStatus
	SetStatus(ProposalStatus)

	GetFinalTallyResult() TallyResult
	SetFinalTallyResult(TallyResult)

	GetSubmitTime() time.Time
	SetSubmitTime(time.Time)

	GetDepositEndTime() time.Time
	SetDepositEndTime(time.Time)

	GetTotalDeposit() DecCoins
	SetTotalDeposit(DecCoins)

	GetVotingStartTime() time.Time
	SetVotingStartTime(time.Time)

	GetVotingEndTime() time.Time
	SetVotingEndTime(time.Time)

	String() string

	GetProtocolDefinition() ProtocolDefinition
	SetProtocolDefinition(ProtocolDefinition)
}

type ProposalKind byte

// Unmarshals from JSON assuming Bech32 encoding
func (pt *ProposalKind) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	bz2, err := ProposalTypeFromString(s)
	if err != nil {
		return err
	}
	*pt = bz2
	return nil
}

// String to proposalType byte. Returns 0xff if invalid.
func ProposalTypeFromString(str string) (ProposalKind, error) {
	switch str {
	case "Text":
		return ProposalTypeText, nil
	case "ParameterChange":
		return ProposalTypeParameterChange, nil
	case "AppUpgrade":
		return ProposalTypeAppUpgrade, nil
	case "DexList":
		return ProposalTypeDexList, nil
	default:
		return ProposalKind(0xff), fmt.Errorf("'%s' is not a valid proposal type", str)
	}
}

type ProposalStatus byte

// Unmarshals from JSON assuming Bech32 encoding
func (status *ProposalStatus) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	bz2, err := ProposalStatusFromString(s)
	if err != nil {
		return err
	}
	*status = bz2
	return nil
}

// ProposalStatusToString turns a string into a ProposalStatus
func ProposalStatusFromString(str string) (ProposalStatus, error) {
	switch str {
	case "DepositPeriod":
		return StatusDepositPeriod, nil
	case "VotingPeriod":
		return StatusVotingPeriod, nil
	case "Passed":
		return StatusPassed, nil
	case "Rejected":
		return StatusRejected, nil
	case "":
		return StatusNil, nil
	default:
		return ProposalStatus(0xff), fmt.Errorf("'%s' is not a valid proposal status", str)
	}
}

type TallyResult struct {
	TotalBonded Dec `json:"total_bonded"`
	TotalVoting Dec `json:"total_voting"`
	Yes         Dec `json:"yes"`
	Abstain     Dec `json:"abstain"`
	No          Dec `json:"no"`
	NoWithVeto  Dec `json:"no_with_veto"`
}

// Basic Proposals
type BasicProposal struct {
	ProposalID   uint64       `json:"proposal_id"`   //  ID of the proposal
	Title        string       `json:"title"`         //  Title of the proposal
	Description  string       `json:"description"`   //  Description of the proposal
	ProposalType ProposalKind `json:"proposal_type"` //  Type of proposal. Initial set {PlainTextProposal, AppUpgradeProposal}

	Status           ProposalStatus `json:"proposal_status"` //  Status of the Proposal {Pending, Active, Passed, Rejected}
	FinalTallyResult TallyResult    `json:"tally_result"`    //  Result of Tallys

	SubmitTime     time.Time `json:"submit_time"`      //  Time of the block where TxGovSubmitProposal was included
	DepositEndTime time.Time `json:"deposit_end_time"` // Time that the Proposal would expire if deposit amount isn't met
	TotalDeposit   DecCoins  `json:"total_deposit"`    //  Current deposit on this proposal. Initial value is set at InitialDeposit

	VotingStartTime time.Time `json:"voting_start_time"` //  Time of the block where MinDeposit was reached. -1 if MinDeposit is not reached
	VotingEndTime   time.Time `json:"voting_end_time"`   // Time that the VotingPeriod for this proposal will end and votes will be tallied
}

// Implements Proposal Interface
var _ Proposal = (*BasicProposal)(nil)

// nolint
func (tp BasicProposal) GetProposalID() uint64                      { return tp.ProposalID }
func (tp *BasicProposal) SetProposalID(proposalID uint64)           { tp.ProposalID = proposalID }
func (tp BasicProposal) GetTitle() string                           { return tp.Title }
func (tp *BasicProposal) SetTitle(title string)                     { tp.Title = title }
func (tp BasicProposal) GetDescription() string                     { return tp.Description }
func (tp *BasicProposal) SetDescription(description string)         { tp.Description = description }
func (tp BasicProposal) GetProposalType() ProposalKind              { return tp.ProposalType }
func (tp *BasicProposal) SetProposalType(proposalType ProposalKind) { tp.ProposalType = proposalType }
func (tp BasicProposal) GetStatus() ProposalStatus                  { return tp.Status }
func (tp *BasicProposal) SetStatus(status ProposalStatus)           { tp.Status = status }
func (tp BasicProposal) GetFinalTallyResult() TallyResult           { return tp.FinalTallyResult }
func (tp *BasicProposal) SetFinalTallyResult(tallyResult TallyResult) {
	tp.FinalTallyResult = tallyResult
}
func (tp BasicProposal) GetSubmitTime() time.Time            { return tp.SubmitTime }
func (tp *BasicProposal) SetSubmitTime(submitTime time.Time) { tp.SubmitTime = submitTime }
func (tp BasicProposal) GetDepositEndTime() time.Time        { return tp.DepositEndTime }
func (tp *BasicProposal) SetDepositEndTime(depositEndTime time.Time) {
	tp.DepositEndTime = depositEndTime
}
func (tp BasicProposal) GetTotalDeposit() DecCoins              { return tp.TotalDeposit }
func (tp *BasicProposal) SetTotalDeposit(totalDeposit DecCoins) { tp.TotalDeposit = totalDeposit }
func (tp BasicProposal) GetVotingStartTime() time.Time          { return tp.VotingStartTime }
func (tp *BasicProposal) SetVotingStartTime(votingStartTime time.Time) {
	tp.VotingStartTime = votingStartTime
}
func (tp BasicProposal) GetVotingEndTime() time.Time { return tp.VotingEndTime }
func (tp *BasicProposal) SetVotingEndTime(votingEndTime time.Time) {
	tp.VotingEndTime = votingEndTime
}

func (tp BasicProposal) String() string {
	return fmt.Sprintf(`Proposal %d:
  Title:              %s
  Type:               %d
  Status:             %d
  Submit Time:        %s
  Deposit End Time:   %s
  Total Deposit:      %s
  Voting Start Time:  %s
  Voting End Time:    %s`, tp.ProposalID, tp.Title, tp.ProposalType,
		tp.Status, tp.SubmitTime, tp.DepositEndTime,
		tp.TotalDeposit, tp.VotingStartTime, tp.VotingEndTime)
}

// software upgrade
func (tp BasicProposal) GetProtocolDefinition() ProtocolDefinition {
	return ProtocolDefinition{}
}
func (tp *BasicProposal) SetProtocolDefinition(ProtocolDefinition) {}

// Text Proposals
type TextProposal struct {
	BasicProposal
}

// Implements Proposal Interface
var _ Proposal = (*TextProposal)(nil)

// DexList Proposals
type DexListProposal struct {
	BasicProposal
	Proposer      AccAddress `json:"proposer"`    //  Proposer of proposal
	ListAsset     string         `json:"list_asset"`  //  Symbol of asset listed on Dex.
	QuoteAsset    string         `json:"quote_asset"` //  Symbol of asset quoted by asset listed on Dex.
	InitPrice     Dec        `json:"init_price"`  //  Init price of asset listed on Dex.
	BlockHeight   uint64         `json:"block_height"`
	MaxPriceDigit uint64         `json:"max_price_digit"` //  Decimal of price
	MaxSizeDigit  uint64         `json:"max_size_digit"`  //  Decimal of trade quantity
	MinTradeSize  string         `json:"min_trade_size"`

	DexListStartTime time.Time `json:"dex_list_start_time"`
	DexListEndTime   time.Time `json:"dex_list_end_time"`
}
// Implements Proposal Interface
var _ Proposal = (*DexListProposal)(nil)

func (tp DexListProposal) String() string {
	return fmt.Sprintf(`Proposal %d:
  Title:               %s
  Type:                %d
  Proposer:            %s
  Status:              %d
  Submit Time:         %s
  Deposit End Time:    %s
  Total Deposit:       %s
  Voting Start Time:   %s
  Voting End Time:     %s
  ListAsset            %s
  QuoteAsset           %s
  InitPrice            %s
  BlockHeight          %d
  MaxPriceDigit        %d
  MaxSizeDigit         %d
  MinTradeSize         %s
  Dex List Start Time: %s
  Dex List End Time: %s`, tp.ProposalID, tp.Title, tp.ProposalType, tp.Proposer,
		tp.Status, tp.SubmitTime, tp.DepositEndTime,
		tp.TotalDeposit, tp.VotingStartTime, tp.VotingEndTime, tp.ListAsset, tp.QuoteAsset, tp.InitPrice,
		tp.BlockHeight, tp.MaxPriceDigit, tp.MaxSizeDigit, tp.MinTradeSize, tp.DexListStartTime, tp.DexListEndTime)
}

// Implements Proposal Interface
var _ Proposal = (*ParameterProposal)(nil)

type ParameterProposal struct {
	BasicProposal
	Params Params `json:"params"`
	Height int64  `json:"height"`
}

type Param struct {
	Subspace string `json:"subspace"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type Params []Param


var _ Proposal = (*AppUpgradeProposal)(nil)

type AppUpgradeProposal struct {
	BasicProposal
	ProtocolDefinition ProtocolDefinition
}

func (aup AppUpgradeProposal) GetProtocolDefinition() ProtocolDefinition {
	return aup.ProtocolDefinition
}

func (aup *AppUpgradeProposal) SetProtocolDefinition(protocolDefinition ProtocolDefinition) {
	aup.ProtocolDefinition = protocolDefinition
}

func (aup AppUpgradeProposal) String() string {
	return fmt.Sprintf(`Proposal %d:
  Title:              %s
  Type:               %d
  Status:             %d
  Submit Time:        %s
  Deposit End Time:   %s
  Total Deposit:      %s
  Voting Start Time:  %s
  Voting End Time:    %s
  Version:            %d
  Software:           %s
  Switch Height:      %d
  Threshold:          %s`, aup.ProposalID, aup.Title, aup.ProposalType, aup.Status, aup.SubmitTime, aup.DepositEndTime,
		aup.TotalDeposit, aup.VotingStartTime, aup.VotingEndTime, aup.ProtocolDefinition.Version,
		aup.ProtocolDefinition.Software, aup.ProtocolDefinition.Height, aup.ProtocolDefinition.Threshold.String())
}
