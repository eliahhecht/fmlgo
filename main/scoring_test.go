package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoringForSimpleCase(t *testing.T) {
	decklist := Decklist{
		Cards: []Card{"Black Lotus", "Black Lotus", "Time Walk"}}
	player := Player{
		Name:  "Test Player",
		Cards: []Card{"Black Lotus"}}

	score := calculateScore([]Player{player}, []Decklist{decklist})

	assert.Equal(t, 2.0, score["Test Player"].Total())
}
