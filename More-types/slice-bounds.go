package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)
	fmt.Println(cap(s))

	s = s[:2]
	fmt.Println(s) //[3 5] => [0:2]
	s = s[1:]
	fmt.Println(s) // [5] => [1:2] *get the high bound declare abrove
	fmt.Println(len(s), cap(s))

	s = s[1:2]
	fmt.Println(len(s), cap(s))

}
