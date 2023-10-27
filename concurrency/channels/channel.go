package main

import (
	"fmt"
	"time"
)

func main() {
	// Buffer Channel
	ch := make(chan int, 10)

	//UnBuffer Channel
	//ch := make(chan int)

	go sendMessage(ch)
	for message := range ch {
		fmt.Println(message)
	}

	//time.Sleep(time.Second * 2)

}

func sendMessage(ch chan int) {
	for i := 0; i < cap(ch); i++ {
		time.Sleep(2 * time.Second)
		ch <- i
	}

	close(ch)
	fmt.Println("finish")

}
