package main

import "fmt"

type User struct{
	Name string
	Email string
	IsVerified bool
}

func main(){
	user1 := User{"aprim", "aprim@go.dev", false}
	verifyUser(&user1)
	fmt.Println("The status of user is", user1)

	addRes := add(10, 23)
	fmt.Println("The sum is ", addRes)

	a,b := swap(99, 299)
	fmt.Printf("Swapped %d and %d \n", a, b)

	res := mul(3, 9)
	fmt.Println("The multiply res is ", res)

	sumM := sumMultiple(2,3,4,5,6,7,8)
	fmt.Println("The sum of multiple num is ", sumM)
	
	f:= func () string {
		return "I am anonynomous func"
	}
	fmt.Println(f())

	func()  {
		fmt.Println("I am immediately invoked func")
	}()

}

func add(num1, num2 int) int{
	return num1 +num2;
}

func swap(num1, num2 int) (int, int){
	return num2, num1;
}

// named return 
func mul(a,b int) (res int){
	res = a * b
	return
}

func sumMultiple(nums ...int) (res int){
	res = 0
	for _, num := range nums{
		res = res + num
	}
	return
}

// func verifyUser(user User){
// 	user.IsVerified = true
// 	fmt.Println("Copy of user1 inside function ", user)
// }

func verifyUser(user *User){
	user.IsVerified = true
	fmt.Println("Copy of user1 inside function ", user)
}