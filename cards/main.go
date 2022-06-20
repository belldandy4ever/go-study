package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, remainingCards := deal(cards, 3)
	// hand.print()
	// remainingCards.print()
	// hand.saveToFile("myHandCard.txt")

	// testCards := newDeckFromFile("myHandCard.txt")
	// fmt.Println(testCards)

}

func newCard() string {
	return "Five of Diamonds"
}
