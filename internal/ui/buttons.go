package ui

import (
	"fmt"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/WiviWonderWoman/yatzy/internal/game"
)

// createScoreButton creates the buttons for scores
func (ui *UI) createScoreButton(
	key string,
	value string,
	calculated bool,
	clickable *widget.Clickable,
) material.ButtonStyle {
	btn := material.Button(ui.theme, clickable, fmt.Sprintf("%s: %s", key, value))

	if calculated {
		// Already scored - disable the button
		btn.Background = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		btn.Color = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	} else {
		// Available for scoring
		btn.Background = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
		btn.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	}

	// Common button styling
	btn.TextSize = unit.Sp(14)
	btn.CornerRadius = unit.Dp(4)
	btn.Inset = layout.Inset{
		Top:    unit.Dp(8),
		Bottom: unit.Dp(8),
		Left:   unit.Dp(12),
		Right:  unit.Dp(12),
	}
	return btn
}

// createDiceButton creates a styled button for a dice
func (ui *UI) createDiceButton(dice *game.Dice) material.ButtonStyle {
	btn := material.Button(ui.theme, dice.Widget, dice.Key)

	if dice.Selected {
		// Style for locked dice (light green, semi-transparent)
		btn.Background = color.NRGBA{R: 100, G: 200, B: 100, A: 100}
		btn.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	} else {
		// Style for unlocked dice (gray)
		btn.Background = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
		btn.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	}

	// Common button styling
	btn.Inset.Top = unit.Dp(0)
	btn.Inset.Bottom = unit.Dp(10)
	btn.TextSize = unit.Sp(15)
	btn.Font.Weight = font.Black
	btn.CornerRadius = unit.Dp(8)

	return btn
}

// createRollButton creates a styled button for rolling dice
func (ui *UI) createRollButton(clickable *widget.Clickable) material.ButtonStyle {
	rollText := fmt.Sprintf("ROLL (%d)", ui.rollsLeft)
	btn := material.Button(ui.theme, clickable, rollText)

	// Disable button appearance when no rolls left
	if ui.rollsLeft <= 0 {
		btn.Background = ui.theme.Palette.Bg
	}

	return btn
}
