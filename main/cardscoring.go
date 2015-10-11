package main

// ScoreCards calculates how many points each card earned given the decklists
func ScoreCards(decklists []Decklist, legalCards *CardCollection) {
	for _, decklist := range decklists {
		for _, decklistCard := range decklist.Cards {
			card := legalCards.GetCard(decklistCard.Name)
			if card != nil {
				card.Score++
			}
		}
	}
}
