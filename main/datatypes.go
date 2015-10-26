package main

import (
	"fmt"

	"github.com/deckarep/golang-set"
)

// A CardType is a Magic card type (Artifact, Creature, etc)
type CardType string

var AllCardTypes []CardType = []CardType{
	"Artifact", "Creature", "Enchantment", "Land", "Planeswalker", "Instant", "Sorcery"}

// A CardName is the name of a Magic card, and uniquely determines that card for FML purposes
type CardName string

// A PlayerName is the name of a FML player
type PlayerName string

// Rarity is the rarity of a card (Common, Uncommon, etc)
type Rarity string

// SetCode is the two- or three-letter code for a set (e.g. BFZ for Battle for Zendikar)
type SetCode string

// A Card represents a Magic card
type Card struct {
	Name      CardName
	Types     mapset.Set
	OtherSide CardName
	Score     int
	Ownership OwnershipTag
	Rarity    Rarity
	SetCodes  mapset.Set
}

type OwnershipTag struct {
	Owner   PlayerName
	OnBench bool
}

func (c Card) IsOwned() bool {
	return c.Ownership.Owner != ""
}

type CardCollection struct {
	CardsByName map[CardName]*Card
}

func (c *CardCollection) Contains(name CardName) bool {
	_, ok := c.CardsByName[name]
	return ok
}

func (c *CardCollection) GetCard(name CardName) *Card {
	card, ok := c.CardsByName[name]
	if !ok {
		panic(fmt.Sprintf("Caller asked for card %v, but it was not found", name))
	}
	return card
}

func (c *CardCollection) GetCardsOfType(cardType CardType) []*Card {
	var results []*Card
	for _, card := range c.CardsByName {
		if card.Types.Contains(cardType) {
			results = append(results, card)
		}
	}
	return results
}

// A Set is a slice of cards plus a set code
type Set struct {
	Code  SetCode
	Cards []*Card
}

// Player represents an FML player
type Player struct {
	Name  PlayerName
	Cards []*Card
	Bench []*Card
}

func (p Player) TotalScore() int {
	total := 0
	for _, card := range p.Cards {
		total += card.Score
	}
	return total
}
