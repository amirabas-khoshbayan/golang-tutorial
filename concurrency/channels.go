package main

import (
	"fmt"
	"time"
)

// Goroutines are basically independent, but they might need to communicate (practical examples will be involved)
// Channels are a way for goroutines communications (transferring data and signals between them)
func Channels() {
	//How to make them
	myChannel := make(chan string)

	//How to send to them
	myChannel <- "Sending to a channel"

	//How to receive from them
	myVal := <-myChannel

	fmt.Println(myVal)
}

// Receiving and sending to channels is blocking.
// A blocking operation is an operation that we wait to complete before going to further.

func main() {
	myChan := make(chan string)

	go routine(myChan)
	<-myChan
	fmt.Println("Outside goroutine")

	time.Sleep(time.Second * 2)

}
func routine(myChan chan string) {
	myChan <- "" // Blocking
	fmt.Println("Inside goroutine")
}
