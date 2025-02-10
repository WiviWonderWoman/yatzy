package ui

import (
	"fmt"

	"gioui.org/app"
	"gioui.org/op"
	"github.com/WiviWonderWoman/yatzy/internal/game"
)

// rollDice rolls all unlocked dice
func (ui *UI) rollDice() {
	for i := range ui.dices {
		if !ui.dices[i].Selected {
			ui.dices[i].Value = game.GetRandomValue()
			ui.dices[i].Key = game.GetKey(ui.dices[i].Value)
		}
	}
}

// Add method to reset turn
func (ui *UI) resetTurn() {
	ui.rollsLeft = 3
	for i := range ui.dices {
		ui.dices[i].Selected = false
	}
	ui.rollDice()
}

// Add method to update total score
func (ui *UI) updateScore() {
	upperSum := game.CalculateUpperSum(ui.upperBoxes)
	lowerSum := game.CalculateLowerSum(ui.lowerBoxes)

	// Add bonus if upper section sum is 63 or more
	bonus := 0
	if upperSum >= 63 {
		bonus = 50
	}

	ui.totalScore = upperSum + bonus + lowerSum
}

// Helper function for integer max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Update handleEvents to handle score box clicks
func (ui *UI) HandleEvents() error {
	var ops op.Ops

	for {
		evt := ui.Window.Event()

		switch e := evt.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Handle upper section scoring
			for i := range ui.upperBoxes {
				if ui.upperBoxes[i].Widget.Clicked(gtx) && !ui.upperBoxes[i].Calculate {
					ui.upperBoxes[i].CountUpperValue(ui.dices)
					ui.upperBoxes[i].Calculate = true
					ui.resetTurn()
				}
			}

			// Handle lower section scoring
			for i := range ui.lowerBoxes {
				if ui.lowerBoxes[i].Widget.Clicked(gtx) && !ui.lowerBoxes[i].Calculate {
					ui.lowerBoxes[i].Score = ui.lowerBoxes[i].CalculateFunc(ui.dices)
					ui.lowerBoxes[i].Value = fmt.Sprintf("%d", ui.lowerBoxes[i].Score)
					ui.lowerBoxes[i].Calculate = true
					ui.resetTurn()
				}
			}

			// Handle dice selection
			for i := range ui.dices {
				if ui.dices[i].Widget.Clicked(gtx) {
					ui.dices[i].Selected = !ui.dices[i].Selected
				}
			}

			// Handle roll button
			if ui.rollButton.Clicked(gtx) && ui.rollsLeft > 0 {
				ui.rollDice()
				ui.rollsLeft--
			}

			// Update total score
			ui.updateScore()

			ui.layout(gtx)
			e.Frame(gtx.Ops)

		case app.DestroyEvent:
			return e.Err
		}
	}
}
