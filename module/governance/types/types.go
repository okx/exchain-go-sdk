package types

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/okex/okexchain-go-sdk/types"
)

// const
const (
	ModuleName                         = "governance"
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
	MsgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(MsgCdc)
}

// RegisterCodec registers the msg type for governance module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgSubmitProposal{}, "okchain/gov/MsgSubmitProposal")
	cdc.RegisterConcrete(MsgDeposit{}, "okchain/gov/MsgDeposit")
	cdc.RegisterConcrete(MsgVote{}, "okchain/gov/MsgVote")

	cdc.RegisterInterface((*Content)(nil))
	cdc.RegisterConcrete(TextProposal{}, "okchain/gov/TextProposal")
	cdc.RegisterConcrete(ParameterChangeProposal{}, "okchain/params/ParameterChangeProposal")
	cdc.RegisterConcrete(DelistProposal{}, "okchain/dex/DelistProposal")
	cdc.RegisterConcrete(CommunityPoolSpendProposal{}, "okchain/distribution/CommunityPoolSpendProposal")
}

type (
	// ProposalJSON - structure for a standard proposal from the JSON file
	ProposalJSON struct {
		Title        string
		Description  string
		ProposalType string
		Deposit      string
	}

	// ParamChangeProposalJSON - structure for a ParamChangeProposal with a deposit used to parse parameter change proposals
	// from the JSON file
	ParamChangeProposalJSON struct {
		Title       string           `json:"title"`
		Description string           `json:"description"`
		Changes     ParamChangesJSON `json:"changes"`
		Deposit     sdk.DecCoins     `json:"deposit"`
		Height      uint64           `json:"height"`
	}

	// DelistProposalJSON - structure for a DelistProposal with a deposit used to parse delist proposals from the JSON file
	DelistProposalJSON struct {
		Title       string       `json:"title"`
		Description string       `json:"description"`
		BaseAsset   string       `json:"base_asset"`
		QuoteAsset  string       `json:"quote_asset"`
		Deposit     sdk.DecCoins `json:"deposit"`
	}

	// CommunityPoolSpendProposalJSON - structure for a CommunityPoolSpendProposal used to parse community pool spend proposals
	// from the JSON file
	CommunityPoolSpendProposalJSON struct {
		Title       string         `json:"title"`
		Description string         `json:"description"`
		Recipient   sdk.AccAddress `json:"recipient"`
		Amount      sdk.DecCoins   `json:"amount"`
		Deposit     sdk.DecCoins   `json:"deposit"`
	}
)

// ParamChangesJSON defines a slice of ParamChangeJSON objects which can be converted to a slice of ParamChange objects
type ParamChangesJSON []ParamChangeJSON

// ToParamChanges converts a slice of ParamChangeJSON objects to a slice of ParamChange
func (pcj ParamChangesJSON) ToParamChanges() []ParamChange {
	res := make([]ParamChange, len(pcj))
	for i, pc := range pcj {
		res[i] = pc.ToParamChange()
	}
	return res
}

// ParamChangeJSON defines a parameter change used in JSON input
// this allows values to be specified in raw JSON instead of being string encoded
type ParamChangeJSON struct {
	Subspace string          `json:"subspace"`
	Key      string          `json:"key"`
	Subkey   string          `json:"subkey,omitempty"`
	Value    json.RawMessage `json:"value"`
}

// ToParamChange converts a ParamChangeJSON object to ParamChange
func (pcj ParamChangeJSON) ToParamChange() ParamChange {
	return ParamChange{
		Subspace: pcj.Subspace,
		Key:      pcj.Key,
		Subkey:   pcj.Subkey,
		Value:    string(pcj.Value),
	}
}

// Content defines an interface that a proposal must implement
// It contains information such as the title and description along with the type and routing information for the appropriate
// handler to process the proposal
type Content interface {
	GetTitle() string
	GetDescription() string
	ProposalRoute() string
	ProposalType() string
	ValidateBasic() sdk.Error
	String() string
}

var (
	_ Content = (*TextProposal)(nil)
	_ Content = (*ParameterChangeProposal)(nil)
	_ Content = (*DelistProposal)(nil)
	_ Content = (*CommunityPoolSpendProposal)(nil)
)

// TextProposal - structure for a text proposal that implements interface Content
type TextProposal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// NewTextProposal is a constructor function for TextProposal
func NewTextProposal(title, description string) Content {
	return TextProposal{title, description}
}

// nolint
func (TextProposal) GetTitle() string         { return "" }
func (TextProposal) GetDescription() string   { return "" }
func (TextProposal) ProposalRoute() string    { return "" }
func (TextProposal) ProposalType() string     { return "" }
func (TextProposal) String() string           { return "" }
func (TextProposal) ValidateBasic() sdk.Error { return nil }

