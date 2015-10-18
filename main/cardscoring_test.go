package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDecklists = []Decklist{Decklist{Cards: makeCards("Black Lotus", "Black Lotus", "Storm Crow")}}

func TestCardScorer_ScoresCardCorrectly(t *testing.T) {
	legalCards := makeCardCollection("Black Lotus", "Storm Crow")
	ScoreCards(testDecklists, legalCards)

	blackLotusScore := legalCards.GetCard("Black Lotus").Score
	if (blackLotusScore != 2) {
		t.Errorf("Wanted %d, got %d", 2, blackLotusScore)
	}
}

var testCardCollection = makeCardCollection("Black Lotus", "Plains", "Storm Crow")

var testPlayer = &Player{
	Name:  "Test Player",
	Cards: []*Card{testCardCollection.GetCard("Black Lotus")},
	Bench: []*Card{testCardCollection.GetCard("Storm Crow")}}

func TestTaggingCardsToOwners(t *testing.T) {
	invokeTagOwners()

	blackLotus := testCardCollection.GetCard("Black Lotus")
	assert.Equal(t, testPlayer.Name, blackLotus.Ownership.Owner)
}

func invokeTagOwners() {
	TagOwners([]*Player{testPlayer}, testCardCollection)
}

func TestTaggingCardsThatAreNotOwned(t *testing.T) {
	invokeTagOwners()

	plains := testCardCollection.GetCard("Plains")
	assert.False(t, plains.IsOwned())
}
