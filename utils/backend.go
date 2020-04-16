package utils

import (
	"bytes"
	"encoding/json"
)

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
