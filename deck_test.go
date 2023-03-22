package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected a deck of length 16, but got %d", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card is not Ace of Spades  %v", d[0])
	}

	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected last card is not Four of Clubs  %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(deck) != 16 {
		t.Errorf("Expected a deck of length 16, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
