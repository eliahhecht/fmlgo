package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecklistParsesCorrectCards(t *testing.T) {
	decklistString := `
    1 Llanowar Wastes
    2 Black Lotus
    57 Swamp

    sideboard
    3 Ancestral Recall
    12 Plains`

	cards := parseDecklist(decklistString)

	assert.Len(t, cards, 75)

	assertContains(t, cards, "Llanowar Wastes", 1)
	assertContains(t, cards, "Black Lotus", 2)
	assertContains(t, cards, "Swamp", 57)
	assertContains(t, cards, "Ancestral Recall", 3)
	assertContains(t, cards, "Plains", 12)
}

func assertContains(t *testing.T, haystack []Card, needle Card, expectedNumber int) {
	numberFound := 0
	for _, card := range haystack {
		if card == needle {
			numberFound++
		}
	}

	assert.Equal(t, expectedNumber, numberFound)
}

//ehtodo test makes sure 75 cards

func TestDecklistParsesMultiplierFromFileName(t *testing.T) {
	systemUnderTest := decklistLoader{
		loadFile: func(string) []byte { return nil }}

	var decklist = systemUnderTest.loadDecklist("foo.1.3.txt")

	assert.Equal(t, 1.3, decklist.ScoreMultiplier)
}
