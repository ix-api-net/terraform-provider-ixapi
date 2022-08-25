package provider

import "testing"

func TestProvider(t *testing.T) {
	p := New("test")
	if p == nil {
		t.Fatal("could not create provider instance")
	}
}
