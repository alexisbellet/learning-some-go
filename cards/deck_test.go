package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected a deck length of 52 and instead got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card of the deck to be Ace of Spades and instead got %v", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected the last card of the deck to be King of Clubs and instead got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck,_ := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected a deck length of 52 and instead got %v", len(loadedDeck))
	}
}