package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	ch := make(chan string, 1024)
	var wg sync.WaitGroup
	workerCount := 200

	for i := 0; i < workerCount; i++ {
		go worker(ch, &wg)
	}

	urls := [3]string{"https://www.google.com", "https://www.time.ir", "https://www.toplearn.com"}
	for i := 0; i < 100; i++ {
		for _, url := range urls {
			wg.Add(1)
			ch <- url
		}
	}
	wg.Wait()
	fmt.Println("The End")
}

func worker(input <-chan string, wg *sync.WaitGroup) {
	for {
		url := <-input
		resp, err := http.Get(url)
		if err != nil {
			log.Print(err.Error())
		} else {
			log.Printf("Get %s, StatusCode %d", url, resp.StatusCode)
		}

		wg.Done()
	}
}
