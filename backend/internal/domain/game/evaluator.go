package game

func EvaluateHand(cards []Card) HandResult {
	if len(cards) < 3 {
		return HandResult{Type: HighCard, Score: 0}
	}

	var nonJokers []Card
	jokerCount := 0
	for _, c := range cards {
		if c.Suit == Joker {
			jokerCount++
		} else {
			nonJokers = append(nonJokers, c)
		}
	}

	if jokerCount >= 3 {
		return HandResult{Type: Leopard, Score: 500 + int(nonJokers[0].Rank)}
	}

	hasPair := false
	hasThreeOfAKind := false
	hasStraight := false
	hasFlush := true

	rankCount := make(map[Rank]int)
	suitCount := make(map[Suit]int)

	for _, c := range nonJokers {
		rankCount[c.Rank]++
		suitCount[c.Suit]++
		if suitCount[c.Suit] > 1 {
			hasFlush = false
		}
	}

	for _, count := range rankCount {
		if count == 2 {
			hasPair = true
		}
		if count == 3 {
			hasThreeOfAKind = true
		}
	}

	if len(nonJokers) >= 3 {
		var ranks []Rank
		for _, c := range nonJokers {
			ranks = append(ranks, c.Rank)
		}
		hasStraight = isConsecutive(ranks)
	}

	if jokerCount >= 2 {
		if hasPair || hasThreeOfAKind {
			return HandResult{Type: Leopard, Score: 500 + int(getHighestRank(nonJokers))}
		}
		if hasStraight && hasFlush {
			return HandResult{Type: StraightFlush, Score: 400 + int(getHighestRank(nonJokers))}
		}
		if hasStraight {
			return HandResult{Type: Straight, Score: 300 + int(getHighestRank(nonJokers))}
		}
		if hasFlush {
			return HandResult{Type: Flush, Score: 200 + int(getHighestRank(nonJokers))}
		}
		if hasPair {
			return HandResult{Type: Pair, Score: 100 + int(getHighestRank(nonJokers))}
		}
	}

	if jokerCount >= 1 {
		if hasThreeOfAKind {
			return HandResult{Type: Leopard, Score: 500 + int(getHighestRank(nonJokers))}
		}
		if hasFlush {
			return HandResult{Type: Flush, Score: 200 + int(getHighestRank(nonJokers))}
		}
		if hasStraight {
			return HandResult{Type: Straight, Score: 300 + int(getHighestRank(nonJokers))}
		}
		if hasPair {
			return HandResult{Type: Pair, Score: 100 + int(getHighestRank(nonJokers))}
		}
	}

	if hasThreeOfAKind {
		return HandResult{Type: Leopard, Score: 500 + int(getHighestRank(nonJokers))}
	}
	if hasStraight && hasFlush {
		return HandResult{Type: StraightFlush, Score: 400 + int(getHighestRank(nonJokers))}
	}
	if hasStraight {
		return HandResult{Type: Straight, Score: 300 + int(getHighestRank(nonJokers))}
	}
	if hasFlush {
		return HandResult{Type: Flush, Score: 200 + int(getHighestRank(nonJokers))}
	}
	if hasPair {
		return HandResult{Type: Pair, Score: 100 + int(getHighestRank(nonJokers))}
	}

	return HandResult{Type: HighCard, Score: int(getHighestRank(nonJokers))}
}

func isConsecutive(ranks []Rank) bool {
	if len(ranks) < 3 {
		return false
	}
	for i := 0; i < len(ranks)-2; i++ {
		if ranks[i]+1 == ranks[i+1] && ranks[i+1]+1 == ranks[i+2] {
			return true
		}
	}
	return false
}

func getHighestRank(cards []Card) Rank {
	max := Rank(0)
	for _, c := range cards {
		if c.Rank > max {
			max = c.Rank
		}
	}
	return max
}
