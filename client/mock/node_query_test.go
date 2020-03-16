package mock

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/client"
	"github.com/okex/okchain-go-sdk/common/libs/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tendermint/go-amino"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"testing"
)

func TestQueryBlock(t *testing.T) {
	Convey("TestQueryBlock", t, func() {
		Convey("Test return valid output", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var height int64 = 100
			result := `{"block_meta":{"block_id":{"hash":"D413558AC98D51D49BAE7CFCBF980314CAF7CB9AC231EB8BF49DBCA4BA13CB20","parts":{"total":"1","hash":"C02AACFF657A58E51B7599262DBA5EBB65AB68992CA3BD7C01E68CB1F8B8B016"}},"header":{"version":{"block":"10","app":"0"},"chain_id":"okchain","height":"12","time":"2020-01-09T05:16:20.229297379Z","num_txs":"0","total_txs":"0","last_block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"last_commit_hash":"6397F129717B7AD48C37144F9F69B130ED1444E1F7FF429F044AC0E2045C1D19","data_hash":"","validators_hash":"A809FCB6E86D5AF31AD1F5BB6A684AD7B189F1750985EA76022C5933CE824E43","next_validators_hash":"A809FCB6E86D5AF31AD1F5BB6A684AD7B189F1750985EA76022C5933CE824E43","consensus_hash":"048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F","app_hash":"8800CD15175047DB2F312EB1A8CA9F074BFE43209FB3BF99AE35842B4AD08193","last_results_hash":"","evidence_hash":"","proposer_address":"7A503503558C8281B3D00455DDB3497B958F23DA"}},"block":{"header":{"version":{"block":"10","app":"0"},"chain_id":"okchain","height":"12","time":"2020-01-09T05:16:20.229297379Z","num_txs":"0","total_txs":"0","last_block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"last_commit_hash":"6397F129717B7AD48C37144F9F69B130ED1444E1F7FF429F044AC0E2045C1D19","data_hash":"","validators_hash":"A809FCB6E86D5AF31AD1F5BB6A684AD7B189F1750985EA76022C5933CE824E43","next_validators_hash":"A809FCB6E86D5AF31AD1F5BB6A684AD7B189F1750985EA76022C5933CE824E43","consensus_hash":"048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F","app_hash":"8800CD15175047DB2F312EB1A8CA9F074BFE43209FB3BF99AE35842B4AD08193","last_results_hash":"","evidence_hash":"","proposer_address":"7A503503558C8281B3D00455DDB3497B958F23DA"},"data":{"txs":null},"evidence":{"evidence":null},"last_commit":{"block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"precommits":[{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.127812004Z","validator_address":"128276A947E61097569338672A391EAAA348226A","validator_index":"0","signature":"jUecU4WfQi0SN0GDJqoykhqqn7Lck+2Hx/P54wnx1TcYxbb891VuDI3NtYEYnHq0mZPNX9NKAZxBc6hbW3UjCg=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.200295355Z","validator_address":"1908D39C6EAA69D6E756DBF3C309FC43B249C45E","validator_index":"1","signature":"9G9w8UDmZC6d7eB+jEnrrZtjl/2CKRlCIkYqAovYPWohkHoJQjDAn2O1brq4CbvsYLJS2kkHJ1bKPRyhkf3PBg=="},null,{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.192665324Z","validator_address":"35C361D16DCB5F5C4F06D0482B4DDE199B3F222F","validator_index":"3","signature":"vDt4+8M3n8aMGGLVzJyHrFSPHPyAkI7FR+usSxqYAxHRiQ1p26ennlpkskyq7Dwfbu+foNnJnZLi7NWT6r/ZCw=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.237437371Z","validator_address":"420AA1EA8A432C771C88744B841418108572E68B","validator_index":"4","signature":"uVWIaGvtFB2SeqnUPAC/tX0IcdlMrdEOS/j5YPmK7b70CVsiAMDUpqUQmSW5NGxlFGDnrfBZtti8G/hcWL8ZAw=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.19173035Z","validator_address":"548261410C22D5AA01C123EAC98DE4DDBEA2EF7E","validator_index":"5","signature":"c6u/Ksncg4G0PtmLw0sRA6jigZg2nlUgoAFXGuzNGGljSfx9pUnCEdH3CqV3+wsVejNA8rUac8psyEBeP4gHBw=="},null,{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.293357439Z","validator_address":"6175F13ED0076BB625DBA93FDED22204C18A9BD6","validator_index":"7","signature":"yG+0KSLSez3SOOVUNuJGzIYt5EwrEzp3xeorygjkjbDPTvzB3NBjMWYkg8nl6RMR4fzIYTIDY86lwFq+cwwPDA=="},null,{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.276770024Z","validator_address":"6B1B2F2C1E01B5DA7C014F0D62DAACE336FE52CE","validator_index":"9","signature":"/zMrnmWvg7EdOhSzZghbewHF0Nc/MGMqxPKOGcnWTk9zQlQ1botV4XI3QcYUrUnEIhThwP1HPW0w6MRJWefaBg=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.306167574Z","validator_address":"7A503503558C8281B3D00455DDB3497B958F23DA","validator_index":"10","signature":"1DTwrV11U/NB6L+4u35NK6btFr48ua5xh5PB5hMld4xP5jz+k/pkk36rt0cEwiSmi0HgXHZq79Q83mbxRzWsDw=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.24689756Z","validator_address":"8E3DECB8BE9120A05FB0011E7386E738AC60F8AA","validator_index":"11","signature":"8EdHKjFddOLSEDfuTTb0z5q72JS6OLd0pI8jmsAzfBLju8m3WUWEDyzj3UFrNsjjM5aoM+9GfgZMQ3B3TwR3Bw=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.141045325Z","validator_address":"91E1C69B68B2E8CD140E2A1D86F0CB31ACAE2113","validator_index":"12","signature":"m4gPaN3/tlm8A6uXMjGoFZfr6xxaV0bdR30pDt77UTEvjW/uk+HPTmhDfcssS5LR1ARjbCIAHtfdCgERmL5kBQ=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.191855015Z","validator_address":"9BF50513FD7EF408D0D677E4B1E540238AEDDB3B","validator_index":"13","signature":"sJ80d9+KEmWRaGZ+bbtOKzB2Jqybd8xRNsu4xCAHrUyhOqY7hYbqE1Uyx1cGRrC5OwAUn24kdDFJdN1NEg9lCQ=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.258319768Z","validator_address":"A1BF7521FBED60C6A61A0B3B3BA84D2E37D8EAEE","validator_index":"14","signature":"wat9HjXVECkdUOiVaOtNROxZnp/475dQIudif8fdIbWtfG7DWzL8DMYAbPJIKt+1y+SQYkFtwWX6A8uGmR2VCQ=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.229297379Z","validator_address":"AF1E3E818AF442E99F4C4BFABB5CE5B772DBF6A3","validator_index":"15","signature":"3SaySV96f3JqdsEd0N/HMEC0X3eiyKcXEVIfzbJobaTz1lu5p1OyR0ND3tkmQdCB4WmZmjgCalk4+CdoK808DQ=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.256895886Z","validator_address":"CB0C1B8373C74BA220FA146AF3683F9BDDBD002A","validator_index":"16","signature":"9xc1lgXcFI9hUbnCsP3ndXGFOctwXcnieY4v67a4fOU5fygRJb0n4H3tnF+GlScDrwkpOg+3WQWs5vTu+YhsCw=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.233387745Z","validator_address":"E40D07FA88FDF04F490F46E55F49E7B6443D5DBE","validator_index":"17","signature":"bOKRFhVNeDwCRBedWndE/3dTkPeRL34pcdyPVwuG5k+isohwhk+mMxZBV2YTFywuzf74tWg+VspQnGcdB/KJBA=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.200333219Z","validator_address":"E6F10F611145D6D7AB4521FCD35BEA6C61957112","validator_index":"18","signature":"jb53jrZ1lwMAEanT1BeT7VbpLtBRHROrSLOX/tBQpVEujfBCaxjn7UrhLKXmzbPo3zVG2s9LnAZ44heHBF8CBQ=="},{"type":2,"height":"11","round":"0","block_id":{"hash":"F256DF354F05B72353F171FCF1235B3330DEF977BB8A41055945D074230E64CA","parts":{"total":"1","hash":"E4647BAE9A9AFB97F86FBBA4738BE62BAAB3853530DC4F4079CE9689DC6F8590"}},"timestamp":"2020-01-09T05:16:20.312783069Z","validator_address":"ED7AFF397240437610647C42D7D821DF0592B6FE","validator_index":"19","signature":"45bfp1Oh/x0vctWvFwLhaTj53CTt97NIH9Uww8+wCamRZDkuWrkmVHkdkLjo7zLazzKRV51/E80BLToZaj13Cg=="},null]}}}`
			var resultBlock ctypes.ResultBlock
			json.Unmarshal([]byte(result), &resultBlock)

			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().Block(&height).Return(&resultBlock, nil)
			cli.SetClient(mockClient)

			resp, err := cli.QueryBlock(&height)
			So(err, ShouldBeNil)
			jsonBytes, err := json.Marshal(resp)
			So(err, ShouldBeNil)
			fmt.Println(string(jsonBytes))
		})

		Convey("Test return err", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var height int64 = 100
			errReturn := errors.New("Test return error")
			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().Block(&height).Return(nil, errReturn)
			cli.SetClient(mockClient)

			_, err := cli.QueryBlock(&height)
			So(err, ShouldBeError)
		})
	})
}

