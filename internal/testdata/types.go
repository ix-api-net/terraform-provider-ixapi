package testdata

// NewOptString creates an optional string with a value
func NewOptString(val string) *string {
	return &val
}

// NewOptBool creates an optional bool with a value
func NewOptBool(val bool) *bool {
	return &val
}

// NewOptInt creates an optional integer with a value
func NewOptInt(val int) *int {
	return &val
}
