package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to API requests in golang")
	// performGetRequest()
	perfromPostJsonRequest()
	performPostFormRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Response is of type: %T\n", response)
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)
	var responeString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responeString.Write(content)
	fmt.Println("Number of bytes written to string: ", byteCount)
	fmt.Println("responseString: ", responeString.String())

}

func perfromPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
	{
		"course": "golang",
		"price": 0,
		"platform": "golang.org"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Response content: ", string(content))


}

func performPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	//form data

	data := url.Values{}
	data.Add("firstname", "Vartika")
	data.Add("lastname", "Gupta")

	response, err:= http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Response content: ", string(content))
}