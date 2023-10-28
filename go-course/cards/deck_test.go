package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as the first deck card, but got %v", d[len(d)-1])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs as the first deck card, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndnewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)
	loadedFromFile := newDeckFromFile(filename)

	if len(loadedFromFile) != 16 {
		t.Errorf("Exprected a file length of 16, got %v", len(loadedFromFile))
	}

	os.Remove(filename)
}
