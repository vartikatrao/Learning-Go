package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format(time.RFC822))

	createdDate := time.Date(202, time.December, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format(time.RFC822))

}
