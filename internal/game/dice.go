package game

import (
	"gioui.org/widget"
	"golang.org/x/exp/rand"
)

type Dice struct {
	Key      string
	Value    int
	Widget   *widget.Clickable
	Selected bool
}

func GetRandomValue() int {
	return rand.Intn(7-1) + 1
}

func GetKey(value int) string {
	str := ""
	switch value {
	case 1:
		str = "  .  "
	case 2:
		str = "    .\n\n.    "
	case 3:
		str = ".    \n  .  \n    ."
	case 4:
		str = ".    .\n\n.    ."
	case 5:
		str = ".    .\n . \n.    ."
	case 6:
		str = ".    .\n.    .\n.    ."
	}

	return str
}
