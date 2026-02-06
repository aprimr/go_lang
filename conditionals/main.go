package main

import (
	"fmt"
)

func main(){
	// if else
	// age:=12
	// if age>18{
	// 	fmt.Println("You are adult")
	// } else {
	// 	fmt.Println("You are not adult")
	// }

	// if.... else if.... else
	// salary := 300
	// if salary>=200 {
	// 	fmt.Println("Boss")
	// } else if salary>100 && salary<200 {
	// 	fmt.Println("Manager")
	// } else {
	// 	fmt.Println("Employee")
	// }

	// We can declare a variable inside if constructor
	// if age:=20; age>18{
	// 	fmt.Println("Adult", age)
	// } else if age < 18{
	// 	fmt.Println("Kid" , age)
	// }


	// SWITCH STATEMENT //
	// number:= 5
	// switch number{
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3: 
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }

	// MULTIPLE CONDITIONS SWITCH //
	// switch time.Now().Weekday(){
	// 	case time.Sunday, time.Saturday: 
	// 		fmt.Println("Weekend")
	// default:
	// 	fmt.Println("WorkDay")
	// }

	typeOf := func(n any){
		switch n.(type){
		case string:
			fmt.Println("String")
		case int:
			fmt.Println("Int")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other")
		}
	}

	typeOf("golang")
	typeOf(12)
	typeOf(true)
}