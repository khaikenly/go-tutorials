package main

import "fmt"

func fibonacci_number(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci_number(n-1) + fibonacci_number(n-2)
}

func main() {
	fmt.Println(fibonacci_number(13))
}
