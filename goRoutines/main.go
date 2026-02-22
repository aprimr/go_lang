package main

import (
	"fmt"
	"sync"
)

func display(num int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Printing num", num)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go display(i, &wg)
	}

	wg.Wait()
}
