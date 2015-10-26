package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/deckarep/golang-set"
)

var cards = loadAllCards()

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

func TestLoadCardsLoadsAllSetsForCards(t *testing.T) {
	anticipate := *cards.GetCard("Anticipate")

	if !anticipate.SetCodes.Equal(mapset.NewSetFromSlice([]interface{} {SetCode("DTK"), SetCode("BFZ")})) {
		t.Errorf("Expected card to have sets BFZ and DTK, but was %v", anticipate.SetCodes)
	}
}