func TestQueryTx(t *testing.T) {
	Convey("TestQueryTx", t, func() {
		Convey("Test return valid output", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// get tx hash bytes
			txHash, err := hex.DecodeString("FA9A423C4D7C051800DB1E7A2F661406D95709D6AD932E99C8C4E517C8BCEE23")
			So(err, ShouldBeNil)

			resultTx := ctypes.ResultTx{
				Hash:   []byte("83ED00217AFDADB522EDD9D822422B1DFDF69C829C863F12524B017C84088D91"),
				Height: 3306876,
				Index:  0,
			}

			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().Tx(txHash, true).Return(&resultTx, nil)
			cli.SetClient(mockClient)

			resp, err := cli.QueryTx(txHash, true)
			So(err, ShouldBeNil)
			jsonBytes, err := json.Marshal(resp)
			So(err, ShouldBeNil)
			fmt.Println(string(jsonBytes))
		})

		Convey("Test return err", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// get tx hash bytes
			txHash, err := hex.DecodeString("FA9A423C4D7C051800DB1E7A2F661406D95709D6AD932E99C8C4E517C8BCEE23")
			So(err, ShouldBeNil)

			errReturn := errors.New("Test return error")
			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().Tx(txHash, true).Return(nil, errReturn)
			cli.SetClient(mockClient)

			_, err = cli.QueryTx(txHash, true)
			So(err, ShouldBeError)
		})
	})

}

