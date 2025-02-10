package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"github.com/WiviWonderWoman/yatzy/internal/ui"
)

func main() {
	go func() {
		ui := ui.NewUI()
		ui.Window = new(app.Window)
		ui.Window.Option(app.Size(unit.Dp(600), unit.Dp(600)))
		ui.Window.Option(app.Title("Y A T Z Y"))

		if err := ui.HandleEvents(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
