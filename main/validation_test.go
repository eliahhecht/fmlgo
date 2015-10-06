package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationForLegalCard(t *testing.T) {
	testPlayer := Player{Cards: makeCards("Black Lotus")}
	testCards := makeCards("Black Lotus", "Ancestral Recall")

	confirmCardsAreLegal([]Player{testPlayer}, testCards)
}

func TestValidationForIllegalCard(t *testing.T) {
	testPlayer := Player{Cards: makeCards("Blacker Lotus")}
	testCards := makeCards("Black Lotus", "Ancestral Recall")

	assert.Panics(t, func() {
		confirmCardsAreLegal([]Player{testPlayer}, testCards)
	})
}
