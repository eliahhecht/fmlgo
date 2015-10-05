package main

// ScoreResult represents the score for a given player for a given week
type ScoreResult struct {
	CardScores      map[Card]int
	SideboardScores map[Card]int
}

func newScoreResult() ScoreResult {
	return ScoreResult{CardScores: make(map[Card]int), SideboardScores: make(map[Card]int)}
}

// OverallResult represents the overall scoring for a given week
type OverallResult struct {
	PlayerScores      map[string]ScoreResult
	UnownedCardScores map[Card]int
}

func newOverallResult() OverallResult {
	return OverallResult{
		PlayerScores:      make(map[string]ScoreResult),
		UnownedCardScores: make(map[Card]int)}
}

// Total is the player's total points for the week
func (sr ScoreResult) Total() int {
	total := 0
	for _, cardScoreValue := range sr.CardScores {
		total += cardScoreValue
	}
	return total
}

func (p *Player) hasCard(needle Card) bool {
	for _, playerCard := range p.Cards {
		if playerCard == needle {
			return true
		}
	}
	return false
}

func calculateScore(players []Player, cardScores map[Card]int) OverallResult {
	result := newOverallResult()
	unownedCards := copy(cardScores)
	result.UnownedCardScores = unownedCards

	for _, player := range players {
		result.PlayerScores[player.Name] = scorePlayer(
			player, cardScores, unownedCards)
	}

	return result
}

func copy(m map[Card]int) map[Card]int {
	copy := make(map[Card]int)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}

func scorePlayer(player Player, cardScores map[Card]int, unownedCards map[Card]int) ScoreResult {
	playerResult := newScoreResult()
	transferScores(playerResult.CardScores, cardScores, player.Cards, unownedCards)
	transferScores(playerResult.SideboardScores, cardScores, player.Sideboard, unownedCards)

	return playerResult
}

func transferScores(
	destinationMap map[Card]int,
	sourceMap map[Card]int,
	cards []Card,
	unownedCards map[Card]int) {
	for _, card := range cards {
		destinationMap[card] = sourceMap[card]
		unownedCards[card] = 0
	}
}
