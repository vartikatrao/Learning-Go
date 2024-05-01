package main

import "fmt"

func main() {
	fmt.Println("Maps in go")
	languages := make (map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Map of languages: ", languages)
	fmt.Println("Js:", languages["JS"])

	delete (languages, "RB")
	fmt.Println("Map of languages: ", languages)

	// loops are intersting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
