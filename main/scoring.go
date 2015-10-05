package main

// ScoreResult represents the score for a given player for a given week
type ScoreResult struct {
	CardScores      map[Card]float64
	SideboardScores map[Card]float64
}

func newScoreResult() ScoreResult {
	return ScoreResult{CardScores: make(map[Card]float64), SideboardScores: make(map[Card]float64)}
}

// OverallResult represents the overall scoring for a given week
type OverallResult struct {
	PlayerScores      map[string]ScoreResult
	UnownedCardScores map[Card]float64
}

func newOverallResult() OverallResult {
	return OverallResult{
		PlayerScores:      make(map[string]ScoreResult),
		UnownedCardScores: make(map[Card]float64)}
}

func (sr *ScoreResult) addCard(c Card, scoreMultiplier float64) {
	sr.CardScores[c] += 1.0 * scoreMultiplier
}

// Total is the player's total points for the week
func (sr ScoreResult) Total() float64 {
	total := 0.0
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

func calculateScore(players []Player, cardScores map[Card]float64) OverallResult {
	result := newOverallResult()
	unownedCards := copy(cardScores)
	result.UnownedCardScores = unownedCards

	for _, player := range players {
		result.PlayerScores[player.Name] = scorePlayer(
			player, cardScores, unownedCards)
	}

	return result
}

func copy(m map[Card]float64) map[Card]float64 {
	copy := make(map[Card]float64)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}

func scorePlayer(player Player, cardScores map[Card]float64, unownedCards map[Card]float64) ScoreResult {
	playerResult := newScoreResult()
	transferScores(playerResult.CardScores, cardScores, player.Cards, unownedCards)
	transferScores(playerResult.SideboardScores, cardScores, player.Sideboard, unownedCards)

	return playerResult
}

func transferScores(
	destinationMap map[Card]float64,
	sourceMap map[Card]float64,
	cards []Card,
	unownedCards map[Card]float64) {
	for _, card := range cards {
		destinationMap[card] = sourceMap[card]
		unownedCards[card] = 0
	}
}
