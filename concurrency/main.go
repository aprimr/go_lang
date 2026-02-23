package main

import (
	"fmt"
	"time"
)

// func greet(name string) {
// 	fmt.Println("Hello", name)
// }

// // func slowTask(){
// // 	time.Sleep(time.Second )
// // 	fmt.Println("Done")
// // }

// func printNumber(num int) {
// 	fmt.Println("Print: ", num)
// }

// func main() {

// 	// go routine
// 	// goroutine is non blocking function
// 	go greet("aprim")

// 	// go slowTask()
// 	// fmt.Println("Next")

// 	for i := 0; i < 10; i++ {
// 		go printNumber(i)
// 	}

// 	time.Sleep(time.Second)
// }

func sayHello() {
	fmt.Println("Hello World")
	time.Sleep(time.Second * 2)
	fmt.Println("Hello world again")
}

func sayHi() {
	fmt.Println("Hi Golang")
	time.Sleep(time.Second * 1)
	fmt.Println("Hi Golang again")
}

func main() {
	fmt.Println("Go routines")

	go sayHello()
	go sayHi()

	time.Sleep(5 * time.Second)

}
