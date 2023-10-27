package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendMessage(ch)

	time.Sleep(time.Second * 2)
	message := <-ch
	fmt.Println(message)

}

func sendMessage(ch chan string) {
	fmt.Println("finish")
	ch <- "WelCome Amirabbas"

}
