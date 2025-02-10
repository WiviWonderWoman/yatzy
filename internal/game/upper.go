package game

import (
	"strconv"

	"gioui.org/widget"
)

type UpperScoreBox struct {
	Key       string
	Value     string
	DiceValue int
	Score     int
	Calculate bool
	Widget    *widget.Clickable
}

var UpperBoxes = []UpperScoreBox{
	{
		Key:       "Ones",
		Value:     "",
		DiceValue: 1,
		Score:     0,
	},
	{
		Key:       "Twos",
		Value:     "",
		DiceValue: 2,
		Score:     0,
	},
	{
		Key:       "Threes",
		Value:     "",
		DiceValue: 3,
		Score:     0,
	},
	{
		Key:       "Fours",
		Value:     "",
		DiceValue: 4,
		Score:     0,
	},
	{
		Key:       "Fives",
		Value:     "",
		DiceValue: 5,
		Score:     0,
	},
	{
		Key:       "Sixes",
		Value:     "",
		DiceValue: 6,
		Score:     0,
	},
}

func (u *UpperScoreBox) CountUpperValue(dices []Dice) {
	for _, d := range dices {
		if d.Value == u.DiceValue {
			u.Score += d.Value
		}
	}

	u.Value = strconv.Itoa(u.Score)
}
