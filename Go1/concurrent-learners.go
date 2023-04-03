package main

import "fmt"

func main() {
	streams := [][]int{
		{0, 1000},
		{500, 2000},
		{2500, 300},
		{400, 1400},
		{1100, 180},
		{1200, 140},
		{500, 2400},
		{2100, 230},
	}
	maxConcurrentStreams := findMaxConcurrentStreams(streams)
	fmt.Printf("The maximum number of concurrent streams is %d.\n", maxConcurrentStreams)
}

func findMaxConcurrentStreams(streams [][]int) int {
	n := len(streams)
	start := make([]int, n)
	end := make([]int, n)
	for i := 0; i < n; i++ {
		start[i] = streams[i][0]
		end[i] = streams[i][1]
	}
	sortByStartTime(start, end)
	maxConcurrentStreams := 0
	currentStreams := 0
	i := 0
	j := 0
	for i < n && j < n {
		if start[i] < end[j] {
			currentStreams++
			i++
			if currentStreams > maxConcurrentStreams {
				maxConcurrentStreams = currentStreams
			}
		} else {
			currentStreams--
			j++
		}
	}
	return maxConcurrentStreams
}

func sortByStartTime(start []int, end []int) {
	n := len(start)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if start[i] > start[j] {
				start[i], start[j] = start[j], start[i]
				end[i], end[j] = end[j], end[i]
			}
		}
	}
}
