package ixapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var flexibleParsingEnabled bool

// SetFlexibleParsing enables or disables lenient JSON parsing for DE-CIX extension API responses.
func SetFlexibleParsing(enabled bool) {
	flexibleParsingEnabled = enabled
}

// FlexibleString is a string type that can unmarshal both JSON strings and JSON arrays into a single string value.
type FlexibleString string

// UnmarshalJSON implements json.Unmarshaler, accepting either a JSON string or a JSON array.
func (fs *FlexibleString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*fs = FlexibleString(str)
		return nil
	}
	if !flexibleParsingEnabled {
		return fmt.Errorf("detail field must be a string, got: %s", string(data))
	}

	var arr []interface{}
	if err := json.Unmarshal(data, &arr); err == nil {
		parts := make([]string, len(arr))
		for i, v := range arr {
			parts[i] = fmt.Sprintf("%v", v)
		}
		*fs = FlexibleString(strings.Join(parts, "; "))
		return nil
	}

	return fmt.Errorf("detail field must be string or array, got: %s", string(data))
}

func (fs FlexibleString) String() string {
	return string(fs)
}

// FlexibleTime is a time.Time wrapper that can unmarshal multiple date formats returned by the DE-CIX extension API.
type FlexibleTime struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler, accepting RFC3339, datetime with offset, and date-only formats.
func (ft *FlexibleTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || len(data) == 0 {
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		ft.Time = t
		return nil
	}
	if !flexibleParsingEnabled {
		return err
	}

	t, err = time.Parse("2006-01-02 15:04:05.999999999-07:00", s)
	if err == nil {
		ft.Time = t
		return nil
	}

	t, err = time.Parse("2006-01-02", s)
	if err == nil {
		ft.Time = t
		return nil
	}

	return err
}

// MarshalJSON implements json.Marshaler, serializing to RFC3339 or null if zero.
func (ft FlexibleTime) MarshalJSON() ([]byte, error) {
	if ft.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + ft.Time.Format(time.RFC3339) + `"`), nil
}

// NewFlexibleTime wraps a *time.Time in a FlexibleTime, returning nil if the input is nil.
func NewFlexibleTime(t *time.Time) *FlexibleTime {
	if t == nil {
		return nil
	}
	return &FlexibleTime{Time: *t}
}

// ToTime returns the underlying *time.Time, or nil if the receiver is nil or zero.
func (ft *FlexibleTime) ToTime() *time.Time {
	if ft == nil || ft.Time.IsZero() {
		return nil
	}
	t := ft.Time
	return &t
}

// Errors

// APIError is a generic api error
type APIError struct {
	ProblemResponse
}

// NotFoundError indicates that a resource was not found
type NotFoundError struct {
	ProblemResponse
}

// AuthenticationError indicates that the authentication
// was not successful.
type AuthenticationError struct {
	ProblemResponse
}

// PermissionError indicates that insufficient rights were
// given, when trying to access a resource.
type PermissionError struct {
	ProblemResponse
}

// ValidationError indicates that the validation of user data
// failed. The Properties attribute should contain
// a list of property names and reasons.
type ValidationError struct {
	ProblemResponse
	Properties []ValidationErrorProp `json:"properties"`
}

// ValidationErrorProp A failed validation
type ValidationErrorProp struct {
	// Name is a name
	Name string `json:"name,omitempty"`

	// Reason is a reason
	Reason json.RawMessage `json:"reason,omitempty"`
}

// Error implements the error interface
func (e ValidationError) Error() string {
	parts := []string{e.ProblemResponse.Error()}
	for _, prop := range e.Properties {
		parts = append(parts, fmt.Sprintf("%s: %s", prop.Name, prop.Reason))
	}
	return strings.Join(parts, ", ")
}

// Error Type Checking

// AsErrAPIFault tries to convert an error into
// an APIError, a generic API error.
func AsErrAPIFault(err error) *APIError {
	var into *APIError
	if err == nil {
		return nil
	}
	if errors.As(err, &into) {
		return into
	}
	return nil
}

// IsErrAPIFault checks if the error is an APIError
func IsErrAPIFault(err error) bool {
	return AsErrAPIFault(err) != nil
}

// AsErrNotFound tries to convert an error into a not found
// error and will return nil if not successful
func AsErrNotFound(err error) *NotFoundError {
	var into *NotFoundError
	if err == nil {
		return nil
	}
	if errors.As(err, &into) {
		return into
	}
	return nil
}

// IsErrNotFound checks if the error is a not found error
func IsErrNotFound(err error) bool {
	return AsErrNotFound(err) != nil
}

// AsErrAuthenticationFailed tries to convert the
// error into a AuthenticationError
func AsErrAuthenticationFailed(err error) *AuthenticationError {
	var into *AuthenticationError
	if err == nil {
		return nil
	}
	if errors.As(err, &into) {
		return into
	}
	return nil
}

// IsErrAuthenticationFailed checks if the error is an
// authentication error
func IsErrAuthenticationFailed(err error) bool {
	return AsErrAuthenticationFailed(err) != nil
}

// AsErrPermissionDenied tries to convert the error
// into a PermissionError
func AsErrPermissionDenied(err error) *PermissionError {
	var into *PermissionError
	if err == nil {
		return nil
	}
	if errors.As(err, &into) {
		return into
	}
	return nil
}

// IsErrPermissionDenied checks if the error is a
// PermissionError
func IsErrPermissionDenied(err error) bool {
	return AsErrPermissionDenied(err) != nil
}

// AsErrValidationFailed tries to convert an error into
// a ValidationError
func AsErrValidationFailed(err error) *ValidationError {
	var into *ValidationError
	if err == nil {
		return nil
	}
	if errors.As(err, &into) {
		return into
	}
	return nil
}

// IsErrValidationFailed checks if this is an validation error
func IsErrValidationFailed(err error) bool {
	return AsErrValidationFailed(err) != nil
}
