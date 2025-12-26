# github.com/Opvra/flagset

Small, dependency-free flag set utilities for Go. This package provides a compact, stable bitmask type and a registry for name <-> bit mapping.

## Why flag sets

- Compact: a single `uint64` fits in caches, tokens, DB rows, and network payloads.
- Fast: bitwise checks are constant time.
- Stable: old bits keep their meaning, new bits append.

## Why not a struct of bools

- Structs do not serialize compactly across boundaries.
- Mapping to SQL/JSON requires custom logic anyway.
- Bit masks enforce backward-compatible growth with minimal surface area.

## Core type

```go
// Each bit is an independent capability or state.
type Flag uint64
```

## Permission (RBAC) example

```go
const (
	PermRead  flagset.Flag = 1 << 0
	PermWrite flagset.Flag = 1 << 1
	PermAdmin flagset.Flag = 1 << 2
)

var perms flagset.Flag
perms.Grant(PermRead | PermWrite)
if perms.Has(PermRead) {
	// allow
}
```

## Feature flag example

```go
const (
	FeatureSearch flagset.Flag = 1 << 0
	FeatureBetaUI flagset.Flag = 1 << 1
)

flags := FeatureSearch
if flags.Any(FeatureBetaUI) {
	// enable beta UI
}
```

## Config/state flag example

```go
const (
	StateWarm flagset.Flag = 1 << 0
	StateBusy flagset.Flag = 1 << 1
)

var state flagset.Flag
state.Grant(StateWarm)
state.Toggle(StateBusy)
```

## JSON usage

JSON uses flag names via an explicit registry. Set it once at startup.

```go
r := flagset.Registry{
	"read":  1 << 0,
	"write": 1 << 1,
	"admin": 1 << 2,
}

_ = flagset.SetJSONRegistry(r)

payload := struct {
	Flags flagset.Flag `json:"flags"`
}{
	Flags: r.MustParse([]string{"read", "admin"}),
}

b, _ := json.Marshal(payload)
// {"flags":["admin","read"]}
```

## SQL usage

Flags implement `sql.Scanner` and `driver.Valuer` for INTEGER columns.

```go
var f flagset.Flag
_ = f.Scan(int64(3))
value, _ := f.Value()
// value is int64(3)
```

## Registry usage

```go
r := flagset.Registry{
	"read":  1 << 0,
	"write": 1 << 1,
	"admin": 1 << 2,
}

flags, _ := r.Parse([]string{"read", "write"})
if flags.Has(r.MustParse([]string{"read"})) {
	// ok
}
```

## Examples

- `examples/loyal_user/main.go`
- `examples/sql_usage/main.go`
