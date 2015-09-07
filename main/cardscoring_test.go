package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardScorer_ScoresCardCorrectly(t *testing.T) {
	decklists := []Decklist{
		Decklist{Cards: []Card{"Black Lotus", "Black Lotus", "Storm Crow"},
			ScoreMultiplier: 99}}

	legalCards := []Card{"Black Lotus"}
	scoreResult := ScoreCards(decklists, legalCards)

	assert.Equal(t, 198.0, scoreResult["Black Lotus"])
}

func TestCardScorer_DoesNotScoreIllegalCards(t *testing.T) {
	decklists := []Decklist{
		Decklist{Cards: []Card{"Black Lotus", "Black Lotus", "Storm Crow"},
			ScoreMultiplier: 99}}

	cards := []Card{"Storm Crow"}

	scoreResult := ScoreCards(decklists, cards)

	assert.Equal(t, 0.0, scoreResult["Black Lotus"])
}