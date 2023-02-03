package main

import "fmt"

type Bootcamp struct {
	Lat float64
	Lon float64
}

func newFunction() (*Bootcamp, *Bootcamp) {
	x := new(Bootcamp)
	y := &Bootcamp{}
	return x, y
}

func main() {
	x, y := newFunction()
	fmt.Println(*x == *y)
}
