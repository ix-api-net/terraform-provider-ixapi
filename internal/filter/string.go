package filter

import "strings"

// Missing compares values, strings case insensitive.
// Returns true if the values do not match, so the
// result can be skipped.
func Missing(a any, b any, apply bool) bool {
	if !apply {
		return false
	}

	var v1, v2 string
	switch v := a.(type) {
	case string:
		v1 = v
	case *string:
		if v == nil {
			return true
		}
		v1 = *v
	}
	switch v := b.(type) {
	case string:
		v2 = v
	case *string:
		if v == nil {
			return true
		}
		v2 = *v
	}
	return strings.ToLower(v1) != strings.ToLower(v2)
}
