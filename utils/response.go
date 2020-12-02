package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/okex/okexchain/x/common"
)

// UnmarshalListResponse unmarshals the list response from data bytes
func UnmarshalListResponse(bz []byte, ptr interface{}) error {
	var lr common.ListResponse
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
	dataBytes, err := getRawDataFromBaseResponse(bytes)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(dataBytes, ptr); err != nil {
		return err
	}
	return nil
}

func getRawDataFromBaseResponse(bz []byte) (data []byte, err error) {
	preIndex, bytesLen := bytes.Index(bz, []byte("data")), len(bz)
	if preIndex == -1 || preIndex+6 > bytesLen-1 {
		return data, errors.New("failed. invalid format of JSON data received")
	}
	return bz[preIndex+6 : len(bz)-1], err
}
