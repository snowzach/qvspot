package qvspot

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Extra json.RawMessage
type Attr map[string]*StringList
type AttrNum map[string]float64

// MarshalJSON for Raw fields is represented as base64
func (e Extra) MarshalJSON() ([]byte, error) {

	return []byte(e), nil

}

// UnmarshalJSON for Raw fields is parsed as base64
func (e *Extra) UnmarshalJSON(in []byte) error {

	*e = make([]byte, len(in))
	copy(*e, in)
	return nil

}

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

// Scan implements the sql.Scanner interface
func (a *Attr) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("could not parse field type %T", src)
	}
	return json.Unmarshal(b, &a)
}

// Value implements the driver.Valuer interface
func (a Attr) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan implements the sql.Scanner interface
func (a *AttrNum) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("could not parse field type %T", src)
	}
	return json.Unmarshal(b, &a)
}

// Value implements the driver.Valuer interface
func (a AttrNum) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan implements the sql.Scanner interface
func (p *Position) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("could not parse field type %T", src)
	}
	return json.Unmarshal(b, &p)
}

// Value implements the driver.Valuer interface
func (p *Position) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan implements the sql.Scanner interface
func (e *Extra) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("could not parse field type %T", src)
	}
	*e = make([]byte, len(b))
	copy(*e, b)
	return nil
}

// Value implements the driver.Valuer interface
func (e Extra) Value() (driver.Value, error) {
	if len([]byte(e)) == 0 {
		return "{}", nil
	}
	return []byte(e), nil
}
