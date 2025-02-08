package yatzy

type Dice struct {
	Key         string
	Value       int
	KeyValueMap map[string]int
	Selected    bool
}

func getDiceMap(value int) map[string]int {
	diceMap := make(map[string]int)
	switch value {
	case 1:
		diceMap["\n . \n"] = 1
	case 2:
		diceMap["  .\n\n.  "] = 2
	case 3:
		diceMap[".  \n . \n  ."] = 3
	case 4:
		diceMap[". .\n\n. ."] = 4
	case 5:
		diceMap[". .\n . \n. ."] = 5
	case 6:
		diceMap[". .\n. .\n. ."] = 6
	}

	return diceMap
}
