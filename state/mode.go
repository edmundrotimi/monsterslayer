package state

//keep track of game mode

var GameMode map[string]int = map[string]int{
	"active":            1,  // track if game has ended
	"round":             1,  // keep track of game round
	"minRandNum":        9,  // min number used for random gen
	"maxRandNum":        15, // max number used for random gen
	"specialMinRandNum": 12, // Special Atk min number used for random gen
	"specialMaxRandNum": 20, // Special Atk max number used for random gen
	"monsterMinRandNum": 15, // Monster Atk max number used for random gen
	"monsterMaxRandNum": 25, // Monster Atk max number used for random gen
}
