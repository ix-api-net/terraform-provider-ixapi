package ixapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClientSetsUserAgent(t *testing.T) {
	var capturedUA string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedUA = r.Header.Get("User-Agent")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode([]CloudRouter{}); err != nil {
			t.Fatalf("failed to encode response: %v", err)
		}
	}))
	defer server.Close()

	client := NewClient(server.URL, "1.2.3")

	_, err := client.DecixCloudRoutersList(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "terraform-provider-ixapi/1.2.3"
	if capturedUA != expected {
		t.Errorf("expected User-Agent %q, got %q", expected, capturedUA)
	}
}
