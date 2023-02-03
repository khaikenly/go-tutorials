package main

import "fmt"

func primaryDiagonal(arr []int, col, m int) int {
	r := arr[col]
	for i := 1; i < len(arr); i++ {
		j := i + m
		if j < len(arr) {
			r += arr[j]
			i = j
		}
	}
	return r
}

func secondaryDiagonal(arr []int, col, m int) int {
	j := m - 1
	r := arr[j]
	z := 0
	for i := j; i < len(arr); z++ {
		if i <= len(arr)-m {
			i += j
			r += arr[i]
		}
	}
	return r
}

func main() {
	arrM := []int{11, 2, 4, 4, 5, 6, 10, 8, -12}
	col := 0
	m := 3
	primaryResult := primaryDiagonal(arrM, col, m)
	secondaryResult := secondaryDiagonal(arrM, col, m)
	fmt.Println(primaryResult, secondaryResult)
}
