package main

import (
	"path"
	"regexp"
	"strconv"
	"strings"
)

// Decklist represents a decklist, containing a set of cards
type Decklist struct {
	Cards           []Card
	ScoreMultiplier float64
}

var cardLineRegex = regexp.MustCompile(`(\d+) (.*)`)

type fileLoader func(path string) []byte

type decklistLoader struct {
	loadFile fileLoader
}

func newDecklistLoader() decklistLoader {
	return decklistLoader{readFile}
}

func (l decklistLoader) loadDecklist(path string) Decklist {
	fileContents := string(l.loadFile(path))
	cards := parseDecklist(fileContents)
	multiplier := extractMultiplier(path)
	return Decklist{cards, multiplier}
}

// matches filenames like FooBar.1.3.txt, where 1.3 is the multiplier
var fileNameRegex = regexp.MustCompile(`[^\.]+\.([\d\.]+)\.txt`)

func extractMultiplier(filePath string) float64 {
	fileName := path.Base(filePath)
	matches := fileNameRegex.FindStringSubmatch(fileName)
	if len(matches) == 2 {
		multiplier, err := strconv.ParseFloat(matches[1], 64)
		if err == nil {
			return multiplier
		}
	}

	return 1.0
}

func parseDecklist(decklist string) []Card {
	lines := strings.Split(decklist, "\n")
	allCards := []Card{}
	for _, line := range lines {
		newCards, ok := parseLine(line)
		if ok {
			allCards = append(allCards, newCards...)
		}
	}
	return allCards
}

func parseLine(line string) (cards []Card, ok bool) {
	trimmed := strings.TrimSpace(line)
	matches := cardLineRegex.FindStringSubmatch(trimmed)

	if len(matches) == 3 { // matched a number and a card name
		numberOfCard, err := strconv.Atoi(matches[1])
		if err != nil {
			panic("Could not parse number of card")
		}

		card := Card(matches[2])
		cards := repeatCard(card, numberOfCard)

		return cards, true
	}
	return nil, false
}

func repeatCard(card Card, numberOfCard int) []Card {
	cards := make([]Card, numberOfCard)
	for i := 0; i < numberOfCard; i++ {
		cards[i] = card
	}
	return cards
}
