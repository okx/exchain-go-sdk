package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	ModuleName = "governance"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for governance module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgSubmitProposal{}, "okchain/gov/MsgSubmitProposal")
	cdc.RegisterInterface((*Content)(nil))
	cdc.RegisterConcrete(TextProposal{}, "okchain/gov/TextProposal")
}

// Proposal - structure for a standard proposal from the JSON file
type Proposal struct {
	Title        string
	Description  string
	ProposalType string
	Deposit      string
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
