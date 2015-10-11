package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDecklists = []Decklist{Decklist{Cards: makeCards("Black Lotus", "Black Lotus", "Storm Crow")}}

func TestCardScorer_ScoresCardCorrectly(t *testing.T) {
	legalCards := makeCardCollection("Black Lotus")
	ScoreCards(testDecklists, legalCards)

	assert.Equal(t, 2, legalCards.GetCard("Black Lotus").Score)
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
