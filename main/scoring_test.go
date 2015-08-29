package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testPlayer = Player{
	Name:  "Test Player",
	Cards: []Card{"Black Lotus"}}

var testDecklist = Decklist{
	Cards:           []Card{"Black Lotus", "Black Lotus"},
	ScoreMultiplier: 1}

func TestScoringForSimpleCase(t *testing.T) {
	score := invokeCalculateScore()

	assert.Equal(t, 2.0, score.Total())
}

func invokeCalculateScore() ScoreResult {
	results := calculateScore([]Player{testPlayer}, []Decklist{testDecklist})
	return results[testPlayer.Name]
}

func TestScoringForWinsMultiplier(t *testing.T) {
	testDecklist.ScoreMultiplier = 1.3

	score := invokeCalculateScore()

	assert.Equal(t, 2.6, score.Total())
}
