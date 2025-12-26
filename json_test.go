package flagset

import (
	"encoding/json"
	"testing"
)

func resetJSONRegistryForTest() {
	jsonRegistryMu.Lock()
	defer jsonRegistryMu.Unlock()
	jsonRegistry = nil
	jsonRegistrySet = false
}

func TestFlagJSONRoundTrip(t *testing.T) {
	resetJSONRegistryForTest()
	r := Registry{
		"read":  Flag(1 << 0),
		"write": Flag(1 << 1),
		"admin": Flag(1 << 2),
	}
	if err := SetJSONRegistry(r); err != nil {
		t.Fatalf("unexpected registry error: %v", err)
	}
	defer resetJSONRegistryForTest()

	input := struct {
		Flags Flag `json:"flags"`
	}{Flags: Flag(1<<0) | Flag(1<<2) | Flag(1<<10)}

	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}

	if string(data) != `{"flags":["admin","read"]}` {
		t.Fatalf("unexpected json: %s", string(data))
	}

	var output struct {
		Flags Flag `json:"flags"`
	}
	if err := json.Unmarshal(data, &output); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}

	expected := Flag(1<<0) | Flag(1<<2)
	if output.Flags != expected {
		t.Fatalf("expected %d, got %d", expected, output.Flags)
	}
}

func TestFlagJSONUnknown(t *testing.T) {
	resetJSONRegistryForTest()
	r := Registry{"read": Flag(1 << 0)}
	if err := SetJSONRegistry(r); err != nil {
		t.Fatalf("unexpected registry error: %v", err)
	}
	defer resetJSONRegistryForTest()

	var output Flag
	if err := json.Unmarshal([]byte(`["write"]`), &output); err == nil {
		t.Fatalf("expected error for unknown name")
	}
}

func TestFlagJSONNoRegistry(t *testing.T) {
	resetJSONRegistryForTest()

	_, err := json.Marshal(Flag(1))
	if err == nil {
		t.Fatalf("expected error when registry is not set")
	}

	var output Flag
	if err := json.Unmarshal([]byte(`["read"]`), &output); err == nil {
		t.Fatalf("expected error when registry is not set")
	}
}

func TestSetJSONRegistryOnce(t *testing.T) {
	resetJSONRegistryForTest()
	defer resetJSONRegistryForTest()

	r := Registry{"read": Flag(1 << 0)}
	if err := SetJSONRegistry(r); err != nil {
		t.Fatalf("unexpected registry error: %v", err)
	}
	if err := SetJSONRegistry(r); err == nil {
		t.Fatalf("expected error on second SetJSONRegistry")
	}
}
