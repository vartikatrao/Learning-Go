package main

import "fmt"
func main() {
	var username string = "Vartika"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)
	q := "Vartika"
	const p = "Vartika"
	fmt.Println("Value of q: ", q)
	fmt.Println("Value of p: ", p)
}
