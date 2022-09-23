package filter

import "testing"

func TestMissingFilter(t *testing.T) {
	a := "foo"

	if Missing(a, "fOO", true) {
		t.Error("should return false, because string is present")
	}
	if Missing(a, "", false) {
		t.Error("should return false, because filter is not present")
	}
	if !Missing(a, "bar", true) {
		t.Error("should return true, because filter does not match")
	}

	// Optional strings
	var opt *string = &a
	if Missing(opt, "foo", true) {
		t.Error("should return true, because opt is foo")
	}

	opt = nil
	if !Missing(opt, "foo", true) {
		t.Error("should not return true, because opt is nil")
	}
}
