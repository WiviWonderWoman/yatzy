package ui

import (
	"gioui.org/app"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/WiviWonderWoman/yatzy/internal/game"
)

// UI contains all ui-parts
type UI struct {
	Window     *app.Window
	theme      *material.Theme
	dices      []game.Dice
	rollButton *widget.Clickable
	rollsLeft  int
	upperBoxes []game.UpperScoreBox
	lowerBoxes []game.LowerScoreBox
	totalScore int
}

// NewUI initialize a new UI
func NewUI() *UI {
	// Initialize dices
	dices := make([]game.Dice, 5)
	for i := range dices {
		dices[i] = game.Dice{
			Widget:   new(widget.Clickable),
			Selected: false,
			Value:    game.GetRandomValue(),
		}
		dices[i].Key = game.GetKey(dices[i].Value)
	}

	// Initialize scoring boxes with clickable widgets
	upperBoxes := make([]game.UpperScoreBox, len(game.UpperBoxes))
	copy(upperBoxes, game.UpperBoxes)
	for i := range upperBoxes {
		upperBoxes[i].Widget = new(widget.Clickable)
	}

	lowerBoxes := make([]game.LowerScoreBox, len(game.LowerBoxes))
	copy(lowerBoxes, game.LowerBoxes)
	for i := range lowerBoxes {
		lowerBoxes[i].Widget = new(widget.Clickable)
	}

	return &UI{
		theme:      material.NewTheme(),
		dices:      dices,
		rollButton: new(widget.Clickable),
		rollsLeft:  3,
		upperBoxes: upperBoxes,
		lowerBoxes: lowerBoxes,
	}
}