func TestQueryCurrentValidators(t *testing.T) {
	Convey("TestQueryCurrentValidators", t, func() {
		Convey("Test return valid output", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			validatorsStr := `{"block_height":"3372617","validators":[{"address":"okchainvalcons1z2p8d228ucgfw45n8pnj5wg74235sgn2zrzerx","pub_key":"okchainvalconspub1zcjduepqyyy4g578mg08nxr6uf6vl9l0wzq2k25u86pmfuf9gyfuqt5gv3qqhhduct","proposer_priority":"-39960447","voting_power":"10000000"},{"address":"okchainvalcons1ryyd88rw4f5ade6km0euxz0ugweyn3z7msd2zv","pub_key":"okchainvalconspub1zcjduepqdtckh0697kdghss24hekfnvwpzhu4ac4dj5uygzef042jl7f6u2q0h499j","proposer_priority":"69956150","voting_power":"10000000"}]}`
			var currentValidators ctypes.ResultValidators
			amino.UnmarshalJSON([]byte(validatorsStr), &currentValidators)

			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().Validators(nil).Return(&currentValidators, nil)
			cli.SetClient(mockClient)

			resp, err := cli.QueryCurrentValidators()
			So(err, ShouldBeNil)
			jsonBytes, err := json.Marshal(resp)
			So(err, ShouldBeNil)
			fmt.Println(string(jsonBytes))
		})

		Convey("Test return err", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			errReturn := errors.New("Test return error")
			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().Validators(nil).Return(nil, errReturn)
			cli.SetClient(mockClient)

			_, err := cli.QueryCurrentValidators()
			So(err, ShouldBeError)
		})
	})

}

