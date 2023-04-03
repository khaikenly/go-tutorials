package main

import "fmt"

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var max int32 = 0
	for _, v := range candles {
		if max < v {
			max = v
		}
	}

	length := 0
	for _, v := range candles {
		if v == max {
			length++
		}
	}

	return int32(length)
}

func main() {
	candles := []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(candles))
}
