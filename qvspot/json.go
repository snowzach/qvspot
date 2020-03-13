package qvspot

import (
	"encoding/json"
)

// MarshalJSON for Raw fields is represented as base64
func (sl *StringList) MarshalJSON() ([]byte, error) {

	if sl == nil {
		return json.Marshal(sl)
	}

	return json.Marshal(sl.List)

}

// UnmarshalJSON for Raw fields is parsed as base64
func (sl *StringList) UnmarshalJSON(in []byte) error {

	if sl == nil {
		sl = new(StringList)
	}

	return json.Unmarshal(in, &sl.List)
}
