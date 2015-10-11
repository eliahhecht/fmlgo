package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cards = loadCards([]string{"ORI", "BFZ"})

func TestLoadSetsDoesNotIncludeBasicLands(t *testing.T) {
	card := cards.GetCard("Swamp")
	assert.Nil(t, card)
}

func TestLoadCardsContainsCorrectCards(t *testing.T) {
	yavimayaCoast := cards.GetCard("Yavimaya Coast")
	assert.NotNil(t, yavimayaCoast)
	transgress := cards.GetCard("Transgress the Mind")
	assert.NotNil(t, transgress)
}

func TestLoadSetsLoadsCardTypes(t *testing.T) {
	hangarbackWalker := *cards.GetCard("Hangarback Walker")

	assert.Contains(t, hangarbackWalker.Types.ToSlice(), CardType("Artifact"))
	assert.Contains(t, hangarbackWalker.Types.ToSlice(), CardType("Creature"))
}

func TestLoadSetsLoadsAllTypesForDoubleFacedCard(t *testing.T) {
	jace := *cards.GetCard("Jace, Vryn's Prodigy")

	assert.Contains(t, jace.Types.ToSlice(), CardType("Planeswalker"))
	assert.Contains(t, jace.Types.ToSlice(), CardType("Creature"))
}
