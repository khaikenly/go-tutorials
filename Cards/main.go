package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	cards := deck{"Two of Diamonds", newCard(), "Six of Diamonds", "Seven of Diamonds"}
	cards = append(cards, "Eight of Diamonds")
	cards.print()

}
