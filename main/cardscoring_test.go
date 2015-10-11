package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDecklists = []Decklist{Decklist{Cards: makeCards("Black Lotus", "Black Lotus", "Storm Crow")}}

func TestCardScorer_ScoresCardCorrectly(t *testing.T) {
	legalCards := makeCards("Black Lotus")
	ScoreCards(testDecklists, legalCards)

	assert.Equal(t, 2, legalCards[0].Score)
}

func TestCardScorer_DoesNotScoreIllegalCards(t *testing.T) {
	legalCards := makeCards("Storm Crow")

	ScoreCards(testDecklists, legalCards)

	assert.Equal(t, 0, legalCards[0].Score)
}
