package game

func CalculateUpperSum(upperBoxes []UpperScoreBox) int {
	sum := 0
	for _, u := range upperBoxes {
		sum += u.Score
	}

	return sum
}

func CalculateLowerSum(lowerBoxes []LowerScoreBox) int {
	sum := 0
	for _, u := range lowerBoxes {
		sum += u.Score
	}

	return sum
}

func CalculateTotal(upperSum int, lowerSum int) int {
	return upperSum + lowerSum
}
