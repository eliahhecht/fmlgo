package main

func makeCards(names ...string) []*Card {
	var cards []*Card
	for _, name := range names {
		cards = append(cards, &Card{Name: CardName(name)})
	}
	return cards
}

func makeCardCollection(names ...string) *CardCollection {
	cards := makeCards(names...)
	collection := CardCollection{make(map[CardName]*Card)}
	for _, card := range cards {
		collection.CardsByName[card.Name] = card
	}
	return &collection
}
