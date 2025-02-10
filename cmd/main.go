package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/WiviWonderWoman/yatzy"
)

// Type aliases for layout package
type (
	C = layout.Context
	D = layout.Dimensions
)

// First, let's extend the UI struct to include scoring
type UI struct {
	window     *app.Window
	theme      *material.Theme
	dices      []yatzy.Dice
	rollButton *widget.Clickable
	rollsLeft  int
	upperBoxes []yatzy.UpperScoreBox // Add upper section scoring
	lowerBoxes []yatzy.LowerScoreBox // Add lower section scoring
	totalScore int                   // Track total score
}

// Update NewUI to initialize scoring
func NewUI() *UI {
	// Initialize dice (existing code)
	dices := make([]yatzy.Dice, 5)
	for i := range dices {
		dices[i] = yatzy.Dice{
			Widget:   new(widget.Clickable),
			Selected: false,
			Value:    yatzy.GetRandomValue(),
		}
		dices[i].Key = yatzy.GetKey(dices[i].Value)
	}

	// Initialize scoring boxes with clickable widgets
	upperBoxes := make([]yatzy.UpperScoreBox, len(yatzy.UpperBoxes))
	copy(upperBoxes, yatzy.UpperBoxes)
	for i := range upperBoxes {
		upperBoxes[i].Widget = new(widget.Clickable)
	}

	lowerBoxes := make([]yatzy.LowerScoreBox, len(yatzy.LowerBoxes))
	copy(lowerBoxes, yatzy.LowerBoxes)
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

// Add scoring box button creation
func (ui *UI) createScoreButton(
	key string,
	value string,
	calculated bool,
	clickable *widget.Clickable,
) material.ButtonStyle {
	btn := material.Button(ui.theme, clickable, fmt.Sprintf("%s: %s", key, value))

	if calculated {
		// Already scored - disable the button
		btn.Background = color.NRGBA{R: 200, G: 200, B: 200, A: 255}
		btn.Color = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	} else {
		// Available for scoring
		btn.Background = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		btn.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	}

	btn.TextSize = unit.Sp(14)
	btn.CornerRadius = unit.Dp(4)
	return btn
}

// Add layout for score section
func (ui *UI) layoutScoreSection(gtx C) D {
	upperSum := yatzy.CalculateUpperSum(ui.upperBoxes)
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
				layout.Rigid(func(gtx C) D {
					bonusText := fmt.Sprintf("Bonus: %d (Need %d more for bonus)",
						bonusScore,
						max(0, 63-upperSum))
					return material.Body2(ui.theme, bonusText).Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Inset{
						Top:    unit.Dp(10),
						Bottom: unit.Dp(5),
					}.Layout(gtx, material.Body2(ui.theme, fmt.Sprintf("Sum: %d", upperSum)).Layout)
				}),
			)
		}),
		// // Add some space between columns
		// layout.Rigid(layout.Spacer{Width: unit.Dp(20)}.Layout),
		// Lower section (right column)
		layout.Rigid(func(gtx C) D {
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(material.H6(ui.theme, "Lower Section").Layout),
				layout.Rigid(ui.layoutLowerBoxes),
				layout.Rigid(func(gtx C) D {
					lowerSum := yatzy.CalculateLowerSum(ui.lowerBoxes)
					return layout.Inset{
						Top: unit.Dp(10),
					}.Layout(gtx, material.Body2(ui.theme, fmt.Sprintf("Sum: %d", lowerSum)).Layout)
				}),
			)
		}),
	)
}

// Add layout for upper score boxes
func (ui *UI) layoutUpperBoxes(gtx C) D {
	children := make([]layout.FlexChild, len(ui.upperBoxes))

	for i := range ui.upperBoxes {
		box := &ui.upperBoxes[i]
		btn := ui.createScoreButton(box.Key, box.Value, box.Calculate, box.Widget)

		children[i] = layout.Rigid(func(gtx C) D {
			return btn.Layout(gtx)
		})
	}

	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}.Layout(gtx, children...)
}

// Add layout for lower score boxes
func (ui *UI) layoutLowerBoxes(gtx C) D {
	children := make([]layout.FlexChild, len(ui.lowerBoxes))

	for i := range ui.lowerBoxes {
		box := &ui.lowerBoxes[i]
		btn := ui.createScoreButton(box.Key, box.Value, box.Calculate, box.Widget)

		children[i] = layout.Rigid(func(gtx C) D {
			return btn.Layout(gtx)
		})
	}

	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}.Layout(gtx, children...)
}

// createDiceButton creates a styled button for a die
// The appearance changes based on whether the die is selected (locked)
func (ui *UI) createDiceButton(dice *yatzy.Dice) material.ButtonStyle {
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

// rollDice rolls all unlocked dice
func (ui *UI) rollDice() {
	for i := range ui.dices {
		if !ui.dices[i].Selected {
			ui.dices[i].Value = yatzy.GetRandomValue()
			ui.dices[i].Key = yatzy.GetKey(ui.dices[i].Value)
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
	upperSum := yatzy.CalculateUpperSum(ui.upperBoxes)
	lowerSum := yatzy.CalculateLowerSum(ui.lowerBoxes)

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
func (ui *UI) handleEvents() error {
	var ops op.Ops

	for {
		evt := ui.window.Event()

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

func main() {
	go func() {
		ui := NewUI()
		ui.window = new(app.Window)
		ui.window.Option(app.Size(unit.Dp(600), unit.Dp(600)))
		ui.window.Option(app.Title("Y A T Z Y"))

		if err := ui.handleEvents(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
