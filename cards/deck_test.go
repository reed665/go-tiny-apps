package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	deckLen := len(deck)

	if deckLen != 52 {
		t.Errorf("Expected deck length of 52 but got %v", deckLen)
	}

	if deck.contains("Ace of Spades") == false {
		t.Errorf("Expected deck to contain 'Ace of Spades'")
	}

	if deck.contains("King of Clubs") == false {
		t.Errorf("Expected deck to contain 'King of Clubs'")
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
