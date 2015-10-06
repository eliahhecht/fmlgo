package main

// ScoreResult represents the score for a given player for a given week
type ScoreResult struct {
	CardScores  map[CardName]int
	BenchScores map[CardName]int
}

func newScoreResult() ScoreResult {
	return ScoreResult{CardScores: make(map[CardName]int), BenchScores: make(map[CardName]int)}
}

// OverallResult represents the overall scoring for a given week
type OverallResult struct {
	PlayerScores      map[string]ScoreResult
	UnownedCardScores map[CardName]int
}

func newOverallResult() OverallResult {
	return OverallResult{
		PlayerScores:      make(map[string]ScoreResult),
		UnownedCardScores: make(map[CardName]int)}
}

// Total is the player's total points for the week
func (sr ScoreResult) Total() int {
	total := 0
	for _, cardScoreValue := range sr.CardScores {
		total += cardScoreValue
	}
	return total
}

func calculateScore(players []Player, cardScores map[CardName]int) OverallResult {
	result := newOverallResult()
	unownedCards := copy(cardScores)
	result.UnownedCardScores = unownedCards

	for _, player := range players {
		result.PlayerScores[player.Name] = scorePlayer(player, cardScores, unownedCards)
	}

	return result
}

func copy(m map[CardName]int) map[CardName]int {
	copy := make(map[CardName]int)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}

func scorePlayer(player Player, cardScores map[CardName]int, unownedCards map[CardName]int) ScoreResult {
	playerResult := newScoreResult()
	transferScores(playerResult.CardScores, cardScores, player.Cards, unownedCards)
	transferScores(playerResult.BenchScores, cardScores, player.Bench, unownedCards)

	return playerResult
}

func transferScores(
	destinationMap map[CardName]int,
	sourceMap map[CardName]int,
	cards []Card,
	unownedCards map[CardName]int) {
	for _, card := range cards {
		destinationMap[card.Name] = sourceMap[card.Name]
		unownedCards[card.Name] = 0
	}
}
