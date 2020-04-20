package types

import (
	"fmt"
	"log"
	"sort"
	"strings"

	abci "github.com/tendermint/tendermint/abci/types"
)

type (
	// Event is a type alias for an ABCI Event
	Event abci.Event

	// Attribute defines an attribute wrapper where the key and value are
	// strings instead of raw bytes.
	Attribute struct {
		Key   string `json:"key"`
		Value string `json:"value,omitempty"`
	}

	// Events defines a slice of Event objects
	Events []Event
)

// String returns a human readable string representation of Attribute
func (a Attribute) String() string {
	return fmt.Sprintf("%s: %s", a.Key, a.Value)
}

type (
	// StringEvent defines en Event object wrapper where all the attributes
	// contain key/value pairs that are strings instead of raw bytes
	StringEvent struct {
		Type       string      `json:"type,omitempty"`
		Attributes []Attribute `json:"attributes,omitempty"`
	}

	// StringEvents defines a slice of StringEvents objects
	StringEvents []StringEvent
)

// String returns a human readable string representation of StringEvents
func (se StringEvents) String() string {
	var sb strings.Builder

	for _, e := range se {
		if _, err := sb.WriteString(fmt.Sprintf("\t\t- %s\n", e.Type)); err != nil {
			log.Println(err)
		}

		for _, attr := range e.Attributes {
			if _, err := sb.WriteString(fmt.Sprintf("\t\t\t- %s\n", attr.String())); err != nil {
				log.Println(err)
			}
		}
	}

	return strings.TrimRight(sb.String(), "\n")
}

// Flatten returns a flattened version of StringEvents by grouping all attributes per unique event type
func (se StringEvents) Flatten() StringEvents {
	flatEvents := make(map[string][]Attribute)

	for _, e := range se {
		flatEvents[e.Type] = append(flatEvents[e.Type], e.Attributes...)
	}

	var (
		res  StringEvents
		keys []string
	)

	for ty := range flatEvents {
		keys = append(keys, ty)
	}

	sort.Strings(keys)
	for _, ty := range keys {
		res = append(res, StringEvent{Type: ty, Attributes: flatEvents[ty]})
	}

	return res
}

// StringifyEvent converts an Event object to a StringEvent object
func StringifyEvent(e abci.Event) StringEvent {
	res := StringEvent{Type: e.Type}

	for _, attr := range e.Attributes {
		res.Attributes = append(
			res.Attributes,
			Attribute{string(attr.Key), string(attr.Value)},
		)
	}

	return res
}

// StringifyEvents converts a slice of Event objects into a slice of StringEvent objects
func StringifyEvents(events []abci.Event) StringEvents {
	var res StringEvents

	for _, e := range events {
		res = append(res, StringifyEvent(e))
	}

	return res.Flatten()
}
