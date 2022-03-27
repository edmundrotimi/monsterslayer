// user select option
package interaction

import (
	"fmt"

	"github.com/edmundrotimi/monsterslayer/action"
	"github.com/edmundrotimi/monsterslayer/state"
)

func GameFunction(userInput int) {

	if userInput == 1 || userInput == 2 || (userInput == 3 && state.GameMode["round"]%3 == 0) {
		currentStatus := status[int(userInput)]

		//perform attack functionality
		if currentStatus == "Attack" {
			action.StatusAttack()
		}
		if currentStatus == "Heal" {
			action.StatusHeal()
		}

		//increase round by 1
		state.GameMode["round"]++
	} else {
		fmt.Println("--------------------------------------")
		fmt.Println("Invalid selection; enter new option")
	}
}
