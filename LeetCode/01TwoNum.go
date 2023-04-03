package main

import "fmt"

func addNum(nums []int, x int) []int {
	// numCludes -> num | position
	numCludes := make(map[int]int)
	for i, num := range nums {
		// If key is in m, ok is true. If not, ok is false.
		// If key is not in the map, then elem is the zero value for the map's element type.
		j, ok := numCludes[x-num]
		if ok {
			fmt.Println([]int{j, i})
		}
		numCludes[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{1, 2, 3, 5, 6}
	target := 8
	result := addNum(nums, target)
	fmt.Println(result)
}
