package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Golang")

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day:= range days {
	// 	fmt.Printf("Day %d of week = %s\n", index, day)
	// }

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 5 {
			goto v
		}
		fmt.Println(rougueValue)
		rougueValue++
	}
v:
	fmt.Printf("Value of v = %d\n", 10)
}
