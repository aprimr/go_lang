package main

import "fmt"


func main()  {
	// Creating Slices in GO

	// var sli1 [] int
	// sli2 := [] int {}

	// fmt.Println(sli1 == nil)
	// fmt.Println(sli2 == nil)

	// 1. From array
	// 2. Using make()
	// 3. Slice literal


	// 1.
	// arr := [5]int {1, 4, 7, 2, 9}
	// s1 := arr[1:3]
	// fmt.Println("The slice is ", s1)
	// fmt.Println("The len of slice is ", len(s1))
	// fmt.Println("The cap of slice is ", cap(s1))

	//  2. 
	// s2 := make([]int, 4, 10)
	// fmt.Println(s2)
	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))

	//  3. 
	// s3 := []int {1, 2, 3, 4,5}
	// fmt.Println(s3)
	// fmt.Println(len(s3))
	// fmt.Println(cap(s3))

	//  appending
	// s4 := make([]int, 3, 5)
	// s4[0] = 1
	// s4[1] = 2
	// s4[2] = 3
	// fmt.Println(s4)
	// s4 = append(s4, 2, 3, 4)
	// fmt.Println(s4)

	// arr1 := [5]int {1, 2, 3, 4, 5}
	// fmt.Println(arr1)
	
	// slice1 := arr1[2:]
	// fmt.Println(slice1)
	// fmt.Println(cap(slice1))
	// slice2 := append(slice1, 20, 30)
	// fmt.Println(slice2)
	// slice1[1] = 00
	// slice2[1] = 99
	// fmt.Println("Slice 1", slice1)
	// fmt.Println("Slice 2", slice2)

	// Copying slice
	// arr := []int {1, 2, 3, 4}
	// arr2 := make([]int, len(arr))

	// copy(arr2, arr) // copy arr to arr2
	// fmt.Println("arr1", arr)
	// fmt.Println("arr2", arr2)

	// deletion of elements from slice
	arr := []int {1,2 ,3, 4,5}
	index := 2
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Println(arr)


}