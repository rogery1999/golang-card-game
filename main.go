package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(&cards, 3)

	// hand.print()
	// remainingDeck.print()

	// cards.saveToAFile("my_cards.txt")
	lastDeck := *newDeckFromFile("my_cards.txt")
	lastDeck.shuffle()
	lastDeck.print()
}
