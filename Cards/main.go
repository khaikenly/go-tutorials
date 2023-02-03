package main

import "fmt"

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	cards := []string{"Two of Diamonds", newCard(), "Six of Diamonds", "Seven of Diamonds"}
	cards = append(cards, "Eight of Diamonds") // append new element to last slice
	for i, v := range cards {
		fmt.Println(i, v)
	}
}
