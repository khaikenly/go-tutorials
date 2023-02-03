package main

import "fmt"

func main() {
	var m = make(map[string]int)

	m["a"] = 123
	fmt.Println(m["a"])

	m["a"] = 444
	fmt.Println(m["a"])

	delete(m, "a")
	fmt.Println(m["a"])

	elem, ok := m["a"]
	fmt.Println(elem, ok)
	// If key is in m, ok is true. If not, ok is false.
	// If key is not in the map, then elem is the zero value for the map's element type.
}
