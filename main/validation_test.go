package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLegalCards = makeCardCollection("Black Lotus", "Ancestral Recall")

func TestValidationForLegalCard(t *testing.T) {
	testPlayer := Player{Cards: makeCards("Black Lotus")}

	confirmCardsAreLegal([]Player{testPlayer}, testLegalCards)
}

func TestValidationForIllegalCard(t *testing.T) {
	testPlayer := Player{Cards: makeCards("Blacker Lotus")}

	assert.Panics(t, func() {
		confirmCardsAreLegal([]Player{testPlayer}, testLegalCards)
	})
}

func TestValidationForBenchedCards(t *testing.T) {
	testPlayer := Player{Cards: makeCards("Black Lotus"), Bench: makeCards("Bad Ass")}

	assert.Panics(t, func() {
		confirmCardsAreLegal([]Player{testPlayer}, testLegalCards)
	})
}
