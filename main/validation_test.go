package main

import "testing"

func TestValidationForLegalCard(t *testing.T) {
    testPlayer := Player {Cards: []Card {{"Black Lotus"}}}
    testSet := Set {Cards: []Card {{"Black Lotus"}, {"Ancestral Recall"}}}

    confirmCardsAreLegal([]Player {testPlayer}, []Set {testSet})
}
