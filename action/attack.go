package action

import (
	"fmt"

	"github.com/edmundrotimi/monsterslayer/state"
)

func StatusAttack() {
	//Player attack monster
	playerAttackDamage := RandGen(state.GameMode["minRandNum"], state.GameMode["maxRandNum"])
	state.HealthStatus["monster"] = state.HealthStatus["monster"] - playerAttackDamage

	// monster attack player
	monsterAttackDamage := RandGen(state.GameMode["monsterMinRandNum"], state.GameMode["monsterMaxRandNum"])
	state.HealthStatus["player"] = state.HealthStatus["player"] - monsterAttackDamage

	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Player attack monster with %v damage level\n", playerAttackDamage)
	fmt.Printf("Monster attack player with %v damage level\n", monsterAttackDamage)
	fmt.Println("*****************************************************")
	fmt.Printf("Player Health: %v\n", state.HealthStatus["player"])
	fmt.Printf("Monster Health: %v\n", state.HealthStatus["monster"])
	fmt.Println("*****************************************************")

}
