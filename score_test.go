package yatzy

import (
	"testing"

	"github.com/matryer/is"
)

func TestCalculateUpperSum(t *testing.T) {
	tests := []struct {
		name       string
		upperBoxes []UpperScoreBox
		expected   int
	}{
		{
			name: "sum less than 63",
			upperBoxes: []UpperScoreBox{
				{Score: 10},
				{Score: 20},
			},
			expected: 30,
		},
		{
			name: "sum equals 63",
			upperBoxes: []UpperScoreBox{
				{Score: 30},
				{Score: 33},
			},
			expected: 113, // 63 + 50 bonus
		},
		{
			name: "sum greater than 63",
			upperBoxes: []UpperScoreBox{
				{Score: 40},
				{Score: 40},
			},
			expected: 130, // 80 + 50 bonus
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			result := calculateUpperSum(tt.upperBoxes)
			is.Equal(result, tt.expected)
		})
	}
}

func TestCalculateLowerSum(t *testing.T) {
	tests := []struct {
		name       string
		lowerBoxes []LowerScoreBox
		expected   int
	}{
		{
			name: "simple sum",
			lowerBoxes: []LowerScoreBox{
				{Score: 10},
				{Score: 20},
			},
			expected: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			result := calculateLowerSum(tt.lowerBoxes)
			is.Equal(result, tt.expected)
		})
	}
}

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		name     string
		upperSum int
		lowerSum int
		expected int
	}{
		{
			name:     "simple total",
			upperSum: 50,
			lowerSum: 100,
			expected: 150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			result := calculateTotal(tt.upperSum, tt.lowerSum)
			is.Equal(result, tt.expected)
		})
	}
}
