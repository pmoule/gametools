package roller

import (
	"fmt"
	"math/rand"
	"time"
)

var randSource rand.Source
var D2 *Die = &Die{2}
var D3 *Die = &Die{3}
var D4 *Die = &Die{4}
var D6 *Die = &Die{6}
var D8 *Die = &Die{8}
var D10 *Die = &Die{10}
var D12 *Die = &Die{12}
var D20 *Die = &Die{20}

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
}

type Die struct {
	Sides int
}

func (d *Die) Roll() int {
	rand := rand.New(randSource)
	value := rand.Intn(d.Sides) + 1

	return value
}

func (d *Die) String() string {
	return fmt.Sprintf("D%d", d.Sides)
}
