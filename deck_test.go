package main

import (
	"fmt"
	"os"
	"testing"
)

const testDeckFileName = "_deckTesting.txt"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "As of Diamonds" {
		t.Errorf("Expected first card of 'As of Diamonds', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filePath := fmt.Sprintf("./files/%v", testDeckFileName)

	deck := newDeck()
	deck.saveToAFile(testDeckFileName)

	_, createdFileError := os.Stat(filePath)
	if createdFileError != nil {
		t.Error("Deck file was not created", createdFileError)
	}

	newDeck := *newDeckFromFile(testDeckFileName)

	if len(newDeck) != len(deck) {
		t.Errorf("Expected deck length was %v, but got %v", len(deck), len(newDeck))
	}

	deleteFileError := os.Remove(filePath)

	if deleteFileError != nil {
		t.Error("An unexpected error occurs while delete the test file for the second time.", deleteFileError)
	}
}
