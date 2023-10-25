package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Start")

	message("Hello AmirAbbas")
	wg.Wait()

	fmt.Println("End")
}
func message(text string) {
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go printMessage(text, i)
	}
}
func printMessage(text string, index int) {
	if index == 2 {
		time.Sleep(2 * time.Second)
	}
	fmt.Println(text, index)
	wg.Done() // or wg.Add(-1)

}
