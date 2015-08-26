package main

import (
	"regexp"
	"strconv"
	"strings"
)

// Decklist represents a decklist, containing a set of cards
type Decklist struct {
	Cards []Card
}

var cardLineRegex = regexp.MustCompile(`(\d+) (.*)`)

func loadDecklist(path string) Decklist {
	fileContents := string(readFile(path))
	return parseDecklist(fileContents)
}

func parseDecklist(decklist string) Decklist {
	lines := strings.Split(decklist, "\n")
	cards := []Card{}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		matches := cardLineRegex.FindStringSubmatch(trimmed)

		if len(matches) == 3 { // matched a number and a card name
			numberOfCard, err := strconv.Atoi(matches[1])
			if err != nil {
				panic("Could not parse number of card")
			}

			card := Card(matches[2])

			for i := 0; i < numberOfCard; i++ {
				cards = append(cards, card)
			}
		}
	}
	return Decklist{cards}
}
