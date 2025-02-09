package main

import (
	"image"
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

type (
	C = layout.Context
	D = layout.Dimensions
)

// UI holds all our app state
type UI struct {
	window     *app.Window
	theme      *material.Theme
	dices      []*widget.Clickable
	values     []string
	rollButton *widget.Clickable
}

// NewUI creates and initializes a new UI
func NewUI() *UI {
	// Initialize dices and values
	dices := make([]*widget.Clickable, 5)
	values := make([]string, 5)
	for i := range dices {
		dices[i] = new(widget.Clickable)
		v := yatzy.GetRandomValue()
		values[i] = yatzy.GetKey(v)
	}

	return &UI{
		theme:      material.NewTheme(),
		dices:      dices,
		values:     values,
		rollButton: &widget.Clickable{},
	}
}

// createDiceButton creates a styled dice button
func (ui *UI) createDiceButton(value string, clickable *widget.Clickable) material.ButtonStyle {
	btn := material.Button(ui.theme, clickable, value)
	btn.Inset.Top = unit.Dp(0)
	btn.Inset.Bottom = unit.Dp(10)
	btn.TextSize = 15.0
	btn.Font.Weight = font.Black
	// btn.Background = ui.theme.Palette.ContrastBg
	btn.CornerRadius = unit.Dp(8)
	return btn
}

// createRollButton creates a styled button
func (ui *UI) createRollButton(clickable *widget.Clickable) material.ButtonStyle {
	btn := material.Button(ui.theme, clickable, "ROLL")
	return btn
}

// layoutRollButton creates the layout for button
func (ui *UI) layoutRollButton(gtx C) D {
	// Create buttons
	btns := make([]material.ButtonStyle, 0, 5)
	for i := range ui.dices {
		btn := ui.createDiceButton(ui.values[i], ui.dices[i])
		btns = append(btns, btn)
	}

	// Create flex children with spacers
	children := make([]layout.FlexChild, len(btns)*2-1)
	for i := range btns {
		btn := btns[i]
		children[i*2] = layout.Rigid(func(gtx C) D {
			gtx.Constraints.Min = image.Point{X: 155, Y: 155}
			gtx.Constraints.Max = image.Point{X: 155, Y: 155}
			return btn.Layout(gtx)
		})

		if i < len(btns)-1 {
			children[i*2+1] = layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout)
		}
	}

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx, children...)
}

// layoutDiceButtons creates the layout for all dice buttons
func (ui *UI) layoutDiceButtons(gtx C) D {
	btn := ui.createRollButton(ui.rollButton)
	return btn.Layout(gtx)
}

// layout handles the main layout of the app
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

// handleEvents handles all window events
func (ui *UI) handleEvents() error {
	var ops op.Ops

	for {
		evt := ui.window.Event()

		switch e := evt.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if ui.rollButton.Clicked(gtx) {
				for i := range ui.dices {
					v := yatzy.GetRandomValue()
					ui.values[i] = yatzy.GetKey(v)
				}
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
