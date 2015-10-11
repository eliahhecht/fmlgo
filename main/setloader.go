package main

import (
	"encoding/json"
	"github.com/deckarep/golang-set"
)

type cardDto struct {
	Name   string
	Rarity string
	Types  []string
	Names  []string
}

type setDto struct {
	Code  string
	Cards []cardDto
}

const cardDataFile string = "../data/AllSets.json"

func loadCards(setNames []string) *CardCollection {
	sets := loadAllSets()

	cards := make(map[CardName]*Card)
	for _, setName := range setNames {
		set := sets[setName]
		for _, card := range set.Cards {
			cards[card.Name] = card
		}
	}
	return &CardCollection{cards}
}

func loadAllSets() map[string]Set {
	var sets map[string]Set
	readJSONFile(cardDataFile, &sets)
	return sets
}

// UnmarshalJSON decodes a Set from an MtgJson file
func (s *Set) UnmarshalJSON(b []byte) error {
	var dto setDto
	err := json.Unmarshal(b, &dto)
	if err == nil {
		s.Code = dto.Code

		cardsByName := make(map[CardName]*Card)
		doubleFacedCards := make([]*Card, 0)

		for _, c := range dto.Cards {
			if c.Rarity != "Basic Land" {
				newCard := makeCard(c)
				cardsByName[newCard.Name] = &newCard
				s.Cards = append(s.Cards, &newCard)

				if newCard.OtherSide != "" {
					doubleFacedCards = append(doubleFacedCards, &newCard)
				}
			}
		}

		mergeDFCTypes(doubleFacedCards, cardsByName)

		return nil
	}
	return err
}

func makeCard(dto cardDto) Card {
	card := Card{Name: CardName(dto.Name), Types: mapset.NewSet()}
	for _, t := range dto.Types {
		card.Types.Add(CardType(t))
	}

	if len(dto.Names) > 1 {
		for _, name := range dto.Names {
			if name != dto.Name {
				card.OtherSide = CardName(name)
			}
		}
	}

	return card
}

func mergeDFCTypes(doubleFacedCards []*Card, cardsByName map[CardName]*Card) {
	for _, dfc := range doubleFacedCards {
		otherSide := cardsByName[dfc.OtherSide]
		dfc.Types = dfc.Types.Union(otherSide.Types)
	}
}
