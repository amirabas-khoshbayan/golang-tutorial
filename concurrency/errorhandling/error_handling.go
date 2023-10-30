package main

import (
	"fmt"
	"net/http"
)

type Res struct {
	err     error
	httpRes *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls []string) <-chan Res {
		responses := make(chan Res)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					responses <- Res{err: err, httpRes: resp}
					continue
				}
				select {
				case <-done:
					return
				case responses <- Res{err: nil, httpRes: resp}:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)
	errCount := 0

	urls := []string{"https://www.google.com", "https://badhost", "cc", "ad"}
	for response := range checkStatus(done, urls) {

		if response.err != nil {
			fmt.Printf("Error : %v\n", response.err)
			errCount++

			if errCount > 2 {
				fmt.Println("error > 2")
				break
			}
			continue
		}

		fmt.Printf("Response: %v\n", response.httpRes.Status)
	}

}
