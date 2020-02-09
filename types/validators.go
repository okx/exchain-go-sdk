package types

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

type ResultValidatorsOutput struct {
	BlockHeight int64             `json:"block_height"`
	Validators  []ValidatorOutput `json:"validators"`
}

func NewResultValidatorsOutput(rv *ctypes.ResultValidators) (rvo ResultValidatorsOutput, err error) {
	rvo.BlockHeight = rv.BlockHeight
	valNum := len(rv.Validators)
	rvo.Validators = make([]ValidatorOutput, valNum)
	for i := 0; i < valNum; i++ {
		if rvo.Validators[i], err = bech32ValidatorOutput(rv.Validators[i]); err != nil {
			return ResultValidatorsOutput{}, err
		}
	}
	return
}

type ValidatorOutput struct {
	Address          ConsAddress `json:"address"`
	PubKey           string      `json:"pub_key"`
	ProposerPriority int64       `json:"proposer_priority"`
	VotingPower      int64       `json:"voting_power"`
}

type ConsAddress []byte

func bech32ValidatorOutput(validator *types.Validator) (ValidatorOutput, error) {
	bechValPubkey, err := Bech32ifyConsPub(validator.PubKey)
	if err != nil {
		return ValidatorOutput{}, err
	}

	return ValidatorOutput{
		Address:          ConsAddress(validator.Address),
		PubKey:           bechValPubkey,
		ProposerPriority: validator.ProposerPriority,
		VotingPower:      validator.VotingPower,
	}, nil
}
