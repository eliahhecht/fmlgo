package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var decklistDir string

func main() {
	parseFlags()

	legalCards := loadCards([]string{"KTK", "FRF", "DTK", "ORI", "BFZ", "EXP"})

	players := buildPlayers(legalCards)

	confirmCardsAreLegal(players, legalCards)

	loadDecklists()
	ScoreCards(decklists, legalCards)

	TagOwners(players, legalCards)

	outputScores(players, legalCards)
}

func parseFlags() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		panic("Must be called with exactly one argument (path to decklists dir)")
	}
	decklistDir = flag.Arg(0)
}

const cardOwnersFile string = "../data/owners.json"

func buildPlayers(allCards *CardCollection) []*Player {
	var playersMap map[string]map[string][]string
	readJSONFile(cardOwnersFile, &playersMap)

	var players []*Player
	for k, v := range playersMap {
		mainDeck := buildCards(v["main"], allCards)
		bench := buildCards(v["bench"], allCards)
		player := Player{Name: PlayerName(k), Cards: mainDeck, Bench: bench}
		players = append(players, &player)
	}
	return players
}

func buildCards(cardNames []string, allCards *CardCollection) []*Card {
	var cards = make([]*Card, 0)
	for _, cardName := range cardNames {
		card := allCards.GetCard(CardName(cardName))
		cards = append(cards, card)
	}
	return cards
}

var decklists []Decklist

func loadDecklists() {
	filepath.Walk(decklistDir, loadDecklist)
}

func loadDecklist(path string, f os.FileInfo, err error) error {
	fmt.Println("Loading ", path)
	if !f.IsDir() {
		loader := newDecklistLoader()
		decklist := loader.loadDecklist(path)
		decklists = append(decklists, decklist)
	}
	return nil
}
