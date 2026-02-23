package main

import "fmt"

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

// func cook(ch chan string) {
// 	ch <- "momo"
// }

// func waiter(ch chan string) {
// 	food := <-ch
// 	fmt.Println("Serving:", food)
// }

// func main() {
// 	ch := make(chan string)

// 	go cook(ch)
// 	go waiter(ch)

// 	time.Sleep(time.Second)
// }

func main() {
	ch := make(chan int, 3) // buffered channel

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 // deadlock

	fmt.Println("Done")
}