// ParameterChangeProposal - structure of a param change proposal that implements interface Content
type ParameterChangeProposal struct {
	SDKParameterChangeProposal `json:"ParameterChangeProposal"`
	Height                     uint64 `json:"height"`
}

// ParamChange - structure to define a parameter change
type ParamChange struct {
	Subspace string `json:"subspace"`
	Key      string `json:"key"`
	Subkey   string `json:"subkey,omitempty"`
	Value    string `json:"value"`
}

// NewParameterChangeProposal is a constructor function for ParameterChangeProposal
func NewParameterChangeProposal(title, description string, changes []ParamChange, height uint64) ParameterChangeProposal {
	return ParameterChangeProposal{
		SDKParameterChangeProposal: SDKParameterChangeProposal{
			Title:       title,
			Description: description,
			Changes:     changes,
		},
		Height: height,
	}
}

// SDKParameterChangeProposal defines a proposal which contains multiple parameter changes under cosmos-sdk
type SDKParameterChangeProposal struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Changes     []ParamChange `json:"changes"`
}

// nolint
func (ParameterChangeProposal) GetTitle() string         { return "" }
func (ParameterChangeProposal) GetDescription() string   { return "" }
func (ParameterChangeProposal) ProposalRoute() string    { return "" }
func (ParameterChangeProposal) ProposalType() string     { return "" }
func (ParameterChangeProposal) String() string           { return "" }
func (ParameterChangeProposal) ValidateBasic() sdk.Error { return nil }

// DelistProposal - structure of a delist proposal that implements interface Content
type DelistProposal struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Proposer    sdk.AccAddress `json:"proposer"`
	BaseAsset   string         `json:"base_asset"`
	QuoteAsset  string         `json:"quote_asset"`
}

// NewDelistProposal is a constructor function for DelistProposal
func NewDelistProposal(title, description string, proposer sdk.AccAddress, baseAsset, quoteAsset string,
) DelistProposal {
	return DelistProposal{
		Title:       title,
		Description: description,
		Proposer:    proposer,
		BaseAsset:   baseAsset,
		QuoteAsset:  quoteAsset,
	}
}

// nolint
func (DelistProposal) GetTitle() string         { return "" }
func (DelistProposal) GetDescription() string   { return "" }
func (DelistProposal) ProposalRoute() string    { return "" }
func (DelistProposal) ProposalType() string     { return "" }
func (DelistProposal) String() string           { return "" }
func (DelistProposal) ValidateBasic() sdk.Error { return nil }

// CommunityPoolSpendProposal - structure of a community pool spend proposal that implements interface Content
type CommunityPoolSpendProposal struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Recipient   sdk.AccAddress `json:"recipient"`
	Amount      sdk.DecCoins   `json:"amount"`
}

// NewCommunityPoolSpendProposal is a constructor function for CommunityPoolSpendProposal
func NewCommunityPoolSpendProposal(title, description string, recipient sdk.AccAddress, amount sdk.DecCoins,
) CommunityPoolSpendProposal {
	return CommunityPoolSpendProposal{
		title,
		description,
		recipient,
		amount,
	}
}

// nolint
func (CommunityPoolSpendProposal) GetTitle() string         { return "" }
func (CommunityPoolSpendProposal) GetDescription() string   { return "" }
func (CommunityPoolSpendProposal) ProposalRoute() string    { return "" }
func (CommunityPoolSpendProposal) ProposalType() string     { return "" }
func (CommunityPoolSpendProposal) String() string           { return "" }
func (CommunityPoolSpendProposal) ValidateBasic() sdk.Error { return nil }

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

// Proposal - structure used by the governance module to allow for voting on network changes
type Proposal struct {
	Content          `json:"content"`
	ProposalID       uint64         `json:"id"`
	Status           ProposalStatus `json:"proposal_status"`
	FinalTallyResult TallyResult    `json:"final_tally_result"`
	SubmitTime       time.Time      `json:"submit_time"`
	DepositEndTime   time.Time      `json:"deposit_end_time"`
	TotalDeposit     sdk.DecCoins   `json:"total_deposit"`
	VotingStartTime  time.Time      `json:"voting_start_time"`
	VotingEndTime    time.Time      `json:"voting_end_time"`
}

// String returns a human readable string representation of Proposal
func (p Proposal) String() string {
	return fmt.Sprintf(`Proposal %d:
  Title:              %s
  Type:               %s
  Status:             %s
  Submit Time:        %s
  Deposit End Time:   %s
  Total Deposit:      %s
  Voting Start Time:  %s
  Voting End Time:    %s
  Description:        %s`,
		p.ProposalID, p.GetTitle(), p.ProposalType(),
		p.Status, p.SubmitTime, p.DepositEndTime,
		p.TotalDeposit, p.VotingStartTime, p.VotingEndTime, p.GetDescription(),
	)
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
