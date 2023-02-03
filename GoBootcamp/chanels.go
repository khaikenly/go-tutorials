package main

import "fmt"

//define the count() function
func count(num int, arr []int, ch chan int) {
	c := 0
	for _, a := range arr {
		if a == num {
			c++
		}
	}
	ch <- c
}

func main() {
	data := []int{12, 45, 88, 42, 0, 98, 102, 42, 77, 42, 1, 8, 7, 55, 4, 12, 87, 90, 42, 42, 11, 2, 6, 53, 90, 100, 4, 32, 8}

	var num int

	fmt.Print("Input num: ")
	fmt.Scanln(&num)

	ch1 := make(chan int)

	ch2 := make(chan int)

	go count(num, data[:len(data)/2], ch1)

	go count(num, data[len(data)/2:], ch2)

	//output the final result here
	resp_c1 := <-ch1
	resp_c2 := <-ch2
	fmt.Println(resp_c1 + resp_c2)
}
