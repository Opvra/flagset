package flagset

import (
	"fmt"
	"sort"
)

// Registry maps human-readable names to flag bits.
type Registry map[string]Flag

// Parse returns a Flag composed from the provided names.
func (r Registry) Parse(names []string) (Flag, error) {
	var out Flag
	for _, name := range names {
		flag, ok := r[name]
		if !ok {
			return 0, fmt.Errorf("flagset: unknown flag name %q", name)
		}
		out |= flag
	}
	return out, nil
}

// MustParse is like Parse but panics on error.
func (r Registry) MustParse(names []string) Flag {
	flag, err := r.Parse(names)
	if err != nil {
		panic(err)
	}
	return flag
}

// Names returns the sorted names for the bits set in f.
// Unknown bits are ignored.
func (r Registry) Names(f Flag) []string {
	if r == nil {
		return nil
	}

	out := make([]string, 0, len(r))
	for name, flag := range r {
		if flag == 0 {
			continue
		}
		if f.Has(flag) {
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}
