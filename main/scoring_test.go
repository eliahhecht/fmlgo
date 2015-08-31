package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPlayer = Player{
	Name:  "Test Player",
	Cards: []Card{"Black Lotus"}}

var testCardScores = map[Card]float64{
	"Black Lotus": 2}

func TestScoringForSimpleCase(t *testing.T) {
	score := invokeCalculateScore().PlayerScores["Test Player"]

	assert.Equal(t, 2.0, score.Total())
}

func invokeCalculateScore() OverallResult {
	return calculateScore([]Player{testPlayer}, testCardScores)
}

func TestScoringForWinsMultiplier(t *testing.T) {
	testCardScores["Black Lotus"] = 2.6

	score := invokeCalculateScore().PlayerScores["Test Player"]

	assert.Equal(t, 2.6, score.Total())
}

func TestScoringForUnownedCards(t *testing.T) {
	testCardScores["Mox Pearl"] = 100

	score := invokeCalculateScore()

	assert.Equal(t, 100.0, score.UnownedCardScores["Mox Pearl"])
}

func TestScoringForUnownedCardsDoesNotIncludeOwnedCards(t *testing.T) {
	score := invokeCalculateScore()

	assert.Equal(t, 0.0, score.UnownedCardScores["Black Lotus"])
}
