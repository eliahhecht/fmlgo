package main

import "fmt"

func confirmCardsAreLegal(players []Player, legalSets []Set) {
    for _, player := range players {
        for _, card := range player.Cards {
            if !anySetContainsCard(legalSets, card) {
                panic(fmt.Sprintf("%s's card %s not found in any legal set", player.Name, card))
            }
        }
    }
}

func anySetContainsCard(sets []Set, card Card) bool {
    for _, set := range sets {
        if set.containsCard(card) {
            return true
        }
    }
    return false
}

func (set Set) containsCard(targetCard Card) bool {
    for _, card := range set.Cards {
        if card == targetCard {
            return true
        }
    }
    return false
}
