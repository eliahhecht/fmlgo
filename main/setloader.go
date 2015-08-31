package main

import "encoding/json"

type cardDto struct {
	Name   string
	Rarity string
}

type setDto struct {
	Code  string
	Cards []cardDto
}

const cardDataFile string = "../data/AllSets.json"

func loadSet(setName string) Set {
	var sets map[string]Set
	readJsonFile(cardDataFile, &sets)
	return sets[setName]
}

// UnmarshalJSON decodes a Set from an MtgJson file
func (s *Set) UnmarshalJSON(b []byte) error {
	var dto setDto
	err := json.Unmarshal(b, &dto)
	if err == nil {
		s.Code = dto.Code
		for _, c := range dto.Cards {
			if c.Rarity != "Basic Land" {
				s.Cards = append(s.Cards, Card(c.Name))
			}
		}
		return nil
	}
	return err
}
