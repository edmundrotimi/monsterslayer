// user select option
package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/edmundrotimi/monsterslayer/state"
)

var status map[int]string = map[int]string{
	1: "Attack",
	2: "Heal",
	3: "SpecialAttack",
}

func SelectGameOption(startOption int) {

	if startOption == 1 {

		for state.GameMode["active"] == 1 {

			//user select option
			fmt.Println("--------------------------------------")
			fmt.Println("*************************************")
			fmt.Println("Select Game Option")
			fmt.Println("*************************************")
			fmt.Println("1. Attack Monster")
			fmt.Println("2. Heal")

			if state.GameMode["round"]%3 == 0 {

				fmt.Println("3. Special Attack")
			}

			//input prompt
			fmt.Print("Enter an option: ")
			reader := bufio.NewReader(os.Stdin)
			userInput, err := reader.ReadString('\n')

			//clean user input
			userInput = strings.TrimSpace(userInput)
			cleanInput, err := strconv.ParseInt(userInput, 0, 64)

			if err != nil {
				fmt.Println("--------------------------------------")
				fmt.Println("Invalid selection; enter new option")
			}

			//gamefunctionality
			GameFunction(int(cleanInput))
		}
	}
}
