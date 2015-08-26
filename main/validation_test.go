package main

import
(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestValidationForLegalCard(t *testing.T) {
    testPlayer := Player {Cards: []Card {"Black Lotus"}}
    testSet := Set {Cards: []Card {"Black Lotus", "Ancestral Recall"}}

    confirmCardsAreLegal([]Player {testPlayer}, []Set {testSet})
}

func TestValidationForIllegalCard(t *testing.T) {
    testPlayer := Player {Cards: []Card {"Blacker Lotus"}}
    testSet := Set {Cards: []Card {"Black Lotus", "Ancestral Recall"}}

    assert.Panics(t, func() {
        confirmCardsAreLegal([]Player {testPlayer}, []Set {testSet})
    })
}
