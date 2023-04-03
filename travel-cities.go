package main

import "fmt"

func unique(intSlice []int) ([]int, int) {
	keys := make(map[int]bool)
	cities := []int{}
	var day int
	for i, entry := range intSlice {
		if value := keys[entry]; !value {
			keys[entry] = true
			cities = append(cities, entry)
			day = i
		}
	}
	return cities, day
}

func Reverse(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func main() {
	travel := []int{7, 4, 7, 3, 4, 1, 7}
	cities, day := unique(travel)
	fmt.Println(cities, day)
	fmt.Println(Reverse("Hello"))
}
