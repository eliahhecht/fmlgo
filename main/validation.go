package main

import "fmt"

func confirmCardsAreLegal(players []*Player, legalCards *CardCollection) {
	for _, player := range players {
		for _, card := range append(player.Cards, player.Bench...) {
			if !legalCards.Contains(card.Name) {
				panic(fmt.Sprintf("%s's card %s not found in any legal set (%v legal cards)",
					player.Name, card, len(legalCards.CardsByName)))
			}
		}
	}
}
