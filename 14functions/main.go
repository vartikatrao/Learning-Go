package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is ", result)

	proResult, _ := proAdder(3, 5, 6, 7, 8, 9)
	fmt.Println("Pro Result is ", proResult)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total:=0
	for _, val := range values{
		total+=val
	}
	return total, "Hi, there"
}
func greeter()  {
	fmt.Println("Namastey from India")
}
