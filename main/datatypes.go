package main

// A CardType is a Magic card type (Artifact, Creature, etc)
type CardType string

// A CardName is the name of a Magic card, and uniquely determines that card for FML purposes
type CardName string

// A Card represents a Magic card
type Card struct {
	Name  CardName
	Types []CardType
}

// A Set is a slice of cards plus a set code
type Set struct {
	Code  string
	Cards []Card
}

// Player represents an FML player
type Player struct {
	Name      string
	Cards     []Card
	Sideboard []Card
}
