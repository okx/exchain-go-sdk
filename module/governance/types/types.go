package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/gov"
	govtypes "github.com/okex/okexchain/x/gov/types"
)

// const
const (
	ModuleName                         = govtypes.ModuleName
	ProposalsPath                      = "custom/gov/proposals"
	OptionYes           VoteOption     = 0x01
	OptionAbstain       VoteOption     = 0x02
	OptionNo            VoteOption     = 0x03
	OptionNoWithVeto    VoteOption     = 0x04
	StatusNil           ProposalStatus = 0x00
	StatusDepositPeriod ProposalStatus = 0x01
	StatusVotingPeriod  ProposalStatus = 0x02
	StatusPassed        ProposalStatus = 0x03
	StatusRejected      ProposalStatus = 0x04
	StatusFailed        ProposalStatus = 0x05
)

var (
	MsgCdc = gosdktypes.NewCodec()
)

func init() {
	RegisterCodec(MsgCdc)
}

// RegisterCodec registers the msg type for governance module
func RegisterCodec(cdc *codec.Codec) {
	gov.RegisterCodec(cdc)
}

// ProposalJSON - structure for a standard proposal from the JSON file
type ProposalJSON struct {
	Title        string
	Description  string
	ProposalType string
	Deposit      string
}

// VoteOption defines a vote option
type VoteOption byte

// MarshalJSON Marshals to JSON using string
func (vo VoteOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(vo.String())
}

// String implements the Stringer interface
func (vo VoteOption) String() string {
	switch vo {
	case OptionYes:
		return "Yes"
	case OptionAbstain:
		return "Abstain"
	case OptionNo:
		return "No"
	case OptionNoWithVeto:
		return "NoWithVeto"
	default:
		return ""
	}
}

// TallyResult - structure of the tally results statistics
type TallyResult struct {
	TotalPower      sdk.Dec `json:"total_power"`
	TotalVotedPower sdk.Dec `json:"total_voted_power"`
	Yes             sdk.Dec `json:"yes"`
	Abstain         sdk.Dec `json:"abstain"`
	No              sdk.Dec `json:"no"`
	NoWithVeto      sdk.Dec `json:"no_with_veto"`
}

// ProposalStatus is a type alias that represents a proposal status as a byte
type ProposalStatus byte

// MarshalJSON marshals to JSON using string
func (status ProposalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(status.String())
}

// UnmarshalJSON unmarshals from JSON assuming Bech32 encoding
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

// String implements the Stringer interface
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

	case StatusFailed:
		return "Failed"

	default:
		return ""
	}
}

// ProposalStatusToString turns a string into a ProposalStatus
func ProposalStatusFromString(status string) (ProposalStatus, error) {
	switch status {
	case "DepositPeriod", "deposit_period":
		return StatusDepositPeriod, nil

	case "VotingPeriod", "voting_period":
		return StatusVotingPeriod, nil

	case "Passed", "passed":
		return StatusPassed, nil

	case "Rejected", "rejected":
		return StatusRejected, nil

	case "Failed", "failed":
		return StatusFailed, nil

	case "":
		return StatusNil, nil

	default:
		return ProposalStatus(0xff), fmt.Errorf("failed. '%s' is not a valid proposal status", status)
	}
}
