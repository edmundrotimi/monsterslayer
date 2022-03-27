// user select option
package interaction

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"

	"github.com/edmundrotimi/monsterslayer/action"
	"github.com/edmundrotimi/monsterslayer/state"
)

func GameFunction(userInput int) {

	//checkfor winner
	playerHealthCheck := state.HealthStatus["player"] > 0
	monsterHealthCheck := state.HealthStatus["monster"] > 0

	if playerHealthCheck && monsterHealthCheck {
		if userInput == 1 || userInput == 2 || (userInput == 3 && state.GameMode["round"]%3 == 0) {
			currentStatus := status[int(userInput)]

			//perform attack functionality
			if currentStatus == "Attack" {
				action.StatusAttack()
			} else if currentStatus == "Heal" {
				action.StatusHeal()
			} else {
				action.StatusSpecialAttack()
			}

			//increase round by 1
			state.GameMode["round"]++
		} else {
			fmt.Println("--------------------------------------")
			fmt.Println("Invalid selection; enter new option")
		}
	} else {
		//end game active
		state.GameMode["active"] = 0
		winner := ""

		func() {
			gameOver := figure.NewFigure("Game Over", "", true)
			gameOver.Print()
			fmt.Println("--------------------------------------")

			switch true {
			case !(state.HealthStatus["player"] > 0):
				winner = "Monster"
			case !(state.HealthStatus["Monster"] > 0):
				winner = "Player"
			}

			fmt.Printf("%v Wins!!!", winner)
		}()
	}
}
