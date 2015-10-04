package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cards = loadCards([]string{"ORI", "BFZ"})

func TestLoadSetsDoesNotIncludeBasicLands(t *testing.T) {
	assert.NotContains(t, cards, Card("Swamp"))
}

func TestLoadCardsContainsCorrectCards(t *testing.T) {
	assert.Contains(t, cards, Card("Yavimaya Coast"))
	assert.Contains(t, cards, Card("Transgress the Mind"))
}
