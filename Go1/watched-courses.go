package main

import (
	"fmt"
	"sort"
)

func main() {
	sessions := [][]string{
		{"Course_001", "Course_002", "Course_003", "Course_004"},
		{"Course_001", "Course_003"},
		{"Course_002", "Course_004", "Course_001"},
		{"Course_004", "Course_002", "Course_003", "Course_001"},
		{"Course_004", "Course_003", "Course_002", "Course_001"},
	}

	for _, course := range sessions {
		sort.Strings(course)
	}

	courseMap := make(map[string]map[string]int)
	for _, session := range sessions {
		for i := 0; i < len(session)-1; i++ {
			if _, ok := courseMap[session[i]]; !ok {
				courseMap[session[i]] = make(map[string]int)
			}
			if _, ok := courseMap[session[i]][session[i+1]]; !ok {
				courseMap[session[i]][session[i+1]] = 0
			}
			courseMap[session[i]][session[i+1]]++
		}
	}

	result := make(map[string]string)
	for course, nextMap := range courseMap {
		maxNext := ""
		maxCount := 0
		for next, count := range nextMap {
			if count > maxCount {
				maxNext = next
				maxCount = count
			}
		}
		result[course] = maxNext
	}

	fmt.Println(result)
}
