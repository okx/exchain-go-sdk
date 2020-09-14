package utils

import (
	"encoding/json"
	"testing"

	backend "github.com/okex/okexchain-go-sdk/module/backend/types"
	"github.com/stretchr/testify/require"
)

// custom struct for test
type TestData struct {
	Int   int
	Str   string
	Float float64
	Bytes []byte
}

func TestUnmarshalListResponse(t *testing.T) {
	// data preparation
	listResp := backend.ListResponse{
		Data: backend.ListDataRes{
			Data: TestData{
				Int:   1,
				Str:   "default string",
				Float: 1.024,
				Bytes: []byte("default bytes"),
			},
		},
	}

	rawBytes, err := json.Marshal(listResp)
	require.NoError(t, err)

	var testData TestData
	err = UnmarshalListResponse(rawBytes, &testData)
	require.NoError(t, err)
	require.Equal(t, 1, testData.Int)
	require.Equal(t, "default string", testData.Str)
	require.Equal(t, 1.024, testData.Float)
	require.Equal(t, []byte("default bytes"), testData.Bytes)

	// bad JSON bytes
	// case 1
	listResp.Data.Data = "bad string not raw JSON bytes of TestData"
	rawBytes, err = json.Marshal(listResp)
	require.NoError(t, err)

	err = UnmarshalListResponse(rawBytes, &testData)
	require.Error(t, err)

	// case 2
	badListResp := "common string not the object of ListResponse"
	rawBytes, err = json.Marshal(badListResp)
	require.NoError(t, err)

	err = UnmarshalListResponse(rawBytes, &testData)
	require.Error(t, err)
}

func TestGetDataFromBaseResponse(t *testing.T) {
	// data preparation
	baseResp := backend.BaseResponse{
		Data: TestData{
			Int:   1,
			Str:   "default string",
			Float: 1.024,
			Bytes: []byte("default bytes"),
		},
	}

	rawBytes, err := json.Marshal(baseResp)
	require.NoError(t, err)

	var testData TestData
	err = GetDataFromBaseResponse(rawBytes, &testData)
	require.NoError(t, err)
	require.Equal(t, 1, testData.Int)
	require.Equal(t, "default string", testData.Str)
	require.Equal(t, 1.024, testData.Float)
	require.Equal(t, []byte("default bytes"), testData.Bytes)

	// bad JSON bytes
	// case 1
	baseResp.Data = "common string not the object of BaseResponse"
	rawBytes, err = json.Marshal(baseResp)
	require.NoError(t, err)

	err = GetDataFromBaseResponse(rawBytes, &testData)
	require.Error(t, err)

	// case 2
	err = GetDataFromBaseResponse([]byte{}, &testData)
	require.Error(t, err)
}
