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
	return fmt.Errorf("failed. ok client query error: %s", errMsg)
}
