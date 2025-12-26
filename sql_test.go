package flagset

import (
	"math"
	"testing"
)

func TestFlagValue(t *testing.T) {
	f := Flag(5)
	value, err := f.Value()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value.(int64) != 5 {
		t.Fatalf("expected 5, got %v", value)
	}

	tooLarge := Flag(math.MaxInt64) + 1
	if _, err := tooLarge.Value(); err == nil {
		t.Fatalf("expected overflow error")
	}
}

func TestFlagScan(t *testing.T) {
	cases := []struct {
		name  string
		input any
		want  Flag
	}{
		{"int64", int64(7), 7},
		{"int", int(9), 9},
		{"string", "12", 12},
		{"bytes", []byte("42"), 42},
		{"nil", nil, 0},
	}

	for _, tc := range cases {
		var f Flag
		if err := f.Scan(tc.input); err != nil {
			t.Fatalf("%s: unexpected error: %v", tc.name, err)
		}
		if f != tc.want {
			t.Fatalf("%s: expected %d, got %d", tc.name, tc.want, f)
		}
	}

	var f Flag
	if err := f.Scan(int64(-1)); err == nil {
		t.Fatalf("expected error for negative value")
	}
}
