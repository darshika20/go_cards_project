package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck size of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of deck as 'Ace of Spades' but found %v", d[0])
	}

	if d[15] != "Four of Clubs" {
		t.Errorf("Expected last card of deck as 'Four of Clubs' but found %v", d[15])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {

	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	newDeck := readFromFile("_deckTesting")
	if len(newDeck) != 16 {
		t.Errorf("Expected deck size of 16, but found %v", len(newDeck))
	}
}