func TestQueryProposals(t *testing.T) {
	Convey("TestQueryProposals", t, func() {
		Convey("Test return valid output", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Be careful, the type here should be same with the type registered by the cli.cdc
			proposalsInfo := `[{"type":"okchain/gov/TextProposal","value":{"BasicProposal":{"proposal_id":"1","title":"test","description":"test","proposal_type":"Text","proposal_status":"DepositPeriod","tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-07-29T03:03:41.765548835Z","deposit_end_time":"2019-07-30T03:03:41.765548835Z","total_deposit":[{"denom":"okt","amount":"80.00000000"}],"voting_start_time":"0001-01-01T00:00:00Z","voting_end_time":"0001-01-01T00:00:00Z"}}},{"type":"okchain/gov/TextProposal","value":{"BasicProposal":{"proposal_id":"2","title":"test","description":"test","proposal_type":"Text","proposal_status":"DepositPeriod","tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-07-29T03:05:59.446787436Z","deposit_end_time":"2019-07-30T03:05:59.446787436Z","total_deposit":[{"denom":"okt","amount":"80.00000000"}],"voting_start_time":"0001-01-01T00:00:00Z","voting_end_time":"0001-01-01T00:00:00Z"}}},{"type":"okchain/gov/ParameterProposal","value":{"BasicProposal":{"proposal_id":"3","title":"Change gov/MinDeposit","description":"","proposal_type":"ParameterChange","proposal_status":"DepositPeriod","tally_result":{"yes":"0","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-07-29T03:18:21.761155019Z","deposit_end_time":"2019-07-30T03:18:21.761155019Z","total_deposit":[{"denom":"okt","amount":"60.00000000"}],"voting_start_time":"0001-01-01T00:00:00Z","voting_end_time":"0001-01-01T00:00:00Z"},"params":[{"subspace":"gov","key":"MinDeposit","value":"1000okt"}],"height":"1000"}}]`
			// If necessary, you can uncomment the next three line to check it
			//var rawProposals sdktypes.Proposals
			//cli.GetCdc().UnmarshalJSON([]byte(proposalsInfo), &rawProposals)
			//fmt.Println(rawProposals)

			var result ctypes.ResultABCIQuery
			result.Response.Code = 0
			result.Response.Value = []byte(proposalsInfo)

			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
			cli.SetClient(mockClient)

			resp, err := cli.QueryProposals()
			So(err, ShouldBeNil)
			jsonBytes, err := json.Marshal(resp)
			So(err, ShouldBeNil)
			fmt.Println(string(jsonBytes))
		})

		Convey("Test return err", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			errReturn := errors.New("Test return error")
			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errReturn)
			cli.SetClient(mockClient)

			_, err := cli.QueryProposals()
			So(err, ShouldBeError)
		})
	})
}

func TestQueryProposalByID(t *testing.T) {
	Convey("TestQueryProposalByID", t, func() {
		Convey("Test return valid output", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			proposalInfo := `{"type":"okchain/gov/DexListProposal","value":{"BasicProposal":{"proposal_id":"1","title":"list bcoin-7a4/okt","description":"","proposal_type":"DexList","proposal_status":"Passed","tally_result":{"yes":"100000000","abstain":"0","no":"0","no_with_veto":"0"},"submit_time":"2019-07-29T03:25:36.759218374Z","deposit_end_time":"2019-07-30T03:25:36.759218374Z","total_deposit":[{"denom":"okt","amount":"21000.00000000"}],"voting_start_time":"2019-07-29T03:31:43.90917706Z","voting_end_time":"2019-08-01T03:31:43.90917706Z"},"proposer":"okchain10q0rk5qnyag7wfvvt7rtphlw589m7frsmyq4ya","list_asset":"bcoin-a69","quote_asset":"okt","init_price":"2500.25000000","block_height":"0","max_price_digit":"4","max_size_digit":"4","min_trade_size":"0.001","dex_list_start_time":"2019-07-29T03:36:57.647117542Z","dex_list_end_time":"2019-07-30T03:36:57.647117542Z"}}`
			var result ctypes.ResultABCIQuery
			result.Response.Code = 0
			result.Response.Value = []byte(proposalInfo)

			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
			cli.SetClient(mockClient)

			resp, err := cli.QueryProposalByID(1)
			So(err, ShouldBeNil)
			jsonBytes, err := json.Marshal(resp)
			So(err, ShouldBeNil)
			fmt.Println(string(jsonBytes))
		})

		Convey("Test return err", func() {
			cli := client.NewClient("rpcUrl")
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			errReturn := errors.New("Test return error")
			mockClient := NewMockClient(ctrl)
			mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errReturn)
			cli.SetClient(mockClient)

			_, err := cli.QueryProposalByID(1)
			So(err, ShouldBeError)
		})
	})

}
