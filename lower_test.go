package yatzy

import (
	"testing"

	"github.com/matryer/is"
)

func TestPair(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "valid pair",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
			},
			expected: 10,
		},
		{
			name: "invalid pair - different values",
			dices: []Dice{
				{Value: 5},
				{Value: 4},
			},
			expected: 0,
		},
		{
			name: "invalid pair - wrong length",
			dices: []Dice{
				{Value: 5},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := pair(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestTwoPair(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "valid two pairs",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
				{Value: 6},
				{Value: 6},
			},
			expected: 22,
		},
		{
			name: "invalid - not pairs",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
				{Value: 6},
				{Value: 4},
			},
			expected: 0,
		},
		{
			name: "invalid - wrong length",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
				{Value: 6},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := twoPair(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestThreesome(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "valid three of a kind",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
				{Value: 5},
			},
			expected: 15,
		},
		{
			name: "invalid - different values",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
				{Value: 6},
			},
			expected: 0,
		},
		{
			name: "invalid - wrong length",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := threesome(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestFoursome(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "valid four of a kind",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
				{Value: 5},
				{Value: 5},
			},
			expected: 20,
		},
		{
			name: "invalid - different values",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
				{Value: 5},
				{Value: 6},
			},
			expected: 0,
		},
		{
			name: "invalid - wrong length",
			dices: []Dice{
				{Value: 5},
				{Value: 5},
				{Value: 5},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := foursome(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestSmallStraight(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "valid small straight",
			dices: []Dice{
				{Value: 1},
				{Value: 2},
				{Value: 3},
				{Value: 4},
				{Value: 5},
			},
			expected: 15,
		},
		{
			name: "invalid - wrong sequence",
			dices: []Dice{
				{Value: 2},
				{Value: 3},
				{Value: 4},
				{Value: 5},
				{Value: 6},
			},
			expected: 0,
		},
		{
			name: "invalid - wrong length",
			dices: []Dice{
				{Value: 1},
				{Value: 2},
				{Value: 3},
				{Value: 4},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := small(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestLargeStraight(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "valid large straight",
			dices: []Dice{
				{Value: 2},
				{Value: 3},
				{Value: 4},
				{Value: 5},
				{Value: 6},
			},
			expected: 20,
		},
		{
			name: "invalid - wrong sequence",
			dices: []Dice{
				{Value: 1},
				{Value: 2},
				{Value: 3},
				{Value: 4},
				{Value: 5},
			},
			expected: 0,
		},
		{
			name: "invalid - wrong length",
			dices: []Dice{
				{Value: 2},
				{Value: 3},
				{Value: 4},
				{Value: 5},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := large(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestFullHouse(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "valid full house",
			dices: []Dice{
				{Value: 6},
				{Value: 6},
				{Value: 6},
				{Value: 5},
				{Value: 5},
			},
			expected: 28,
		},
		{
			name: "invalid - wrong combination",
			dices: []Dice{
				{Value: 6},
				{Value: 6},
				{Value: 6},
				{Value: 5},
				{Value: 4},
			},
			expected: 0,
		},
		{
			name: "invalid - wrong length",
			dices: []Dice{
				{Value: 6},
				{Value: 6},
				{Value: 6},
				{Value: 5},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := fullHouse(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestChance(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "sum all dices",
			dices: []Dice{
				{Value: 6},
				{Value: 5},
				{Value: 4},
				{Value: 3},
				{Value: 2},
			},
			expected: 20,
		},
		{
			name:     "empty dice array",
			dices:    []Dice{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := sumDices(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestYatzy(t *testing.T) {
	tests := []struct {
		name     string
		dices    []Dice
		expected int
	}{
		{
			name: "valid yatzy",
			dices: []Dice{
				{Value: 6},
				{Value: 6},
				{Value: 6},
				{Value: 6},
				{Value: 6},
			},
			expected: 30,
		},
		{
			name: "invalid - different values",
			dices: []Dice{
				{Value: 6},
				{Value: 6},
				{Value: 5},
				{Value: 6},
				{Value: 6},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := yatzy(tt.dices)
			is.Equal(result, tt.expected)
		})
	}
}

func TestCheckLength(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		length   int
		expected bool
	}{
		{
			name:     "correct length",
			slice:    []int{1, 2, 3},
			length:   3,
			expected: true,
		},
		{
			name:     "incorrect length",
			slice:    []int{1, 2},
			length:   3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			result := checkLength(tt.slice, tt.length)
			is.Equal(result, tt.expected)
		})
	}
}
