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

// TagOwners associates each card with the player who owns it
func TagOwners(players []*Player, cards *CardCollection) {
	for _, player := range players {
		for _, card := range player.Cards {
			cards.GetCard(card.Name).Ownership = OwnershipTag{Owner: player.Name, OnBench: false}
		}
		for _, card := range player.Bench {
			cards.GetCard(card.Name).Ownership = OwnershipTag{Owner: player.Name, OnBench: true}
		}
	}

}
