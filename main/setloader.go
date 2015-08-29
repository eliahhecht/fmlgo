package main

type cardDto struct {
	Name string
}

type setDto struct {
	Code  string
	Cards []cardDto
}

func loadSet(setName string) Set {
	var sets map[string]setDto
	readJsonFile("./AllSets.json", &sets)
	return convertToSet(sets[setName])
}

func convertToSet(dto setDto) Set {
	numberOfCards := len(dto.Cards)
	cards := make([]Card, numberOfCards)

	for i, c := range dto.Cards {
		cards[i] = Card(c.Name)
	}

	return Set{dto.Code, cards}
}
