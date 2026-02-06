package main

import (
	"fmt"
)

func main(){

	// While loop using for
	// i:=1
	// for i<=5{
	// 	fmt.Println(i)
	// 	i++
	// }

	// infinite loop
	// for{
	// 	fmt.Println(i)
	// 	i++
	// }

	// classic for loop
	for i:=1 ; i<=5; i++{
		// if i==2{
		// 	break
		// }
		// if i==2{
		// 	continue
		// }
		fmt.Println(i)
	}

	// range
	for i:= range 10{
		fmt.Println(i)
	}

}