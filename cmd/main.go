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

// UI holds the state for the application
type UI struct {
	window     *app.Window
	theme      *material.Theme
	dices      []yatzy.Dice // Slice of dice for the game
	rollButton *widget.Clickable
	rollsLeft  int // Number of rolls remaining in current turn
}

// NewUI creates and initializes a new UI instance
func NewUI() *UI {
	// Initialize five dice
	dices := make([]yatzy.Dice, 5)
	for i := range dices {
		dices[i] = yatzy.Dice{
			Widget:   new(widget.Clickable),
			Selected: false,
			Value:    yatzy.GetRandomValue(),
		}
		dices[i].Key = yatzy.GetKey(dices[i].Value)
	}

	return &UI{
		theme:      material.NewTheme(),
		dices:      dices,
		rollButton: new(widget.Clickable),
		rollsLeft:  3, // Start with 3 rolls per turn
	}
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

// layoutDiceButtons creates the layout for all dice buttons
func (ui *UI) layoutDiceButtons(gtx C) D {
	// Create a slice for buttons and spacers
	children := make([]layout.FlexChild, len(ui.dices)*2-1)

	for i := range ui.dices {
		btn := ui.createDiceButton(&ui.dices[i])

		// Add button to flex layout
		children[i*2] = layout.Rigid(func(gtx C) D {
			gtx.Constraints.Min = image.Point{X: 155, Y: 155}
			gtx.Constraints.Max = image.Point{X: 155, Y: 155}
			return btn.Layout(gtx)
		})

		// Add spacer between buttons (except after the last one)
		if i < len(ui.dices)-1 {
			children[i*2+1] = layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout)
		}
	}

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx, children...)
}

// layoutRollButton creates the layout for the roll button
func (ui *UI) layoutRollButton(gtx C) D {
	btn := ui.createRollButton(ui.rollButton)
	return btn.Layout(gtx)
}

// layout handles the main layout of the application
func (ui *UI) layout(gtx C) D {
	margins := layout.Inset{
		Top:    unit.Dp(25),
		Bottom: unit.Dp(25),
		Right:  unit.Dp(35),
		Left:   unit.Dp(35),
	}

	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceEnd,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx C) D {
				return margins.Layout(gtx, ui.layoutRollButton)
			},
		),
		layout.Rigid(
			func(gtx C) D {
				return margins.Layout(gtx, ui.layoutDiceButtons)
			},
		),
	)
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

// handleEvents processes all window events
func (ui *UI) handleEvents() error {
	var ops op.Ops

	for {
		evt := ui.window.Event()

		switch e := evt.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Handle dice selection
			for i := range ui.dices {
				if ui.dices[i].Widget.Clicked(gtx) {
					ui.dices[i].Selected = !ui.dices[i].Selected
				}
			}

			// Handle roll button clicks
			if ui.rollButton.Clicked(gtx) && ui.rollsLeft > 0 {
				ui.rollDice()
				ui.rollsLeft--
			}

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
		ui.window.Option(app.Title("Y A T Z Y"))

		if err := ui.handleEvents(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
