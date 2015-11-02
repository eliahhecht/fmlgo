package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func outputScores(players []*Player, allCards *CardCollection, standardLegalSets []SetCode) {
	csvFile, err := os.Create("out.csv")
	check(err)
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for _, player := range players {
		fmt.Printf("%s: %d\n", player.Name, player.TotalScore())
		writeCsvLine(csvWriter, string(player.Name))
		printCardScoresForPlayer(csvWriter, player.Cards)
		writeCsvLine(csvWriter, "Bench")
		printCardScoresForPlayer(csvWriter, player.Bench)
		writeCsvLine(csvWriter, "")
	}

	writeCsvLine(csvWriter, "")
	writeCsvLine(csvWriter, "Card scores by type")
	writeCsvLine(csvWriter, "")

	standardLegalCards := getStandardLegalCards(allCards, standardLegalSets)

	for _, cardType := range AllCardTypes {
		writeCsvLine(csvWriter, string(cardType))
		cardsForType := standardLegalCards.GetCardsOfType(cardType)
		printCardScoresForType(csvWriter, cardsForType)
		writeCsvLine(csvWriter, "")
	}
}

func writeCsvLine(w *csv.Writer, content string) {
	w.Write([]string{content})
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getStandardLegalCards(allCards *CardCollection, standardLegalSets []SetCode) *CardCollection {
	cards := make(map[CardName]*Card)
	for _, card := range allCards.CardsByName {
		for _, setCode := range standardLegalSets {
			if card.SetCodes.Contains(setCode) {
				cards[card.Name] = card
				break
			}
		}
	}

	return &CardCollection{cards}
}

func printCardScoresForPlayer(w *csv.Writer, cards []*Card) {
	printCardScores(w, cards, 10000, false)
}

func printCardScoresForType(w *csv.Writer, cards []*Card) {
	printCardScores(w, cards, 15, true)
}

func printCardScores(w *csv.Writer, cards []*Card, max int, includeOwner bool) {
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

func printCardScoreWithOwner(w *csv.Writer, card *Card) {
	owner := ""
	bench := ""
	if card.Ownership.OnBench {
		bench = ("(Bench)")
	}
	if card.IsOwned() {
		owner = string(card.Ownership.Owner)
	}

	if card.IsOwned() || card.Score > 0 {
		record := []string{string(card.Name), strconv.Itoa(card.Score), owner, bench}
		w.Write(record)
	}
}

func printCardScoreWithoutOwner(w *csv.Writer, card *Card) {
	w.Write([]string{string(card.Name), strconv.Itoa(card.Score)})
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
