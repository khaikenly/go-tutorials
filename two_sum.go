package main

func main() {
	nums := []int{6, 4, 7, 9, 3, 12}
	target := 9

	m := make(map[int]int)
	for _, v := range nums {
		t := target - v
		ok = m[t]
		if ok {

		}
		m[t] = v
	}

	// for _, v := range nums {
	// 	for _, x := range nums {
	// 		if v != x && v+x == target {
	// 			fmt.Println(v, x)
	// 		}
	// 	}
	// }
}
