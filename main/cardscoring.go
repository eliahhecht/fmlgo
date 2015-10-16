package main

// ScoreCards calculates how many points each card earned given the decklists
func ScoreCards(decklists []Decklist, legalCards *CardCollection) {
	for _, decklist := range decklists {
		for _, decklistCard := range decklist.Cards {
			card := legalCards.GetCard(decklistCard.Name)
			if card != nil && card.Rarity != "Basic Land" {
				card.Score++
			}
		}
	}
}

// TagOwners associates each card with the player who owns it
func TagOwners(players []*Player, cards *CardCollection) {
	for _, player := range players {
		tagCards(player, player.Cards, false)
		tagCards(player, player.Bench, true)
	}
}

func tagCards(player *Player, cardsToTag []*Card, onBench bool) {
	for _, card := range cardsToTag {
		card.Ownership = OwnershipTag{Owner: player.Name, OnBench: onBench}
	}
}
