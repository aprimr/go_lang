package main

import "fmt"

type Address struct{
	City string
	Country string
}

type User struct{
	Name		 string
	Email	 			string
	IsVerified bool
	Age int
	Address Address
}

func main()  {
	add := Address{"Ghorahu","Nepal"}
	user1 := User{"Aprim", "aprim.dev@gmail.com", true, 20,add }
	fmt.Println("The user data is ", user1)
	
	fmt.Printf("The user details is %+v \n", user1)
	fmt.Printf("The user name is %v \n", user1.Name)

}