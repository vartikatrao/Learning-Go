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
	vartika.GetStatus()
	vartika.NewEmail()
	fmt.Println("Updated email is ", vartika.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail()  {
	u.Email= "test@go.dev"
	fmt.Println("Email is ", u.Email)
}