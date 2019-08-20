package types

import (
	//"bytes"
	//"fmt"
	//
	//"github.com/pkg/errors"
	//"github.com/tendermint/tendermint/crypto/tmhash"

	amino "github.com/tendermint/go-amino"

	"github.com/tendermint/tendermint/crypto"
	//"github.com/tendermint/tendermint/crypto/merkle"
)

// Evidence represents any provable malicious activity by a validator
type Evidence interface {
	Height() int64                                     // height of the equivocation
	Address() []byte                                   // address of the equivocating validator
	Bytes() []byte                                     // bytes which compromise the evidence
	Hash() []byte                                      // hash of the evidence
	Verify(chainID string, pubKey crypto.PubKey) error // verify the evidence
	Equal(Evidence) bool                               // check equality of evidence

	ValidateBasic() error
	String() string
}

func RegisterEvidences(cdc *amino.Codec) {
	cdc.RegisterInterface((*Evidence)(nil), nil)
	cdc.RegisterConcrete(&DuplicateVoteEvidence{}, "tendermint/DuplicateVoteEvidence", nil)
}

// DuplicateVoteEvidence contains evidence a validator signed two conflicting
// votes.
type DuplicateVoteEvidence struct {
	PubKey crypto.PubKey
	VoteA  *Vote
	VoteB  *Vote
}

// EvidenceList is a list of Evidence. Evidences is not a word.
type EvidenceList []Evidence