package main

import "fmt"

// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	nums := []int{1, 2, 3}
// 	names := []string{"go lang", "ts", "c", "c++"}
// 	bools := []bool{true, false, true, true}
// 	printSlice(nums)
// 	printSlice(names)
// 	printSlice(bools)
// }

type stack [T int | string] struct {
	elements []T
}

func main() {
	myStack := stack[string]{
		elements: []string{"1", "2", "3"},
	}

	fmt.Println(myStack)
}
