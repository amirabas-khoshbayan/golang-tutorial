package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Start")
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printMessage(i, &wg)
	}
	wg.Wait()

	fmt.Println("End")
}
func printMessage(index int, wg *sync.WaitGroup) {
	fmt.Println(index)
	wg.Done() // or wg.Add(-1)

}
