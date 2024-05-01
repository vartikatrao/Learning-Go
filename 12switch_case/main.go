package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// switch case in golang
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got one")
	case 2:
		fmt.Println("You got two")
	case 3:
		fmt.Println("You got three")
	case 4:
		fmt.Println("You got four")
	case 5:
		fmt.Println("You got five")
	case 6:
		fmt.Println("You got six")
	default:
		fmt.Println("Invalid dice number")
	}

}
