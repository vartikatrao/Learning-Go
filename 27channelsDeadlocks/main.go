package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5  // deadlock!
	// fmt.Println(<-myCh)
	wg.Add(2)
	// receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch

		fmt.Println("Channel is open: ", isChannelOpen)
		fmt.Println("Value: ", val)

		fmt.Println(<-ch)
		wg.Done()

	}(myCh, wg)
	// send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(ch)
		// ch <- 5
		// ch <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
