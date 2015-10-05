package main

// ScoreCards calculates how many points each card earned given the decklists
func ScoreCards(decklists []Decklist, legalCards []Card) map[Card]int {
	result := make(map[Card]int)
	legalCardsMap := convertToMap(legalCards)
	for _, decklist := range decklists {
		for _, decklistCard := range decklist.Cards {
			if legalCardsMap[decklistCard] {
				result[decklistCard]++
			}
		}
	}
	return result
}

func convertToMap(cards []Card) map[Card]bool {
	m := make(map[Card]bool)
	for _, card := range cards {
		m[card] = true
	}
	return m
}
