package action

import (
	"fmt"

	"github.com/edmundrotimi/monsterslayer/state"
)

func StatusHeal() {

	//Player healing
	playerHealBalm := RandGen(state.GameMode["minRandNum"], state.GameMode["maxRandNum"])
	// ensure health do not exceed 100
	healthLevel := state.HealthStatus["player"] + playerHealBalm

	if healthLevel > 100 {
		state.HealthStatus["player"] = 100
	} else {
		state.HealthStatus["player"] = state.HealthStatus["player"] + playerHealBalm
	}

	// monster attack player
	monsterAttackDamage := RandGen(state.GameMode["monsterMinRandNum"], state.GameMode["monsterMaxRandNum"])
	state.HealthStatus["player"] = state.HealthStatus["player"] - monsterAttackDamage

	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Player used health toolbox with %v heal level\n", playerHealBalm)
	fmt.Printf("Monster attack player with %v damage level\n", monsterAttackDamage)
	fmt.Println("*****************************************************")
	fmt.Printf("Player Health: %v\n", state.HealthStatus["player"])
	fmt.Printf("Monster Health: %v\n", state.HealthStatus["monster"])
	fmt.Println("*****************************************************")

}
