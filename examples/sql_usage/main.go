package main

import (
	"fmt"
	"log"

	"github.com/Opvra/flagset"
)

const (
	PermRead  flagset.Flag = 1 << 0
	PermWrite flagset.Flag = 1 << 1
)

func main() {
	// Simulate reading an INTEGER column from SQL.
	var f flagset.Flag
	if err := f.Scan(int64(3)); err != nil {
		log.Fatal(err)
	}

	if f.Has(PermRead | PermWrite) {
		fmt.Println("read+write enabled")
	}

	// Simulate writing back to SQL.
	value, err := f.Value()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("db value: %v\n", value)
}
