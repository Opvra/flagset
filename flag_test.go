package flagset

import "testing"

func TestFlagQueries(t *testing.T) {
	read := Flag(1 << 0)
	write := Flag(1 << 1)
	admin := Flag(1 << 2)

	f := read | admin

	if !f.Has(read) {
		t.Fatalf("expected Has(read) to be true")
	}
	if f.Has(write) {
		t.Fatalf("expected Has(write) to be false")
	}
	if !f.Any(read | write) {
		t.Fatalf("expected Any(read|write) to be true")
	}
	if f.Any(write) {
		t.Fatalf("expected Any(write) to be false")
	}
	if !f.All(admin) {
		t.Fatalf("expected All(admin) to be true")
	}
}

func TestFlagMutations(t *testing.T) {
	read := Flag(1 << 0)
	write := Flag(1 << 1)

	var f Flag
	f.Grant(read)
	if !f.Has(read) {
		t.Fatalf("expected Grant to set read")
	}
	f.Grant(write)
	if !f.Has(read | write) {
		t.Fatalf("expected Grant to set both bits")
	}
	f.Revoke(read)
	if f.Has(read) {
		t.Fatalf("expected Revoke to clear read")
	}
	f.Toggle(write)
	if f.Any(write) {
		t.Fatalf("expected Toggle to clear write")
	}
	f.Toggle(read)
	if !f.Has(read) {
		t.Fatalf("expected Toggle to set read")
	}
	f.Reset()
	if f != 0 {
		t.Fatalf("expected Reset to clear all bits")
	}
}
