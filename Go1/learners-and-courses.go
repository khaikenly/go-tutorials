package main

import "fmt"

func main() {
	learnerCourses := map[string][]string{
		"Learner-0001": []string{
			"Course-0001",
			"Course-0002",
			"Course-0003",
			"Course-0001",
			"Course-0002",
			"Course-0003",
		},
		"Learner-0002": []string{
			"Course-0002",
			"Course-0003",
			"Course-0004",
		},
	}
	coursesCounter := make(map[string]int)
	for _, courses := range learnerCourses {
		for _, course := range courses {
			coursesCounter[course] = coursesCounter[course] + 1
		}
	}

	uniqueLearnerCourses := []string{}
	for course, counter := range coursesCounter {
		if counter == 1 {
			uniqueLearnerCourses = append(uniqueLearnerCourses, course)
		}
	}
	fmt.Println(uniqueLearnerCourses)
}
