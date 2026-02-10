package main

import (
	"fmt"
)

func main()  {
	adder := createAdder()
	fmt.Println(adder(10))
	fmt.Println(adder(20))


	debug := makeLogger("DEBUG")
	err := makeLogger("ERROR")

	debug("Memory Allocated!!")
	err("Unexpected args")

	resM := resMonitor(1000)
	fmt.Println(resM(50))
	fmt.Println(resM(500))
	fmt.Println(resM(300))
	fmt.Println(resM(200))
	fmt.Println(resM(200))
	
}

func resMonitor(limit int) func(int) bool{
	totalUsage := 0
	return func(usage int) bool {
		if totalUsage + usage > limit{
			fmt.Println("Limit excedded by ", totalUsage+usage - limit)
			return  true
		} else {
			totalUsage += usage
			fmt.Println("The total usage is", totalUsage)
			return false
		}
	}
}

func createAdder() func(int) int{
	sum := 0
	return func (x int) int {
		sum += x
		return sum
	}
}

func makeLogger(prefix string) func(string){
	return func(message string) {
		fmt.Printf("[%v]: %v \n", prefix, message)
	}
}

