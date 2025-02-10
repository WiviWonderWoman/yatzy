package ui

import (
	"gioui.org/app"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/WiviWonderWoman/yatzy/internal/game"
)

// First, let's extend the UI struct to include scoring
type UI struct {
	Window     *app.Window
	theme      *material.Theme
	dices      []game.Dice
	rollButton *widget.Clickable
	rollsLeft  int
	upperBoxes []game.UpperScoreBox // Add upper section scoring
	lowerBoxes []game.LowerScoreBox // Add lower section scoring
	totalScore int                  // Track total score
}

// Update NewUI to initialize scoring
func NewUI() *UI {
	// Initialize dice (existing code)
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
