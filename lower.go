package yatzy

import (
	"sort"

	"gioui.org/widget"
)

type LowerScoreBox struct {
	Key           string
	Value         string
	Score         int
	Calculate     bool
	Widget        *widget.Clickable
	CalculateFunc func(dices []Dice) int
}

var LowerBoxes = []LowerScoreBox{
	{
		Key:           "One Pair",
		Value:         "",
		Score:         0,
		CalculateFunc: pair,
	},
	{
		Key:           "Two Pairs",
		Value:         "",
		Score:         0,
		CalculateFunc: twoPair,
	},
	{
		Key:           "Three of a Kind",
		Value:         "",
		Score:         0,
		CalculateFunc: threesome,
	},
	{
		Key:           "Four of a Kind",
		Value:         "",
		Score:         0,
		CalculateFunc: foursome,
	},
	{
		Key:           "Small Straight",
		Value:         "",
		Score:         0,
		CalculateFunc: small,
	},
	{
		Key:           "Large Straight",
		Value:         "",
		Score:         0,
		CalculateFunc: large,
	},
	{
		Key:           "Full House",
		Value:         "",
		Score:         0,
		CalculateFunc: fullHouse,
	},
	{
		Key:           "Chance",
		Value:         "",
		Score:         0,
		CalculateFunc: sumDices,
	},
	{
		Key:           "Yatzy",
		Value:         "",
		Score:         0,
		CalculateFunc: yatzy,
	},
}

func pair(dices []Dice) int {
	// must be two equal dices
	if !checkLength(dices, 2) || dices[0].Value != dices[1].Value {
		return 0
	}

	return sumDices(dices)
}

func twoPair(dices []Dice) int {
	// must be four dices
	if !checkLength(dices, 4) {
		return 0
	}

	// extract values from dices
	values := make([]int, 0, 4)
	for _, d := range dices {
		values = append(values, d.Value)
	}

	// sort values
	sort.Ints(values)

	// find two different pairs
	if values[0] != values[1] && values[2] != values[3] {
		return 0
	}

	if values[1] == values[2] {
		return 0
	}

	return sumDices(dices)
}

func threesome(dices []Dice) int {
	// must be three dices
	if !checkLength(dices, 3) {
		return 0
	}

	// extract valid values from dices
	values := make([]int, 0, 3)
	for i, d := range dices {
		if i == 0 || d.Value == values[i-1] {
			values = append(values, d.Value)
		}
	}

	// must be three equal values
	if !checkLength(values, 3) {
		return 0
	}

	return sumDices(dices)
}

func foursome(dices []Dice) int {
	// must be four dices
	if !checkLength(dices, 4) {
		return 0
	}

	// extract valid values from dices
	values := make([]int, 0, 4)
	for i, d := range dices {
		if i == 0 || d.Value == values[i-1] {
			values = append(values, d.Value)
		}
	}

	// must be four equal values
	if !checkLength(values, 4) {
		return 0
	}

	return sumDices(dices)
}

func small(dices []Dice) int {
	sum := 0
	// must be five dices
	if !checkLength(dices, 5) {
		return sum
	}

	return checkStraight(dices, 1)
}

func large(dices []Dice) int {
	sum := 0
	// must be five dices
	if !checkLength(dices, 5) {
		return sum
	}

	return checkStraight(dices, 2)
}

func checkStraight(dices []Dice, add int) int {
	sum := 0

	// extract values from dices
	values := make([]int, 0, 5)
	for _, d := range dices {
		values = append(values, d.Value)
	}

	// sort values
	sort.Ints(values)

	// check values
	for i, v := range values {
		if v != i+add {
			return 0
		}
		sum += v
	}

	return sum
}

func fullHouse(dices []Dice) int {
	// must be five dices
	if !checkLength(dices, 5) {
		return 0
	}
	// Count occurrences of each dice value
	counts := make(map[int]int)
	for _, d := range dices {
		counts[d.Value]++
	}

	// Check if we have exactly two different values
	if len(counts) != 2 {
		return 0
	}

	// Check if one value appears 3 times and another 2 times
	hasThree := false
	hasTwo := false
	for _, count := range counts {
		if count == 3 {
			hasThree = true
		}
		if count == 2 {
			hasTwo = true
		}
	}

	if !hasThree && !hasTwo {
		return 0
	}

	return sumDices(dices)
}

func yatzy(dices []Dice) int {
	// extract valid values from dices
	value := dices[0].Value
	for _, d := range dices {
		if d.Value != value {
			return 0
		}
	}

	return 50
}

func checkLength[T any](dices []T, l int) bool {
	return len(dices) == l
}

func sumDices(dices []Dice) int {
	sum := 0

	for _, d := range dices {
		sum += d.Value
	}

	return sum
}
