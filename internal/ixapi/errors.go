package ixapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

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
	props := ""
	plen := len(e.Properties) - 1
	for i, prop := range e.Properties {
		props += fmt.Sprintf("%s: %s", prop.Name, prop.Reason)
		if i < plen {
			props += ", "
		}
	}
	return fmt.Sprintf("%s %s",
		e.Title, props)
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
