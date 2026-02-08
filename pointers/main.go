package main

import "fmt"

func main()  {
	fmt.Println("POINTERS")

	// var ptr *int
	// fmt.Println(ptr)

	num := 60
	ptr := &num
	fmt.Println("Mem address",ptr)
	fmt.Println("Actual value",&ptr)

	*ptr = *ptr +10
	fmt.Println("New Val is", num)
	
}