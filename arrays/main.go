package main

import "fmt"

func main()  {
	fmt.Println("Arrays")
	
	var fruits [5]string
	var bools  [3]bool
	var nums  [3]int

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "peach"

	fmt.Println(fruits)
	fmt.Println(bools)
	fmt.Println(nums)

	fmt.Println("Arr length is ",len(fruits))
}