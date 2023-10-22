package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	totalBytes := make([]byte, 0)
	bs := make([]byte, 4)
	for err != io.EOF {
		_, err = resp.Body.Read(bs)
		totalBytes = append(totalBytes, bs...)
	}
	fmt.Println(string(totalBytes))
}
