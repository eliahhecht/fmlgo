package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPlayer = Player{
	Name:      "Test Player",
	Cards:     []Card{"Black Lotus"},
	Sideboard: []Card{"Storm Crow"}}

var testCardScores = map[Card]int{
	"Black Lotus": 2,
	"Storm Crow":  100}

func TestScoringForSimpleCase(t *testing.T) {
	score := invokeCalculateScore().PlayerScores["Test Player"]

	assert.Equal(t, 2, score.Total())
}

func invokeCalculateScore() OverallResult {
	return calculateScore([]Player{testPlayer}, testCardScores)
}

func TestScoringForUnownedCards(t *testing.T) {
	testCardScores["Mox Pearl"] = 100

	score := invokeCalculateScore()

	assert.Equal(t, 100, score.UnownedCardScores["Mox Pearl"])
}

func TestScoringForUnownedCardsDoesNotIncludeOwnedCards(t *testing.T) {
	score := invokeCalculateScore()

	assert.Equal(t, 0, score.UnownedCardScores["Black Lotus"])
}

func TestScoringForSideboardCards(t *testing.T) {
	score := invokeCalculateScore().PlayerScores["Test Player"]

	assert.Equal(t, 100, score.SideboardScores["Storm Crow"])
}
