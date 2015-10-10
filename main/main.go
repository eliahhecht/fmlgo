package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"
)

var decklistDir string

func main() {
	parseFlags()
	flag.Parse()

	legalCards := loadCards([]string{"KTK", "FRF", "DTK", "ORI", "BFZ", "EXP"})

	players := buildPlayers()

	confirmCardsAreLegal(players, legalCards)

	loadDecklists()
	oriCardScores := ScoreCards(decklists, legalCards)

	scores := calculateScore(players, oriCardScores)

	outputScores(scores)
}

func parseFlags() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		panic("Must be called with exactly one argument (path to decklists dir)")
	}
	decklistDir = flag.Arg(0)
}

func outputScores(scores OverallResult) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	var playerNames []string
	for playerName := range scores.PlayerScores {
		playerNames = append(playerNames, playerName)
	}
	sort.Strings(playerNames)
	for _, playerName := range playerNames {
		score := scores.PlayerScores[playerName]
		fmt.Fprintf(w, "== %s: \t%d ==\n", playerName, score.Total())
		printCardScores(w, score.CardScores, math.MaxInt32)
		fmt.Fprintln(w, "  Sideboard:")
		printCardScores(w, score.BenchScores, math.MaxInt32)
		fmt.Fprintln(w, "\t")
	}

	fmt.Fprintln(w, "Highest-scoring unfielded cards: ")
	printCardScores(w, scores.UnownedCardScores, 50)

	w.Flush()
}

type sortedScoreMap struct {
	m    map[CardName]int
	keys []CardName
}

func newSortedScoreMap(cardScores map[CardName]int) sortedScoreMap {
	sorted := sortedScoreMap{m: cardScores}
	for key := range cardScores {
		sorted.keys = append(sorted.keys, key)
	}
	return sorted
}

func (sm *sortedScoreMap) Len() int {
	return len(sm.keys)
}

func (sm *sortedScoreMap) Less(i, j int) bool {
	iVal := sm.m[sm.keys[i]]
	jVal := sm.m[sm.keys[j]]

	if iVal == jVal {
		return sm.keys[i] < sm.keys[j]
	}
	return iVal < jVal
}

func (sm *sortedScoreMap) Swap(i, j int) {
	sm.keys[i], sm.keys[j] = sm.keys[j], sm.keys[i]
}

func printCardScores(w io.Writer, cardScores map[CardName]int, max int) {
	sortable := newSortedScoreMap(cardScores)
	sort.Sort(sort.Reverse(&sortable))

	for i, cardName := range sortable.keys {
		cardScore := sortable.m[cardName]
		fmt.Fprintf(w, "   %v \t%d\n", cardName, cardScore)

		if i >= max {
			break
		}
	}
}

const cardOwnersFile string = "../data/owners.json"

func buildPlayers() []Player {
	var playersMap map[string]map[string][]string
	readJSONFile(cardOwnersFile, &playersMap)

	var players []Player
	for k, v := range playersMap {
		mainDeck := buildCards(v["main"])
		bench := buildCards(v["bench"])
		player := Player{Name: k, Cards: mainDeck, Bench: bench}
		players = append(players, player)
	}
	return players
}

func buildCards(cardNames []string) []*Card {
	var cards = make([]*Card, 0)
	for _, cardName := range cardNames {
		cards = append(cards, &Card{Name: CardName(cardName)})
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
