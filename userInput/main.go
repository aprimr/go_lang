package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	message:= "Welcome to go lang"
	fmt.Println(message)

	// reader
	reader := bufio.NewReader(os.Stdin)


	// fmt.Println("Enter your name")
	// name, _ := reader.ReadString('\n')
	// fmt.Println("WELCOME! ", name)
	// fmt.Printf("The type of variable is %T \n", name)

	fmt.Println("Enter your age")
	age, _ := reader.ReadString('\n')
	fmt.Println("Your age is ",age)
	fmt.Printf("The type of variable is %T \n", age)

	intAge, err := strconv.ParseFloat(strings.TrimSpace(age), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New age is:", intAge + 5)
	}


}