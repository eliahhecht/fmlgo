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
	score := invokeCalculateScore()

	assert.Equal(t, 2.0, score.Total())
}

func invokeCalculateScore() ScoreResult {
	results := calculateScore([]Player{testPlayer}, testCardScores)
	return results[testPlayer.Name]
}

func TestScoringForWinsMultiplier(t *testing.T) {
	testCardScores["Black Lotus"] = 2.6

	score := invokeCalculateScore()

	assert.Equal(t, 2.6, score.Total())
}
