// // simple http server

package main

import (
	"fmt"
	"net/http"
	"time"
)

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Request recieved")
// 	fmt.Println(r.Method)
// 	fmt.Println(r.URL.Path)

// 	// w.Write([]byte("Hwllo world"))
// 	// fmt.Fprintln(w, "Hwllo World")

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{message: "Hello World"}`))
// }

// func main() {

// 	http.HandleFunc("/hello", helloHandler)

// 	fmt.Println("Server running on port 8080.")
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Println("Error occured: ", err)
// 	}

// }

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed (Use GET method)", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Fprintln(w, "Welcome to my API")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Fprintln(w, "This is about route.")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Fprintln(w, time.Now())
}

func main() {
	port := ":8080"
	mux := http.NewServeMux()

	// Welcome route
	mux.HandleFunc("/", welcomeHandler)

	// About route
	mux.HandleFunc("/about", aboutHandler)

	// Time route
	mux.HandleFunc("/time", timeHandler)

	fmt.Println("Server started at port", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}

}
