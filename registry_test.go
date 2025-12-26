package flagset

import "testing"

func TestRegistryParse(t *testing.T) {
	r := Registry{
		"read":  Flag(1 << 0),
		"write": Flag(1 << 1),
		"admin": Flag(1 << 2),
	}

	flag, err := r.Parse([]string{"read", "admin"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !flag.Has(Flag(1<<0)|Flag(1<<2)) {
		t.Fatalf("expected parsed flag to include read and admin")
	}

	if _, err := r.Parse([]string{"missing"}); err == nil {
		t.Fatalf("expected error for unknown name")
	}
}

func TestRegistryMustParse(t *testing.T) {
	r := Registry{"read": Flag(1 << 0)}

	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic from MustParse")
		}
	}()

	_ = r.MustParse([]string{"missing"})
}

func TestRegistryNames(t *testing.T) {
	r := Registry{
		"write": Flag(1 << 1),
		"read":  Flag(1 << 0),
		"admin": Flag(1 << 2),
		"zero":  0,
	}

	flag := Flag(1<<0) | Flag(1<<2) | Flag(1<<10)
	names := r.Names(flag)
	if len(names) != 2 {
		t.Fatalf("expected 2 names, got %d", len(names))
	}
	if names[0] != "admin" || names[1] != "read" {
		t.Fatalf("expected sorted names [admin read], got %v", names)
	}
}
