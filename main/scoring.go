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

func (r *OverallResult) GetScoresForType(cardType CardType) map[CardName]int {
	return nil
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

func calculateScore(players []Player, cards *CardCollection) OverallResult {
	result := newOverallResult()
	tagOwners(players, cards)
	//	result.UnownedCardScores = unownedCards

	for _, player := range players {
		result.PlayerScores[player.Name] = scorePlayer(player, cards)
	}

	return result
}

func tagOwners(players []Player, cards *CardCollection) {

}

func scorePlayer(player Player, cardScores *CardCollection) ScoreResult {
	playerResult := newScoreResult()

	transferScores(playerResult.CardScores, cardScores, player.Cards)
	transferScores(playerResult.BenchScores, cardScores, player.Bench)

	return playerResult
}

func transferScores(
	destinationMap map[CardName]int,
	allCards *CardCollection,
	cardsToTransfer []*Card) {
	for _, card := range cardsToTransfer {
		destinationMap[card.Name] = allCards.GetCard(card.Name).Score
	}
}
