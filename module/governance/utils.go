package governance

import (
	"encoding/json"
	"github.com/okex/okchain-go-sdk/module/governance/types"
	"github.com/okex/okchain-go-sdk/utils"
	"io/ioutil"
)

func parseProposalFromFile(path string) (proposal types.ProposalJSON, err error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if err = json.Unmarshal(contents, &proposal); err != nil {
		return proposal, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

func parseParamChangeProposalFromFile(path string) (proposal types.ParamChangeProposalJSON, err error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if err = types.MsgCdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

func parseDelistProposalFromFile(path string) (proposal types.DelistProposalJSON, err error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if err = types.MsgCdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}

func parseCommunityPoolSpendProposalFromFile(path string) (proposal types.CommunityPoolSpendProposalJSON, err error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if err = types.MsgCdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
