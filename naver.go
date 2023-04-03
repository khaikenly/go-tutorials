package main

import "fmt"

func slution(s string) {
	tmp := 0
	for i := 0; i < len(s); i++ {
		if tmp != 0 {
			tmp++
		}
		if string(s[i]) == "1" {
			tmp = i
		}
	}
	fmt.Println(tmp)
}

func main() {
	str := "1000100001"
	slution(str)
}
