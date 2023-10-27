package main

import (
	"fmt"
	"sync"
)

// go run -race memory_share.go
// check Mutex for access control Goroutine to Memory

type Counter struct {
	sync.Mutex
	Value int
}

func main() {
	fmt.Println("Start")
	var wg sync.WaitGroup
	counter := Counter{}

	//for i := 1; i <= 10; i++ {
	//	wg.Add(1)
	//	go printMessage(i, &wg)
	//}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incerement(&counter, &wg)
	}

	wg.Wait()

	fmt.Println("value: ", counter.Value)
	fmt.Println("End")
}

func incerement(counter *Counter, wg *sync.WaitGroup) {
	counter.Lock()
	counter.Value++
	wg.Done()
	counter.Unlock()
}

//func printMessage(index int, wg *sync.WaitGroup) {
//	fmt.Println(index)
//	wg.Done() // or wg.Add(-1)
//
//}
