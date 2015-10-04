package main

import "fmt"

func confirmCardsAreLegal(players []Player, legalCards []Card) {
	for _, player := range players {
		for _, card := range player.Cards {
			if !cardIsLegal(legalCards, card) {
				panic(fmt.Sprintf("%s's card %s not found in any legal set (%v legal cards)",
					player.Name, card, len(legalCards)))
			}
		}
	}
}

func cardIsLegal(legalCards []Card, card Card) bool {
	for _, legalCard := range legalCards {
		if card == legalCard {
			return true
		}
	}
	return false
}
