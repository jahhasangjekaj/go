package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	var card = "Ace of Spades"
	if d[0] != card {
		t.Errorf("Expected first card of %v, but got %v", card, d[0])
	}

	card = "King of Clubs"
	if d[len(d)-1] != card {
		t.Errorf("Expected fourth card of %v, but got %v", card, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	var fileName = "_decktesting"

	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := newDeckFronFile(fileName)
	var expectedLength = 52
	if len(loadedDeck) != expectedLength {
		t.Errorf("Expected %v cards in the deck, got %v", expectedLength, len(loadedDeck))
	}

	os.Remove(fileName)
}
