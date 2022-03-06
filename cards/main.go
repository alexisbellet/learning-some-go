package main

func main() {
	cards, err := newDeckFromFile("deck.txt")
	if (err != nil) {
		cards = newDeck()
	}
	cards.shuffle()
	cards.print()
	saveToFileErr := cards.saveToFile("deck.txt")
	if (saveToFileErr != nil) {
		panic(saveToFileErr)
	}
}