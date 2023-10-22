package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	//resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
	resp, err := client.Get("https://dummyjson.com/products")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	var myMap = map[string]any{}

	json.NewDecoder(resp.Body).Decode(&myMap)
	fmt.Println(myMap)

	//io.Copy(os.Stdout, resp.Body)
	//totalBySlice, err := io.ReadAll(resp.Body)
	//json.Unmarshal(totalBySlice, &myMap)

}
