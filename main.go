package main

func main() {
	// cards := newDeck()

	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("xmy_cards")
	cards.print()

}
