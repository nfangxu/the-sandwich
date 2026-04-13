package game

import "sort"

func EvaluateHand(cards []Card) HandResult {
	if len(cards) != 3 {
		return HandResult{Type: HighCard}
	}

	// Separate normal cards and jokers
	var normalCards []Card
	jokerCount := 0
	for _, c := range cards {
		if c.Suit == Joker {
			jokerCount++
		} else {
			normalCards = append(normalCards, c)
		}
	}

	// Sort normal cards descending
	sort.Slice(normalCards, func(i, j int) bool {
		return normalCards[i].Rank > normalCards[j].Rank
	})

	// Logic for 3 cards
	isFlush := len(normalCards) > 0
	for i := 1; i < len(normalCards); i++ {
		if normalCards[i].Suit != normalCards[0].Suit {
			isFlush = false
			break
		}
	}

	isStraight := false
	if len(normalCards) > 1 {
		gap := int(normalCards[0].Rank - normalCards[len(normalCards)-1].Rank)
		if gap <= 2+jokerCount && len(normalCards) == 3 && normalCards[0].Rank != normalCards[1].Rank && normalCards[1].Rank != normalCards[2].Rank {
            isStraight = true
        } else if len(normalCards) == 2 && normalCards[0].Rank != normalCards[1].Rank && gap <= 1+jokerCount {
			isStraight = true
		} else if len(normalCards) == 1 {
            isStraight = true
        }
	} else {
		isStraight = true
	}

    isLeopard := false
    if len(normalCards) <= 1 {
        isLeopard = true
    } else if len(normalCards) == 2 && normalCards[0].Rank == normalCards[1].Rank {
        isLeopard = true
    } else if len(normalCards) == 3 && normalCards[0].Rank == normalCards[1].Rank && normalCards[1].Rank == normalCards[2].Rank {
        isLeopard = true
    }

	// Determine best hand
	if isLeopard {
		return HandResult{Type: Leopard, Cards: cards}
	}
	if isFlush && isStraight {
		return HandResult{Type: StraightFlush, Cards: cards}
	}
	if isFlush {
		return HandResult{Type: Flush, Cards: cards}
	}
	if isStraight {
		return HandResult{Type: Straight, Cards: cards}
	}

	// Pair check
	isPair := false
	if jokerCount > 0 {
		isPair = true
	} else if len(normalCards) == 3 && (normalCards[0].Rank == normalCards[1].Rank || normalCards[1].Rank == normalCards[2].Rank) {
		isPair = true
	}

	if isPair {
		return HandResult{Type: Pair, Cards: cards}
	}

	return HandResult{Type: HighCard, Cards: cards}
}
