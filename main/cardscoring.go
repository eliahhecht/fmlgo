package main

// ScoreCards calculates how many points each card earned given the decklists
func ScoreCards(decklists []Decklist, legalCards []*Card) {
	cardsMap := convertToMap(legalCards)
	for _, decklist := range decklists {
		for _, decklistCard := range decklist.Cards {
			card := cardsMap[decklistCard.Name]
			if card != nil {
				card.Score++
			}
		}
	}
}

func convertToMap(cards []*Card) map[CardName]*Card {
	m := make(map[CardName]*Card)
	for _, card := range cards {
		m[card.Name] = card
	}
	return m
}
