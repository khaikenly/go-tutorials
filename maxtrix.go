package main

import "fmt"

func main() {
	arrMaxtrix := [][]int{
		[]int{9, 9, 9, 9, 9, 9, 9},
		[]int{9, 9, 9, 9},
	}
	var result []int
	arrF := arrMaxtrix[0]
	tmp := 0
	for i := 1; i < len(arrMaxtrix); i++ {
		for j, v := range arrMaxtrix[i] {

			sum := arrF[j] + v

			if tmp != 0 {
				sum += tmp
				tmp = 0
			}

			if sum >= 10 {
				tmp = sum / 10
				sum -= tmp * 10
			}

			result = append(result, sum)

			if j == (len(arrMaxtrix[i])-1) && tmp != 0 {
				result = append(result, tmp)
			}
		}
	}
	fmt.Println(result)
}
