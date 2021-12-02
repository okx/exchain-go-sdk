package utils

import (
	"testing"
	"time"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	stakingtypes "github.com/okex/exchain/x/staking/types"
	"github.com/stretchr/testify/require"
)

const (
	valAddrStr = "exvaloper1qwuag8gx408m9ej038vzx50ntt0x4yrq8qwdtq"
)

func TestParseValAddresses(t *testing.T) {
	valAddrsStr := []string{valAddrStr}
	valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
	require.NoError(t, err)

	valAddrs, err := ParseValAddresses(valAddrsStr)
	require.NoError(t, err)
	require.Equal(t, 1, len(valAddrs))
	require.Equal(t, valAddr, valAddrs[0])

	// bad val address
	valAddrsStr = append(valAddrsStr, valAddrStr[1:])
	_, err = ParseValAddresses(valAddrsStr)
	require.Error(t, err)
}

func TestConvertToDelegatorResponse(t *testing.T) {
	expectedAccAddr, err := sdk.AccAddressFromBech32(defaultAddr)
	require.NoError(t, err)
	expectedValAddrs, err := ParseValAddresses([]string{valAddrStr})
	require.NoError(t, err)
	expectedDec := sdk.NewDec(1024)
	expectedTime := time.Now()
	expectedDelegator := stakingtypes.Delegator{
		DelegatorAddress:     expectedAccAddr,
		ValidatorAddresses:   expectedValAddrs,
		Shares:               expectedDec,
		IsProxy:              true,
		TotalDelegatedTokens: expectedDec,
		ProxyAddress:         expectedAccAddr,
	}

	expectedUndelegationInfo := stakingtypes.NewUndelegationInfo(expectedAccAddr, expectedDec, expectedTime)
	delResp := ConvertToDelegatorResponse(expectedDelegator, expectedUndelegationInfo)
	require.True(t, expectedAccAddr.Equals(delResp.DelegatorAddress))
	require.Equal(t, 1, len(delResp.ValidatorAddresses))
	require.True(t, expectedValAddrs[0].Equals(delResp.ValidatorAddresses[0]))
	require.True(t, expectedDec.Equal(delResp.Shares))
	require.True(t, expectedDec.Equal(delResp.UnbondedTokens))
	require.True(t, expectedTime.Equal(delResp.CompletionTime))
	require.True(t, delResp.IsProxy)
	require.True(t, expectedDec.Equal(delResp.TotalDelegatedTokens))
	require.True(t, expectedAccAddr.Equals(delResp.ProxyAddress))
}
