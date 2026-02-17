package main

import (
	"fmt"
	"io"
	"net/http"
)

var url string = "https://jsonplaceholder.typicode.com/todos/3"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic("Error fetching url")
	}
	fmt.Printf("The type of response is %T. \n", res)
	defer res.Body.Close()
	dataBytes, e := io.ReadAll(res.Body)
	if e != nil {
		panic("Error occured")
	}
	content := string(dataBytes)
	fmt.Println("The data is:", content)
}