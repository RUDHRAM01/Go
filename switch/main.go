package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Switch statement

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		println("You rolled a 1")
	case 2:
		println("You rolled a 2")
	case 3:
		println("You rolled a 3")
	case 4:
		println("You rolled a 4")
	case 5:
		println("You rolled a 5")
	case 6:
		println("You rolled a 6")
	default:
		println("Invalid dice number")
	}
}
