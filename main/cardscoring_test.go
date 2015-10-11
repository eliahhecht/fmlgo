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
