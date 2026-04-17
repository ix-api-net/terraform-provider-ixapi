package schemas

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGetTimeOpt_RFC3339(t *testing.T) {
	s := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"timestamp": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

	rd := s.Data(nil)
	_ = rd.Set("timestamp", "2024-08-01T14:30:00Z")

	res := ResourceData{ResourceData: rd}
	ts, err := res.GetTimeOpt("timestamp")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ts == nil {
		t.Fatal("expected non-nil time")
	}

	expected := time.Date(2024, 8, 1, 14, 30, 0, 0, time.UTC)
	if !ts.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, *ts)
	}
}

func TestGetTimeOpt_DateOnly(t *testing.T) {
	s := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"date": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

	rd := s.Data(nil)
	_ = rd.Set("date", "2024-08-01")

	res := ResourceData{ResourceData: rd}
	ts, err := res.GetTimeOpt("date")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ts == nil {
		t.Fatal("expected non-nil time")
	}

	expected := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)
	if !ts.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, *ts)
	}
}

func TestGetTimeOpt_NotSet(t *testing.T) {
	s := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"timestamp": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

	rd := s.Data(nil)

	res := ResourceData{ResourceData: rd}
	ts, err := res.GetTimeOpt("timestamp")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ts != nil {
		t.Errorf("expected nil, got %v", *ts)
	}
}

func TestGetTimeOpt_InvalidFormat(t *testing.T) {
	s := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"timestamp": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

	rd := s.Data(nil)
	_ = rd.Set("timestamp", "not-a-date")

	res := ResourceData{ResourceData: rd}
	_, err := res.GetTimeOpt("timestamp")
	if err == nil {
		t.Error("expected error for invalid date format")
	}
}
