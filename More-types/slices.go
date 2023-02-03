package main

import "fmt"

func main() {
	var s []int
	s = []int{1, 2, 2, 2, 2, 2, 3, 4, 5, 5}
	fmt.Println(s)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s1 []int = primes[1:4]
	fmt.Println(s1)
}
