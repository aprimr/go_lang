package main

import "fmt"


func main(){
	fmt.Println("Hello World")

	// var name string ="golang"
	// var phone int = 92121212

	// type infer
	// var name = "golang"
	// var isAdult = true
	// var age = 20

	// fmt.Println(name)
	// fmt.Println(phone)
	// fmt.Println(isAdult)
	// fmt.Println(age)


	// short hand syntax
	// name:= "golang"
	// age:= 22
	// isAdult:= true

	// constants
	// const name = "golang"
	// const age = 14
	// const isAdult = false

	// constant grouping
	const(
		name = "golang"
		age= 43
		isAdult = true
	)
	
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isAdult)

}