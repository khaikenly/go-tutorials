package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
	F float64
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S, t.F)
}

func main() {
	var i I = T{S: "Hello", F: math.Pi}
	i.M()
}
