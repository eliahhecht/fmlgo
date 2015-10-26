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

	assertContains(t, cards, &Card{Name: "Llanowar Wastes"}, 1)
	assertContains(t, cards, &Card{Name: "Black Lotus"}, 2)
	assertContains(t, cards, &Card{Name: "Swamp"}, 57)
	assertContains(t, cards, &Card{Name: "Ancestral Recall"}, 3)
	assertContains(t, cards, &Card{Name: "Plains"}, 12)
}

func assertContains(t *testing.T, haystack []*Card, needle *Card, expectedNumber int) {
	numberFound := 0
	for _, card := range haystack {
		if card.Name == needle.Name {
			numberFound++
		}
	}

	assert.Equal(t, expectedNumber, numberFound)
}

//ehtodo test makes sure 75 cards
