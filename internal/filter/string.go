package filter

import "strings"

// String compares strings, case insensitive.
// Returns true if the strings do not match, so the
// result can be skipped.
func String(a string, b any, apply bool) bool {
	return apply && (strings.ToLower(a) != strings.ToLower(b.(string)))
}
