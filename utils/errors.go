package utils

import "fmt"

// ErrMarshalJSON returns an error when it failed in marshaling JSON
func ErrMarshalJSON(errMsg string) error {
	return fmt.Errorf("failed. marshal JSON error: %s", errMsg)
}

// ErrUnmarshalJSON returns an error when it failed in unmarshaling JSON
func ErrUnmarshalJSON(errMsg string) error {
	return fmt.Errorf("failed. unmarshal JSON error: %s", errMsg)
}

// ErrClientQuery returns an error when client failed in query
func ErrClientQuery(errMsg string) error {
	return fmt.Errorf("failed. client query error: %s", errMsg)
}

// ErrFilterDataFromBaseResponse returns an error when it failed to filter data from backend base response
func ErrFilterDataFromBaseResponse(kind, errMsg string) error {
	return fmt.Errorf("failed. filter %s data from base response error: %s", kind, errMsg)
}

// ErrFilterDataFromListResponse returns an error when it failed to filter data from backend list response
func ErrFilterDataFromListResponse(kind, errMsg string) error {
	return fmt.Errorf("failed. filter %s data from list response error: %s", kind, errMsg)
}
