package main

import "fmt"

func main() {
	var ptr *int

	fmt.Println("Value of the pointer: ", ptr)

	myNumber := 23
	ptr = &myNumber

	fmt.Println("Value of the pointer: ", ptr)
	fmt.Println("Value in the pointer: ", *ptr)


	*ptr= *ptr + 2
	fmt.Println("Value in the pointer: ", *ptr )
}
