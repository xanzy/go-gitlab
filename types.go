package gitlab

import (
	"encoding/json"
)

// This file contains types to help dealing with GitLab API.

// BoolValue is a boolean value with advanced json unmarshaling features.
type BoolValue bool

// UnmarshalJSON implements json.Unmarshaler.
func (t *BoolValue) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case `"1"`:
		*t = true
		return nil
	case `"0"`:
		*t = false
		return nil
	default:
		var v bool
		err := json.Unmarshal(b, &v)
		*t = BoolValue(v)
		return err
	}
}
