package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go worker1(ch1)
	go worker2(ch2)
	go worker3(ch3)

	select {
	case msg := <-ch1:
		fmt.Println("Channel1 : ", msg)
	case msg := <-ch2:
		fmt.Println("Channel2 : ", msg)
	case msg := <-ch3:
		fmt.Println("Channel3 : ", msg)
	}

	//fmt.Println("Channel1 :", <-ch1)
	//fmt.Println("Channel2 :", <-ch2)
	//fmt.Println("Channel3 :", <-ch3)

	fmt.Println("The End")
}

func worker1(ch chan int) {
	time.Sleep(8 * time.Second)
	ch <- 1
}
func worker2(ch chan int) {
	time.Sleep(7 * time.Second)
	ch <- 2
}
func worker3(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 3
}
