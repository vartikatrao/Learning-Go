package main

import (
	"fmt"
	"net/url"
)

const myurl = "httpls://www.wikipedia.org:8080/learn?coursename=reactjs&author=facebook"

func main() {
	fmt.Println("Welcome to urls")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of qparams: %T\n", qparams)

	fmt.Println("Query params: ", qparams["coursename"], "author", qparams["author"])

	for key, val := range qparams {
		fmt.Println(key,": ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "www.wikipedia.org",
		Path: "/learn",
		RawQuery: "coursename=reactjs&author=facebook",
	}

	fmt.Println(partsOfUrl.String())

}
