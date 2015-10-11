package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
)

func outputScores(players []*Player, allCards *CardCollection) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	for _, player := range players {
		fmt.Fprintf(w, "== %s: \t%d ==\n", player.Name, player.TotalScore())
		printCardScoresForPlayer(w, player.Cards)
		fmt.Fprintln(w, "  Bench: \t")
		printCardScoresForPlayer(w, player.Bench)
		fmt.Fprintln(w, "\t")
	}

	fmt.Fprintln(w, "Card scores by type: \t\t")

	for _, cardType := range AllCardTypes {
		fmt.Fprintf(w, "\t\t\t\n%s: \t\t\n", cardType)
		cardsForType := allCards.GetCardsOfType(cardType)
		printCardScoresForType(w, cardsForType)
	}

	w.Flush()
}

func printCardScoresForPlayer(w io.Writer, cards []*Card) {
	printCardScores(w, cards, 10000, false)
}

func printCardScoresForType(w io.Writer, cards []*Card) {
	printCardScores(w, cards, 15, true)
}

func printCardScores(w io.Writer, cards []*Card, max int, includeOwner bool) {
	cardsByScore := ByScore(cards)
	sort.Sort(sort.Reverse(&cardsByScore))

	for i, card := range cardsByScore {
		if includeOwner {
			printCardScoreWithOwner(w, card)
		} else {
			printCardScoreWithoutOwner(w, card)
		}

		if i >= max {
			break
		}
	}
}

func printCardScoreWithOwner(w io.Writer, card *Card) {
	owner := ""
	bench := ""
	if card.Ownership.OnBench {
		bench = (" (Bench)")
	}
	if card.IsOwned {
		owner = string(card.Ownership.Owner)
	}

	if owner == "" && card.Score == 0 {
		continue
	}

	fmt.Fprintf(w, "   %v \t%d\t%v%v\n", card.Name, card.Score, owner, bench)
}

func printCardScoreWithoutOwner(w io.Writer, card *Card) {
	fmt.Fprintf(w, "   %v \t%d\n", card.Name, card.Score)
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
