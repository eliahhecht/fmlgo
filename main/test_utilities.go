package main

func makeCards(names ...string) []Card {
	var cards []Card
	for _, name := range names {
		cards = append(cards, Card{Name: CardName(name)})
	}
	return cards
}
