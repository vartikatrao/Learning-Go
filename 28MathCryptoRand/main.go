package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to 28MathCryptoRand")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 3.3

	// fmt.Println("Addition: ", mynumberOne+mynumberTwo)

	// random number
	// fmt.Println(rand.Intn(5) + 1) // 1-5

	// random number from crypto/rand

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)

}
