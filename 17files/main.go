package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Go!")
	content := "this is a file content"
	file, err := os.Create("myFile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes to file\n", length)
	defer file.Close()
	readFile("myFile.txt")

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("File content is: ", string(databyte))

}
