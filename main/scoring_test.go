package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPlayer = &Player{
	Name:  "Test Player",
	Cards: makeCards("Black Lotus"),
	Bench: makeCards("Storm Crow")}

var testCardScores = &CardCollection{map[CardName]*Card{
	"Black Lotus": &Card{Score: 2},
	"Storm Crow":  &Card{Score: 100},
	"Plains": &Card{Score: 42}}}


func TestScoringForSimpleCase(t *testing.T) {
	score := invokeCalculateScore().PlayerScores["Test Player"]

	assert.Equal(t, 2, score.Total())
}

func invokeCalculateScore() OverallResult {
	return calculateScore([]*Player{testPlayer}, testCardScores)
}

func TestScoringForUnownedCards(t *testing.T) {
	//	testCardScores.CardsByName["Mox Pearl"] = &Card{Score: 100}
	//
	//	score := invokeCalculateScore()
	//
	//	assert.Equal(t, 100, score.UnownedCardScores["Mox Pearl"])
}

func TestScoringForUnownedCardsDoesNotIncludeOwnedCards(t *testing.T) {
	score := invokeCalculateScore()

	assert.Equal(t, 0, score.UnownedCardScores["Black Lotus"])
}

func TestScoringForBenchCards(t *testing.T) {
	score := invokeCalculateScore().PlayerScores["Test Player"]

	assert.Equal(t, 100, score.BenchScores["Storm Crow"])
}

func TestTaggingCardsToOwners(t *testing.T) {
	invokeCalculateScore()

	blackLotus := testCardScores.GetCard("Black Lotus")
	assert.Equal(t, testPlayer.Name, blackLotus.Ownership.Owner)
}

func TestTaggingCardsThatAreNotOwned(t *testing.T) {
	invokeCalculateScore()

	plains := testCardScores.GetCard("Plains")
	assert.False(t, plains.IsOwned())
}
