package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	arrStr := strings.Fields(s)
	m := make(map[string]int)

	for _, v := range arrStr {
		m[v]++
	}

	return m
}

func main() {
	str := "A man a plan a canal panama."
	fmt.Println(WordCount(str))
}
