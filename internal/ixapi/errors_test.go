package ixapi

import (
	"testing"
)

func TestIsErrAPIFault(t *testing.T) {
	var err error = &APIError{}
	if !IsErrAPIFault(err) {
		t.Error("unexpected error:", err)
	}
}

func TestIsErrNotFound(t *testing.T) {
	var err error = &NotFoundError{}
	if !IsErrNotFound(err) {
		t.Error("unexpected error:", err)
	}
}

func TestIsErrAuthenticationFailed(t *testing.T) {
	var err error = &AuthenticationError{}
	if !IsErrAuthenticationFailed(err) {
		t.Error("unexpected error:", err)
	}
}

func TestIsErrPermissionDenied(t *testing.T) {
	var err error = &PermissionError{}
	if !IsErrPermissionDenied(err) {
		t.Error("unexpected error:", err)
	}
}

func TestIsErrValidationFailed(t *testing.T) {
	var err error = &ValidationError{}
	if !IsErrValidationFailed(err) {
		t.Error("unexpected error:", err)
	}
}
