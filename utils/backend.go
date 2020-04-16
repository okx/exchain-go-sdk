package utils

import (
	"bytes"
	"encoding/json"
	bkdtypes "github.com/okex/okchain-go-sdk/module/backend/types"
)

// UnmarshalListResponse unmarshals the list response from bytes data
func UnmarshalListResponse(bz []byte, ptr interface{}) error {
	var lr bkdtypes.ListResponse
	if err := json.Unmarshal(bz, &lr); err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(lr.Data.Data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, ptr); err != nil {
		return err
	}
	return nil
}

// GetDataFromBaseResponse gets the detail data from the base response bytes
func GetDataFromBaseResponse(bytes []byte, ptr interface{}) error {
	dataBytes := getRawDataFromBaseResponse(bytes)
	if err := json.Unmarshal(dataBytes, ptr); err != nil {
		return err
	}
	return nil
}

func getRawDataFromBaseResponse(bz []byte) []byte {
	preIndex := bytes.Index(bz, []byte("data"))
	//sufIndex := bytes.LastIndex(bz, []byte("detail_msg"))
	return bz[preIndex+6 : len(bz)-1 ]
}
