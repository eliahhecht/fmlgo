package main

import "encoding/json"

type cardDto struct {
	Name string
}

type setDto struct {
	Code  string
	Cards []cardDto
}

func loadSet(setName string) Set {
	var sets map[string]Set
	readJsonFile("./AllSets.json", &sets)
	return sets[setName]
}

// UnmarshalJSON decodes a Set from an MtgJson file
func (s *Set) UnmarshalJSON(b []byte) error {
	var dto setDto
	err := json.Unmarshal(b, &dto)
	if err == nil {
		s.Code = dto.Code
		s.Cards = make([]Card, len(dto.Cards))
		for i, c := range dto.Cards {
			s.Cards[i] = Card(c.Name)
		}
		return nil
	}
	return err
}
