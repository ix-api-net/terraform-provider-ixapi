package ixapi

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDetailString_UnmarshalJSON_String(t *testing.T) {
	var fs DetailString
	err := json.Unmarshal([]byte(`"simple string"`), &fs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if fs.String() != "simple string" {
		t.Errorf("expected 'simple string', got '%s'", fs.String())
	}
}

func TestDetailString_UnmarshalJSON_Array(t *testing.T) {
	SetFlexibleParsing(true)
	t.Cleanup(func() { SetFlexibleParsing(false) })
	var fs DetailString
	err := json.Unmarshal([]byte(`["error one", "error two", "error three"]`), &fs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "error one; error two; error three"
	if fs.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, fs.String())
	}
}

func TestDetailString_UnmarshalJSON_MixedArray(t *testing.T) {
	SetFlexibleParsing(true)
	t.Cleanup(func() { SetFlexibleParsing(false) })
	var fs DetailString
	err := json.Unmarshal([]byte(`["string", 123, true]`), &fs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "string; 123; true"
	if fs.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, fs.String())
	}
}

func TestDetailString_UnmarshalJSON_Invalid(t *testing.T) {
	var fs DetailString
	err := json.Unmarshal([]byte(`{"invalid": "object"}`), &fs)
	if err == nil {
		t.Fatal("expected error for invalid JSON type")
	}
}

func TestApiTimestamp_UnmarshalJSON_RFC3339(t *testing.T) {
	var ft ApiTimestamp
	err := json.Unmarshal([]byte(`"2024-08-01T14:30:00Z"`), &ft)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := time.Date(2024, 8, 1, 14, 30, 0, 0, time.UTC)
	if !ft.Time.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, ft.Time)
	}
}

func TestApiTimestamp_UnmarshalJSON_DateOnly(t *testing.T) {
	SetFlexibleParsing(true)
	t.Cleanup(func() { SetFlexibleParsing(false) })
	var ft ApiTimestamp
	err := json.Unmarshal([]byte(`"2024-08-01"`), &ft)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)
	if !ft.Time.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, ft.Time)
	}
}

func TestApiTimestamp_UnmarshalJSON_Null(t *testing.T) {
	var ft ApiTimestamp
	err := json.Unmarshal([]byte(`null`), &ft)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ft.Time.IsZero() {
		t.Errorf("expected zero time, got %v", ft.Time)
	}
}

func TestApiTimestamp_UnmarshalJSON_EmptyString(t *testing.T) {
	var ft ApiTimestamp
	err := json.Unmarshal([]byte(`""`), &ft)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ft.Time.IsZero() {
		t.Errorf("expected zero time, got %v", ft.Time)
	}
}

func TestApiTimestamp_MarshalJSON(t *testing.T) {
	ft := ApiTimestamp{Time: time.Date(2024, 8, 1, 14, 30, 0, 0, time.UTC)}
	data, err := json.Marshal(ft)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `"2024-08-01T14:30:00Z"`
	if string(data) != expected {
		t.Errorf("expected %s, got %s", expected, string(data))
	}
}

func TestApiTimestamp_MarshalJSON_Zero(t *testing.T) {
	ft := ApiTimestamp{}
	data, err := json.Marshal(ft)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := `null`
	if string(data) != expected {
		t.Errorf("expected %s, got %s", expected, string(data))
	}
}

func TestProblemResponse_Error_WithDetailString(t *testing.T) {
	pr := ProblemResponse{
		Title:  "Validation Error",
		Status: 400,
		Detail: DetailString("Invalid input provided"),
	}
	expected := "Validation Error (400), Invalid input provided"
	if pr.Error() != expected {
		t.Errorf("expected '%s', got '%s'", expected, pr.Error())
	}
}

func TestProblemResponse_UnmarshalJSON_DetailAsString(t *testing.T) {
	jsonData := `{"title":"Error","status":400,"detail":"Simple error message"}`
	var pr ProblemResponse
	err := json.Unmarshal([]byte(jsonData), &pr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pr.Detail.String() != "Simple error message" {
		t.Errorf("expected 'Simple error message', got '%s'", pr.Detail.String())
	}
}

func TestProblemResponse_UnmarshalJSON_DetailAsArray(t *testing.T) {
	SetFlexibleParsing(true)
	t.Cleanup(func() { SetFlexibleParsing(false) })
	jsonData := `{"title":"Error","status":400,"detail":["Error 1","Error 2"]}`
	var pr ProblemResponse
	err := json.Unmarshal([]byte(jsonData), &pr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "Error 1; Error 2"
	if pr.Detail.String() != expected {
		t.Errorf("expected '%s', got '%s'", expected, pr.Detail.String())
	}
}
