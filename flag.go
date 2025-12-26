package flagset

// Flag represents a set of bit flags.
type Flag uint64

// Has reports whether all bits in flag are set in f.
func (f Flag) Has(flag Flag) bool {
	return f&flag == flag
}

// Any reports whether any bit in flag is set in f.
func (f Flag) Any(flag Flag) bool {
	return f&flag != 0
}

// All is an alias for Has.
func (f Flag) All(flag Flag) bool {
	return f.Has(flag)
}

// Grant sets the given bits.
func (f *Flag) Grant(flag Flag) {
	*f |= flag
}

// Revoke clears the given bits.
func (f *Flag) Revoke(flag Flag) {
	*f &^= flag
}

// Toggle flips the given bits.
func (f *Flag) Toggle(flag Flag) {
	*f ^= flag
}

// Reset clears all bits.
func (f *Flag) Reset() {
	*f = 0
}
