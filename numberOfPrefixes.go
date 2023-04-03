package main

import (
	"fmt"
	"strings"
)

func numberOfPrefixes(arr1 []string, arr2 []string) []int {
	result := make([]int, len(arr2))
	for i, str2 := range arr2 {
		for _, str1 := range arr1 {
			if strings.HasPrefix(str1, str2) {
				result[i] += 1
			}
		}
	}
	return result
}

func checkIsReduce(n int) bool {
	var k1 int
	var k2 int
	for n != 0 {
		k1 = n % 10
		n = n / 10
		k2 = n % 10
		if k1 > k2 {
			return false
		}
		fmt.Println(k1, k2, n)
	}

	return true
}

func main() {
	// arr1 := []string{"qetrqdqtihpojtm", "njweoiruwpkplnuk", "ylpzqgttahturo", "qetrqdqtihfanfkclsaj", "qetrqdqtihgchocxbvrf"}
	// arr2 := []string{"okl", "qetrq", "qetrqdqtihgch", "ylpzqgt", "pycgykhehq", "qetrqdqtihpo", "ylpzqgttaht", "ylpzqgttahturo"}
	// fmt.Println(numberOfPrefixes(arr1, arr2))
	fmt.Println(checkIsReduce(321))
}
