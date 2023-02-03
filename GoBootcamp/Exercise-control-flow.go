package main

import (
	"fmt"
)

func coinsForUser(name string) int {
	total := 0
	for _, n := range name {
		switch string(n) {
		case "a", "A":
			total++
		case "e", "E":
			total++
		case "i", "I":
			total += 2
		case "o", "O":
			total += 3
		case "u", "U":
			total += 4
		}
	}
	return total
}

func main() {
	coins := 50
	users := []string{
		"Matthew",
		"Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}

	distribution := make(map[string]int, len(users))

	for _, u := range users {
		v := coinsForUser(u)
		if v > 10 {
			v = 10
		}
		distribution[u] = v
		coins -= v
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
