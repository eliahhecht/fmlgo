package main

// A Card represents a Magic card
type Card string

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
