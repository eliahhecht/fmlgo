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

func loadAllCards() *CardCollection {
	sets := loadAllSets()

	cards := make(map[CardName]*Card)
	for _, set := range sets {
		for _, card := range set.Cards {
			if existingCard, ok := cards[card.Name]; ok {
				existingCard.SetCodes.Add(set.Code)
			} else {
				cards[card.Name] = card
			}
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
		s.Code = SetCode(dto.Code)

		cardsByName := make(map[CardName]*Card)
		doubleFacedCards := make([]*Card, 0)

		for _, c := range dto.Cards {
			newCard := makeCard(c, SetCode(s.Code))
			cardsByName[newCard.Name] = &newCard
			s.Cards = append(s.Cards, &newCard)

			if newCard.OtherSide != "" {
				doubleFacedCards = append(doubleFacedCards, &newCard)
			}
		}

		mergeDFCTypes(doubleFacedCards, cardsByName)

		return nil
	}
	return err
}

func makeCard(dto cardDto, setCode SetCode) Card {
	card := Card{Name: CardName(dto.Name), Types: mapset.NewSet(), SetCodes: mapset.NewSet()}
	card.SetCodes.Add(setCode)
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

	card.Rarity = Rarity(dto.Rarity)

	return card
}

func mergeDFCTypes(doubleFacedCards []*Card, cardsByName map[CardName]*Card) {
	for _, dfc := range doubleFacedCards {
		otherSide := cardsByName[dfc.OtherSide]
		dfc.Types = dfc.Types.Union(otherSide.Types)
	}
}
