package qvspot

import (
	"encoding/json"
)

// MarshalJSON for Raw fields is represented as base64
func (sl *StringList) MarshalJSON() ([]byte, error) {

	if sl == nil {
		return json.Marshal(sl)
	}

	if len(sl.List) == 0 {
		return []byte(""), nil
	}

	if len(sl.List) == 1 {
		return json.Marshal(sl.List[0])
	}

	return json.Marshal(sl.List)

}

// UnmarshalJSON for Raw fields is parsed as base64
func (sl *StringList) UnmarshalJSON(in []byte) error {

	if sl == nil {
		sl = new(StringList)
	}

	// If it's a string
	if len(in) > 0 && in[0] == '"' {
		var s string
		err := json.Unmarshal(in, &s)
		if err != nil {
			return err
		}
		sl.List = []string{s}
		return nil
	}

	return json.Unmarshal(in, &sl.List)
}
