package main

import "fmt"

func main() {
	fmt.Println("Structs in go")
	// no inheritance in go
	vartika := User {
		Name: "Vartika",
		Email: "vartikatrao@gmail.com", 
		Status: true,
		Age: 21, 
	}
	fmt.Println("User: ", vartika)
	fmt.Printf("User: %+v\n", vartika)
	fmt.Printf("Name is %v\n", vartika.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
