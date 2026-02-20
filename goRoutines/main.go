package main

import (
	"fmt"
	"time"
)

func display(num int) {
	fmt.Println("Printing num", num)
}

func main() {
	for i := 0; i <= 10; i++ {
		go display(i)
	}

	time.Sleep(time.Second * 2)
}
