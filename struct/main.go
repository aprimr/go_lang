package main

import (
	"fmt"
	"time"
)

// type Address struct{
// 	City string
// 	Country string
// }

// type User struct{
// 	Name		 string
// 	Email	 			string
// 	IsVerified bool
// 	Age int
// 	Address Address
// }

// func main()  {
// 	add := Address{"Ghorahu","Nepal"}
// 	user1 := User{"Aprim", "aprim.dev@gmail.com", true, 20,add }
// 	fmt.Println("The user data is ", user1)

// 	fmt.Printf("The user details is %+v \n", user1)
// 	fmt.Printf("The user name is %v \n", user1.Name)

// }

type Order struct {
	id        int
	name      string
	price     float32
	createdAt time.Time
}

// Constructor
func NewOrder(id int, name string, price float32) *Order {
	myOrder := Order{
		id:        id,
		name:      name,
		price:     price,
		createdAt: time.Now(),
	}

	return &myOrder
}

func main() {
	order1 := NewOrder(1, "Chips", 100)
	order2 := NewOrder(2, "Vegetables", 450)

	fmt.Println("Order 1:", order1.name)
	fmt.Println("Order 2:", order2)

	order3 := struct {
		id        int
		name      string
		price     float32
		createdAt time.Time
	}{3, "Hello", 100, time.Now()}

	fmt.Println("Order 3:", order3)
}
