package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex //pointers
func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, error := http.Get(endpoint)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("Status Code:%d for site %s \n", res.StatusCode, endpoint)

	}
}
