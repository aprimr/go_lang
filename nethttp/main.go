// simple http server

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request recieved")
	w.Write([]byte("Hwllo world"))
}

func main() {

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running on port 8080.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}

}
