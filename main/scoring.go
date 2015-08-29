package main

// ScoreResult represents the score for a given player for a given week
type ScoreResult struct {
	Total      float64
	CardScores map[Card]float64
}

func newScoreResult() ScoreResult {
	return ScoreResult{CardScores: make(map[Card]float64)}
}

func (sr *ScoreResult) addCard(c Card) {
	sr.CardScores[c] += 1.0
	sr.Total += 1.0 //ehtodo make total a func
}

func (p *Player) hasCard(needle Card) bool {
	for _, playerCard := range p.Cards {
		if playerCard == needle {
			return true
		}
	}
	return false
}

func calculateScore(players []Player, decklists []Decklist) map[string]ScoreResult {
	result := make(map[string]ScoreResult)

	for _, player := range players {
		result[player.Name] = scorePlayer(player, decklists)
	}

	return result
}

func scorePlayer(player Player, decklists []Decklist) ScoreResult {
	playerResult := newScoreResult()
	for _, decklist := range decklists {
		for _, decklistCard := range decklist.Cards {
			if player.hasCard(decklistCard) {
				playerResult.addCard(decklistCard)
			}
		}
	}

	return playerResult
}
