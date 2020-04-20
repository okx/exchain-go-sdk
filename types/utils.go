package types

import (
	"encoding/json"
)

// MustSortJSON is like SortJSON but panic if an error occurs, e.g., if the passed JSON isn't valid
func MustSortJSON(toSortJSON []byte) []byte {
	js, err := SortJSON(toSortJSON)
	if err != nil {
		panic(err)
	}
	return js
}

// SortedJSON takes any JSON and returns it sorted by keys. Also, all white-spaces are removed
func SortJSON(toSortJSON []byte) ([]byte, error) {
	var c interface{}
	err := json.Unmarshal(toSortJSON, &c)
	if err != nil {
		return nil, err
	}
	js, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return js, nil
}
