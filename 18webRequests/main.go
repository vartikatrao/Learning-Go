package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.wikipedia.org/"

func main() {
	// Start the server
	fmt.Println("web requests in golang")

	respone, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", respone)

	defer respone.Body.Close()

	databytes, err:= io.ReadAll(respone.Body)

	if err!=nil{
		panic (err)
	}
	content:= string(databytes)

	fmt.Println(content)


}
