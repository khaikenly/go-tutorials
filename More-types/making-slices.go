package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) //make a slice with length is 2
	copy(slice2, slice1)     // copy 1 to 2
	fmt.Println(slice1, slice2)
}
