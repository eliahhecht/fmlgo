package main

// ScoreResult represents the score for a given player for a given week
type ScoreResult struct {
	CardScores map[Card]float64
}

func newScoreResult() ScoreResult {
	return ScoreResult{CardScores: make(map[Card]float64)}
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

func calculateScore(players []Player, cardScores map[Card]float64) map[string]ScoreResult {
	result := make(map[string]ScoreResult)

	for _, player := range players {
		result[player.Name] = scorePlayer(player, cardScores)
	}

	return result
}

func scorePlayer(player Player, cardScores map[Card]float64) ScoreResult {
	playerResult := newScoreResult()
	for _, playerCard := range player.Cards {
		playerResult.CardScores[playerCard] = cardScores[playerCard]
	}
	return playerResult
}
