package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

var decklistDir string

func cliMain() {
	parseFlags()

	allCards := loadAllCards()

	players := buildPlayers(allCards)

	confirmCardsAreLegal(players, allCards)

	loadDecklists()
	ScoreCards(decklists, allCards)

	TagOwners(players, allCards)

	outputScores(players, allCards, []SetCode{"KTK", "FRF", "DTK", "ORI", "BFZ", "EXP"})
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
	if looksLikeDecklist(path, f) {
		loader := newDecklistLoader()
		decklist := loader.loadDecklist(path)
		decklists = append(decklists, decklist)
	}
	return nil
}

func looksLikeDecklist(path string, f os.FileInfo) bool {
	if f.IsDir() {
		return false
	}
	return strings.HasSuffix(path, ".txt")
}
