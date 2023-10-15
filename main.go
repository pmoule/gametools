package main

import (
	"fmt"

	"github.com/pmoule/gametools/roller"
)

func main() {
	dice := []*roller.Die{roller.D2, roller.D3, roller.D4, roller.D6, roller.D8, roller.D10, roller.D12, roller.D20}

	for _, d := range dice {
		value := d.Roll()
		fmt.Printf("Rolled %s: %d\n", d.String(), value)
	}
}
