package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// creating channel
// 	ch := make(chan int)

// 	// send
// 	go func(){
// 		ch <- 10
// 	}()

// 	// recieve
// 	data := <-ch

// 	fmt.Println(data)
// }

func cook(ch chan string) {
	ch <- "momo"
}

func waiter(ch chan string) {
	food := <-ch
	fmt.Println("Serving:", food)
}

func main() {
	ch := make(chan string)

	go cook(ch)
	go waiter(ch)

	time.Sleep(time.Second)
}
