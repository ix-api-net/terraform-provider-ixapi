package ptr

import (
	"testing"
)

func TestOf(t *testing.T) {
	v := 42
	p := Of(v)
	if p == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *p != v {
		t.Fatalf("expected %d, got %d", v, *p)
	}
}

func TestOfIndependentCopy(t *testing.T) {
	v := 42
	p := Of(v)
	v = 99
	if *p != 42 {
		t.Fatal("pointer should hold an independent copy of the value")
	}
}

func TestOfString(t *testing.T) {
	s := "hello"
	p := Of(s)
	if *p != s {
		t.Fatalf("expected %q, got %q", s, *p)
	}
}
