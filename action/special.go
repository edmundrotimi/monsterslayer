package action

import (
	"fmt"

	"github.com/edmundrotimi/monsterslayer/state"
)

func StatusSpecialAttack() {

	//Player attack monster
	playerAttackDamage := RandGen(state.GameMode["specialMinRandNum"], state.GameMode["specialMaxRandNum"])
	monsterAttackDamage := RandGen(state.GameMode["monsterMinRandNum"], state.GameMode["monsterMaxRandNum"])
	state.HealthStatus["monster"] = state.HealthStatus["monster"] - playerAttackDamage

	// monster attack player
	state.HealthStatus["player"] = state.HealthStatus["player"] - monsterAttackDamage

	fmt.Println("-----------------------------------------------------")
	fmt.Printf("Player attack monster with %v damage level\n", playerAttackDamage)
	fmt.Printf("Monster attack player with %v damage level\n", monsterAttackDamage)
	fmt.Println("*****************************************************")
	fmt.Printf("Player Health: %v\n", state.HealthStatus["player"])
	fmt.Printf("Monster Health: %v\n", state.HealthStatus["monster"])
	fmt.Println("*****************************************************")

}
