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

var decklistDir string

func main() {
	parseFlags()
	flag.Parse()

	legalCards := loadCards([]string{"KTK", "FRF", "DTK", "ORI", "BFZ", "EXP"})

	players := buildPlayers(legalCards)

	confirmCardsAreLegal(players, legalCards)

	loadDecklists()
	ScoreCards(decklists, legalCards)

	calculateScore(players, legalCards)

	outputScores(players, legalCards)
}

func parseFlags() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		panic("Must be called with exactly one argument (path to decklists dir)")
	}
	decklistDir = flag.Arg(0)
}

func outputScores(players []*Player, allCards *CardCollection) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	for _, player := range players {
		fmt.Fprintf(w, "== %s: \t%d ==\n", player.Name, player.TotalScore())
		printAllCardScores(w, player.Cards)
		fmt.Fprintln(w, "  Bench:")
		printAllCardScores(w, player.Bench)
		fmt.Fprintln(w, "\t")
	}

	fmt.Fprintln(w, "Card scores by type: ")

	for _, cardType := range AllCardTypes {
		fmt.Fprintf(w, "\n%s:\n", cardType)
		cardsForType := allCards.GetCardsOfType(cardType)
		printCardScores(w, cardsForType, 10)
	}

	w.Flush()
}


type ByScore []*Card

func (bs ByScore) Len() int {
	return len(bs)
}

func (bs ByScore) Less(i, j int) bool {
	return bs[i].Score < bs[j].Score
}

func (bs ByScore) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func printAllCardScores(w io.Writer, cards []*Card) {
	printCardScores(w, cards, 10000)
}

func printCardScores(w io.Writer, cards []*Card, max int) {
	cardsByScore := ByScore(cards)
	sort.Sort(sort.Reverse(&cardsByScore))

	for i, card := range cardsByScore {
		fmt.Fprintf(w, "   %v \t%d\n", card.Name, card.Score)

		if i >= max {
			break
		}
	}
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
