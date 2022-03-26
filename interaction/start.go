package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var startPrompt bool = true

type options []string

var startStatus map[string]options = map[string]options{
	"input": []string{
		"Shuting down .....",
		"Starting Game ....",
		"Invalid input; enter new input",
	},
}

func StartGame() {
	for startPrompt {

		fmt.Println("-----------------------------------------------------")
		fmt.Print("Press 1 to start game or 0 to stop : ")

		reader := bufio.NewReader(os.Stdin)
		userInput, err := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if err != nil {
			fmt.Println(startStatus["error"])
		}

		if userInput == "0" || userInput == "1" {
			input, err := strconv.ParseInt(userInput, 0, 64)

			if err != nil {
				fmt.Println(startStatus["input"][2])
			} else {
				fmt.Println(startStatus["input"][input])
				startPrompt = false
			}

		} else {
			fmt.Println(startStatus["input"][2])
		}
	}
}
