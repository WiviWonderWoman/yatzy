package ui

import (
	"fmt"
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/WiviWonderWoman/yatzy/internal/game"
)

// Type aliases for layout package
type (
	C = layout.Context
	D = layout.Dimensions
)

// / I main.go, uppdatera layoutScoreSection:
func (ui *UI) layoutScoreSection(gtx C) D {
	upperSum := game.CalculateUpperSum(ui.upperBoxes)
	lowerSum := game.CalculateLowerSum(ui.lowerBoxes)
	bonusScore := 0
	if upperSum >= 63 {
		bonusScore = 50
	}

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Start,
		Spacing:   layout.SpaceBetween,
	}.Layout(gtx,

		// Upper section (left column)
		layout.Rigid(func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(material.H6(ui.theme, "Upper Section").Layout),
				layout.Rigid(ui.layoutUpperBoxes),

				// Bonus
				layout.Rigid(func(gtx C) D {
					bonusText := fmt.Sprintf("Bonus: %d (Need %d more for bonus)",
						bonusScore,
						max(0, 63-upperSum))
					return material.Body2(ui.theme, bonusText).Layout(gtx)
				}),
			)
		}),

		// Lower section (right column)
		layout.Rigid(func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(material.H6(ui.theme, "Lower Section").Layout),
				layout.Rigid(ui.layoutLowerBoxes),

				// Total score (including upper, lower and bonus)
				layout.Rigid(func(gtx C) D {
					total := upperSum + lowerSum + bonusScore
					return layout.Inset{
						Top: unit.Dp(10),
					}.Layout(gtx, material.H6(ui.theme, fmt.Sprintf("Total Score: %d", total)).Layout)
				}),
			)
		}),
	)
}

// Add layout for upper score boxes
func (ui *UI) layoutUpperBoxes(gtx C) D {
	children := make([]layout.FlexChild, len(ui.upperBoxes)*2-1) // För knappar och mellanrum

	for i := range ui.upperBoxes {
		box := &ui.upperBoxes[i]
		btn := ui.createScoreButton(box.Key, box.Value, box.Calculate, box.Widget)

		// Lägg till knappen med fast storlek
		children[i*2] = layout.Rigid(func(gtx C) D {
			// Sätt en fast minsta bredd på knapparna
			gtx.Constraints.Min.X = gtx.Dp(200)
			return btn.Layout(gtx)
		})

		// Lägg till mellanrum mellan knapparna (utom efter sista)
		if i < len(ui.upperBoxes)-1 {
			children[i*2+1] = layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout)
		}
	}

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Start,
	}.Layout(gtx, children...)
}

// Add layout for lower score boxes
func (ui *UI) layoutLowerBoxes(gtx C) D {
	children := make([]layout.FlexChild, len(ui.lowerBoxes)*2-1) // För knappar och mellanrum

	for i := range ui.lowerBoxes {
		box := &ui.lowerBoxes[i]
		btn := ui.createScoreButton(box.Key, box.Value, box.Calculate, box.Widget)

		// Lägg till knappen med fast storlek
		children[i*2] = layout.Rigid(func(gtx C) D {
			// Sätt en fast minsta bredd på knapparna
			gtx.Constraints.Min.X = gtx.Dp(200)
			return btn.Layout(gtx)
		})

		// Lägg till mellanrum mellan knapparna (utom efter sista)
		if i < len(ui.lowerBoxes)-1 {
			children[i*2+1] = layout.Rigid(layout.Spacer{Height: unit.Dp(4)}.Layout)
		}
	}

	return layout.Flex{
		Axis:      layout.Vertical,
		Alignment: layout.Start,
	}.Layout(gtx, children...)
}

// Uppdatera layoutRollButton för att visa antal kast kvar
func (ui *UI) layoutRollButton(gtx C) D {
	rollText := fmt.Sprintf("ROLL (%d)", ui.rollsLeft)
	btn := material.Button(ui.theme, ui.rollButton, rollText)
	if ui.rollsLeft <= 0 {
		btn.Background = ui.theme.Palette.Bg
	}
	return btn.Layout(gtx)
}

// Uppdaterad layoutDiceButtons för att använda Dice struct
func (ui *UI) layoutDiceButtons(gtx C) D {
	// Skapa en slice för alla knappar och mellanrum
	children := make([]layout.FlexChild, len(ui.dices)*2-1)

	for i := range ui.dices {
		// Skapa en knapp för varje tärning
		btn := ui.createDiceButton(&ui.dices[i])

		// Lägg till knappen i flexbox
		children[i*2] = layout.Rigid(func(gtx C) D {
			gtx.Constraints.Min = image.Point{X: 155, Y: 155}
			gtx.Constraints.Max = image.Point{X: 155, Y: 155}
			return btn.Layout(gtx)
		})

		// Lägg till mellanrum mellan knapparna (utom efter sista)
		if i < len(ui.dices)-1 {
			children[i*2+1] = layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout)
		}
	}

	// Returnera flex-layouten
	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx, children...)
}

// layout handles the main layout of the app
func (ui *UI) layout(gtx C) D {
	margins := layout.Inset{
		Top:    unit.Dp(20),
		Bottom: unit.Dp(20),
		Left:   unit.Dp(20),
		Right:  unit.Dp(20),
	}

	return margins.Layout(gtx, func(gtx C) D {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			// Score sections at the top
			layout.Rigid(ui.layoutScoreSection),

			// Spacer to push dice to bottom
			layout.Flexed(1, layout.Spacer{}.Layout),

			// Dice and roll button at the bottom
			layout.Rigid(func(gtx C) D {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					// Dice buttons
					layout.Rigid(ui.layoutDiceButtons),
					// Small space
					layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
					// Roll button below dice
					layout.Rigid(func(gtx C) D {
						// Center the roll button
						return layout.Center.Layout(gtx, ui.layoutRollButton)
					}),
				)
			}),
		)
	})
}
