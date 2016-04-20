package phonenumber

import (
	"encoding/json"
)

// Error represents errors when handling phone numbers
type Error string

// Error implements error interface
func (e *Error) Error() string {
	return string(*e)
}

// NewError constructs a new error
func NewError(msg string) *Error {
	err := Error(msg)
	return &err
}

// MarshalJSON implements json.Marshaler
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(*e)
}

// UnmarshalJSON implements json.Unmarshaler
func (e *Error) UnmarshalJSON(data []byte) error {
	var errorStr string
	err := json.Unmarshal(data, &errorStr)
	if err != nil {
		return err
	}
	if errorStr != "" {
		*e = Error(errorStr)
	}
	return nil
}
