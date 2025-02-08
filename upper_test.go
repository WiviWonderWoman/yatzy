package yatzy

import (
	"testing"

	"github.com/matryer/is"
)

func TestUpperScoreBoxCountUpperValue(t *testing.T) {
	tests := []struct {
		name     string
		box      UpperScoreBox
		dices    []Dice
		expected string
		expScore int
	}{
		{
			name: "count ones",
			box: UpperScoreBox{
				DiceValue: 1,
			},
			dices: []Dice{
				{Value: 1},
				{Value: 1},
				{Value: 2},
			},
			expected: "2",
			expScore: 2,
		},
		{
			name: "count sixes",
			box: UpperScoreBox{
				DiceValue: 6,
			},
			dices: []Dice{
				{Value: 6},
				{Value: 6},
				{Value: 6},
			},
			expected: "18",
			expScore: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			tt.box.countUpperValue(tt.dices)
			is.Equal(tt.box.Value, tt.expected)
			is.Equal(tt.box.Score, tt.expScore)
		})
	}
}
