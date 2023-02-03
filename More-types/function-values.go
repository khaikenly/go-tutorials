package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot, 4, 5))
	fmt.Println(compute(math.Pow, 6, 7))
}
