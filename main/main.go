package main

func main() {
	var sets map[string]Set
	readJsonFile("./AllSets.json", &sets)
	ori := sets["ORI"]

	players := buildPlayers()

	confirmCardsAreLegal(players, []Set{ori})
}

func buildPlayers() []Player {
	var playersMap map[string][]string
	readJsonFile("./owners.json", &playersMap)

	players := make([]Player, 0)
	for k, v := range playersMap {
		player := Player{Name: k, Cards: buildCards(v)}
		players = append(players, player)
	}
	return players
}

func buildCards(cardNames []string) []Card {
	var cards = make([]Card, 0)
	for _, cardName := range cardNames {
		cards = append(cards, Card(cardName))
	}
	return cards
}
