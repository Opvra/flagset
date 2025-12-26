package flagset

import (
	"bytes"
	"encoding/json"
	"errors"
	"sync"
)

var (
	jsonRegistry    Registry
	jsonRegistrySet bool
	jsonRegistryMu  sync.Mutex
)

var errNoJSONRegistry = errors.New("flagset: JSON registry not set")
var errJSONRegistryAlreadySet = errors.New("flagset: JSON registry already set")
var errNilJSONRegistry = errors.New("flagset: JSON registry is nil")

// MarshalJSON encodes f as a JSON array of flag names.
func (f Flag) MarshalJSON() ([]byte, error) {
	if !jsonRegistrySet {
		return nil, errNoJSONRegistry
	}
	return json.Marshal(jsonRegistry.Names(f))
}

// UnmarshalJSON decodes a JSON array of flag names into f.
func (f *Flag) UnmarshalJSON(data []byte) error {
	if !jsonRegistrySet {
		return errNoJSONRegistry
	}
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		*f = 0
		return nil
	}

	var names []string
	if err := json.Unmarshal(data, &names); err != nil {
		return err
	}
	flag, err := jsonRegistry.Parse(names)
	if err != nil {
		return err
	}
	*f = flag
	return nil
}

// SetJSONRegistry sets the global JSON registry once at process startup.
func SetJSONRegistry(r Registry) error {
	if r == nil {
		return errNilJSONRegistry
	}
	jsonRegistryMu.Lock()
	defer jsonRegistryMu.Unlock()
	if jsonRegistrySet {
		return errJSONRegistryAlreadySet
	}
	jsonRegistry = r
	jsonRegistrySet = true
	return nil
}
