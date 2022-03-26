package interaction

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func UserInstruction() {

	//game title
	gameTitle := figure.NewFigure("Monster Slayer", "", true)
	gameTitle.Print()
	fmt.Print("-----------------------------------------------------")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1. Both Player and Monster starts up with 100% health score")
	fmt.Println("2. Prayer can attack endlessly as long as health score is not zero")
	fmt.Println("3. Prayer can use special attack every third round")
	fmt.Println("4. Prayer can heal endlessly as long as health score is not zero")
	fmt.Println("5. Monster will always attack irrespective of players' attack or healing.")
	fmt.Println("6. Whoever reaches a health score of zero losses")

	StartGame()

}
