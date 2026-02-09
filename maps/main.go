package main

import "fmt"

func main()  {
	marks := make(map[string]int)
	marks["MP"] = 25
	marks["IT"] = 40
	marks["DL"] = 20

	fmt.Println("The marks map is ", marks)
	delete(marks, "IT")
	fmt.Println("The marks map is ", marks)

	for key, value := range marks{
		fmt.Printf("The marks for %v is %d \n", key, value)
	}
}