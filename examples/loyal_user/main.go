package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Opvra/flagset"
)

const (
	FlagLoyal   flagset.Flag = 1 << 0
	FlagRegular flagset.Flag = 1 << 1
)

type UserPayload struct {
	UserID string       `json:"user_id"`
	Flags  flagset.Flag `json:"flags"`
}

func main() {
	err := flagset.SetJSONRegistry(flagset.Registry{
		"loyal":   FlagLoyal,
		"regular": FlagRegular,
	})
	if err != nil {
		log.Fatal(err)
	}

	raw := []byte(`{"user_id":"u-123","flags":["loyal"]}`)

	var p UserPayload
	if err := json.Unmarshal(raw, &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(renderLanding(p.Flags))
}

func renderLanding(flags flagset.Flag) string {
	switch {
	case flags.Has(FlagLoyal):
		return `<div class="hero">Thanks for being loyal!</div>`
	case flags.Has(FlagRegular):
		return `<div class="hero">Welcome back!</div>`
	default:
		return `<div class="hero">Hello!</div>`
	}
}
