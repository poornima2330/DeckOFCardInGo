package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))

	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades ,but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs , but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove the file if it exists
	os.Remove("_decktesting")

	deck := newDeck()

	// Save the deck to a file
	deck.saveToFile("_decktesting")

	// Load the deck from the file
	loadedDeck := newDeckFromFile("_decktesting")

	// Check if the loaded deck has the correct number of cards
	if len(loadedDeck) != 56 {
		t.Errorf("Expected 64 cards in deck, got %v", len(loadedDeck))
	}

	// Remove the testing file after the test
	os.Remove("_decktesting")
}
