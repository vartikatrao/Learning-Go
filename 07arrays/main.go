package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Onion"}
	fmt.Println("Veg list is: ", vegList)
}
