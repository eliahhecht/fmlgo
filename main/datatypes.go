package main

import "github.com/deckarep/golang-set"

// A CardType is a Magic card type (Artifact, Creature, etc)
type CardType string

// A CardName is the name of a Magic card, and uniquely determines that card for FML purposes
type CardName string

// A PlayerName is the name of a FML player
type PlayerName string

// A Card represents a Magic card
type Card struct {
	Name      CardName
	Types     mapset.Set
	OtherSide CardName
	Score     int
	Ownership  OwnershipTag
}

type OwnershipTag struct {
	Owner PlayerName
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
	return c.CardsByName[name]
}

// A Set is a slice of cards plus a set code
type Set struct {
	Code  string
	Cards []*Card
}

// Player represents an FML player
type Player struct {
	Name  PlayerName
	Cards []*Card
	Bench []*Card
}
