package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDecklists = []Decklist{Decklist{Cards: makeCards("Black Lotus", "Black Lotus", "Storm Crow")}}

func TestCardScorer_ScoresCardCorrectly(t *testing.T) {
	legalCards := makeCards("Black Lotus")
	scoreResult := ScoreCards(testDecklists, legalCards)

	assert.Equal(t, 2, scoreResult["Black Lotus"])
}

func TestCardScorer_DoesNotScoreIllegalCards(t *testing.T) {
	legalCards := makeCards("Storm Crow")

	scoreResult := ScoreCards(testDecklists, legalCards)

	assert.Equal(t, 0, scoreResult["Black Lotus"])
}
