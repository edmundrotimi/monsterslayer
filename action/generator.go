package action

import (
	"math/rand"
	"time"
)

func RandGen(min int, max int) int {
	//generate random attack
	randSource := rand.NewSource(time.Now().UnixNano())
	randN := rand.New(randSource)
	randGenerate := randN.Intn(max-min) + min

	return randGenerate
}
