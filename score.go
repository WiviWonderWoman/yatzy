package yatzy

func calculateUpperSum(upperBoxes []UpperScoreBox) int {
	sum := 0
	for _, u := range upperBoxes {
		sum += u.Score
	}

	if sum >= 63 {
		sum += 50
	}

	return sum
}

func calculateLowerSum(lowerBoxes []LowerScoreBox) int {
	sum := 0
	for _, u := range lowerBoxes {
		sum += u.Score
	}

	return sum
}

func calculateTotal(upperSum int, lowerSum int) int {
	return upperSum + lowerSum
}
