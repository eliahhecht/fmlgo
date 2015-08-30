package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"
)

var decklistDir = flag.String(
	"decklistDir", "decklists", "path to decklists for scoring")

func main() {
	ori := loadSet("ORI")

	players := buildPlayers()

	confirmCardsAreLegal(players, []Set{ori})

	loadDecklists()

	scores := calculateScore(players, decklists)

	outputScores(scores)
}

func outputScores(scores map[string]ScoreResult) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	var playerNames []string
	for playerName := range scores {
		playerNames = append(playerNames, playerName)
	}
	sort.Strings(playerNames)
	for _, playerName := range playerNames {
		score := scores[playerName]
		fmt.Fprintf(w, "== %s: \t%.1f ==\n", playerName, score.Total())
		printCardScores(w, score)
		fmt.Fprintln(w, "\t")
	}

	w.Flush()
}

type sortedScoreMap struct {
	m    map[Card]float64
	keys []Card
}

func newSortedScoreMap(scores ScoreResult) sortedScoreMap {
	sorted := sortedScoreMap{m: scores.CardScores}
	for key := range scores.CardScores {
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

func printCardScores(w io.Writer, score ScoreResult) {
	sorted := newSortedScoreMap(score)
	sort.Sort(&sorted)

	for cardName, cardScore := range sorted.m {
		fmt.Fprintf(w, "   %v: \t%.1f\n", cardName, cardScore)
	}
}

func buildPlayers() []Player {
	var playersMap map[string][]string
	readJsonFile("./owners.json", &playersMap)

	var players []Player
	for k, v := range playersMap {
		player := Player{Name: k, Cards: buildCards(v)}
		players = append(players, player)
	}
	return players
}

func buildCards(cardNames []string) []Card {
	var cards = make([]Card, 0)
	for _, cardName := range cardNames {
		cards = append(cards, Card(cardName))
	}
	return cards
}

var decklists []Decklist

func loadDecklists() {
	filepath.Walk(*decklistDir, loadDecklist)
}

func loadDecklist(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		loader := newDecklistLoader()
		decklist := loader.loadDecklist(path)
		decklists = append(decklists, decklist)
	}
	return nil
}
