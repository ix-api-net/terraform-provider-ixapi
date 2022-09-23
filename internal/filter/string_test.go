package filter

import "testing"

func TestStringFilter(t *testing.T) {
	a := "foo"

	if String(a, "fOO", true) {
		t.Error("should return false, because string is present")
	}
	if String(a, "", false) {
		t.Error("should return false, because filter is not present")
	}
	if !String(a, "bar", true) {
		t.Error("should return true, because filter does not match")
	}
}
