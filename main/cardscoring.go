package main

// ScoreCards calculates how many points each card earned given the decklists
func ScoreCards(decklists []Decklist, legalCards []Card) map[CardName]int {
	result := make(map[CardName]int)
	legalCardsMap := convertToMap(legalCards)
	for _, decklist := range decklists {
		for _, decklistCard := range decklist.Cards {
			if legalCardsMap[decklistCard.Name] {
				result[decklistCard.Name]++
			}
		}
	}
	return result
}

func convertToMap(cards []Card) map[CardName]bool {
	m := make(map[CardName]bool)
	for _, card := range cards {
		m[card.Name] = true
	}
	return m
}
