package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	errMsg = "default error message"
	kind   = "default kind"
)

func TestErrors(t *testing.T) {
	require.True(t, strings.Contains(ErrMarshalJSON(errMsg).Error(), errMsg))
	require.True(t, strings.Contains(ErrUnmarshalJSON(errMsg).Error(), errMsg))
	require.True(t, strings.Contains(ErrClientQuery(errMsg).Error(), errMsg))

	errStr := ErrFilterDataFromBaseResponse(kind, errMsg).Error()
	require.True(t, strings.Contains(errStr, errMsg) && strings.Contains(errStr, kind))
	errStr = ErrFilterDataFromListResponse(kind, errMsg).Error()
	require.True(t, strings.Contains(errStr, errMsg) && strings.Contains(errStr, kind))
}
