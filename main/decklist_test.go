package main

import 
(
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

    decklist := parseDecklist(decklistString)

    assert.Len(t, decklist.Cards, 75)

    assertContains(t, decklist, "Llanowar Wastes", 1)
    assertContains(t, decklist, "Black Lotus", 2)
    assertContains(t, decklist, "Swamp", 57)
    assertContains(t, decklist, "Ancestral Recall", 3)
    assertContains(t, decklist, "Plains", 12)
}

func assertContains(t *testing.T, decklist Decklist, needle Card, expectedNumber int) {
    numberFound := 0
    for _, card := range decklist.Cards {
        if card == needle {
            numberFound++
        }
    }

    assert.Equal(t, expectedNumber, numberFound)
}