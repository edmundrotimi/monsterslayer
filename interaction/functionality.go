// user select option
package interaction

import (
	"fmt"

	"github.com/edmundrotimi/monsterslayer/action"
	"github.com/edmundrotimi/monsterslayer/state"
)

func GameFunction(userInput int) {

	if userInput == 1 || userInput == 2 || userInput == 3 {
		currentStatus := status[int(userInput)]

		//perform attack functionality
		if currentStatus == "Attack" {
			action.StatusAttack()
		}

		//increase round by 1
		state.GameMode["round"]++
	} else {
		fmt.Println("Invalid selection; enter new option")
	}
}
