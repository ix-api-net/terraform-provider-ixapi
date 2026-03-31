package ixapi

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync/atomic"
	"time"
)

// flexibleParsingEnabled controls whether lenient JSON parsing is active.
// It is set atomically so it is safe to call SetFlexibleParsing concurrently.
var flexibleParsingEnabled atomic.Bool

// SetFlexibleParsing enables or disables lenient JSON parsing for DE-CIX extension API responses.
func SetFlexibleParsing(enabled bool) {
	flexibleParsingEnabled.Store(enabled)
}

// FlexibleString is a string type used for fields that the DE-CIX extension API
// may return as either a JSON string or a JSON array of strings — an inconsistency
// in the extension API that does not exist in the core IX-API. When flexible parsing
// is disabled (the default), only a JSON string is accepted.
type FlexibleString string

// UnmarshalJSON implements json.Unmarshaler, accepting either a JSON string or a JSON array.
func (fs *FlexibleString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*fs = FlexibleString(str)
		return nil
	}
	if !flexibleParsingEnabled.Load() {
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

// UnmarshalJSON implements json.Unmarshaler. It tries formats in order:
//  1. time.RFC3339 — the standard IX-API format.
//  2. "2006-01-02 15:04:05.999999999-07:00" — a non-standard variant produced by the DE-CIX
//     extension API that uses a space separator instead of 'T' and an explicit numeric timezone
//     offset instead of 'Z'. Only attempted when flexible parsing is enabled.
//  3. "2006-01-02" — date-only values returned for fields such as decommission_at when no
//     time component is set. Only attempted when flexible parsing is enabled.
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
	if !flexibleParsingEnabled.Load() {
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
