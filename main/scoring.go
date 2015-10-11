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
	PlayerScores      map[PlayerName]ScoreResult
	UnownedCardScores map[CardName]int
}

func (r *OverallResult) GetScoresForType(cardType CardType) map[CardName]int {
	return nil
}

func newOverallResult() OverallResult {
	return OverallResult{
		PlayerScores:      make(map[PlayerName]ScoreResult),
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

func calculateScore(players []*Player, cards *CardCollection) OverallResult {
	result := newOverallResult()
	tagOwners(players, cards)
	//	result.UnownedCardScores = unownedCards

	return result
}

func tagOwners(players []*Player, cards *CardCollection) {
	for _, player := range players {
		for _, card := range player.Cards {
			cards.GetCard(card.Name).Ownership = OwnershipTag{Owner: player.Name, OnBench: false}
		}
		for _, card := range player.Cards {
			cards.GetCard(card.Name).Ownership = OwnershipTag{Owner: player.Name, OnBench: true}
		}
	}

}

