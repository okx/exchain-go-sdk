package types

import (
	"encoding/json"

	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	ModuleName = "governance"
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
	cdc.RegisterInterface((*Content)(nil))
	cdc.RegisterConcrete(TextProposal{}, "okchain/gov/TextProposal")
	cdc.RegisterConcrete(ParameterChangeProposal{}, "okchain/params/ParameterChangeProposal")
}

type (
	// Proposal - structure for a standard proposal from the JSON file
	Proposal struct {
		Title        string
		Description  string
		ProposalType string
		Deposit      string
	}

	// Proposal - structure for a standard proposal from the JSON file
	ParamChangeProposal struct {
		Title       string           `json:"title"`
		Description string           `json:"description"`
		Changes     ParamChangesJSON `json:"changes"`
		Deposit     sdk.DecCoins     `json:"deposit"`
		Height      uint64           `json:"height"`
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
