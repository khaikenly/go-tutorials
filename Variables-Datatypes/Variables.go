package main

import "fmt"

func main() {
	// Declaring variables
	var i int
	var s string
	// Initializing Variables
	i = 20
	s = "Some String"
	fmt.Println(i)
	fmt.Println(s)

	// Creating and initializing variables
	var num1 int = 35
	var num2 int = 50
	fmt.Println(num1)
	fmt.Println(num2)

	//Short variable declaration
	r := 10
	a := "abc"
	fmt.Println(r, a)

	// Declaring multiple variables
	firstName, LastName := "Nguyen", "Khai"
	fmt.Println(firstName, LastName)

	// Variable Declaration Block
	var (
		h = "A"
		k = 2
	)

	fmt.Println(h, k)
}
