package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cards = loadCards([]string{"ORI", "BFZ"})

func TestLoadSetsDoesNotIncludeBasicLands(t *testing.T) {
	card := findCard("Swamp")
	assert.Nil(t, card)
}

func findCard(needle CardName) *Card {
	for _, card := range cards {
		if card.Name == needle {
			return &card
		}
	}
	return nil
}

func TestLoadCardsContainsCorrectCards(t *testing.T) {
	yavimayaCoast := findCard("Yavimaya Coast")
	assert.NotNil(t, yavimayaCoast)
	transgress := findCard("Transgress the Mind")
	assert.NotNil(t, transgress)
}

func TestLoadSetsLoadsCardTypes(t *testing.T) {
	hangarbackWalker := *findCard("Hangarback Walker")

	assert.Contains(t, hangarbackWalker.Types, CardType("Artifact"))
	assert.Contains(t, hangarbackWalker.Types, CardType("Creature"))
}

//ehtodo test types for DFC
