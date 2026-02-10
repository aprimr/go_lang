package main

import "fmt"

// type User struct{
// 	Name string
// 	email string
// 	Status bool
// }

// // method
// func (u User) GetStatus(){ // value reciever
// 	fmt.Println("Active Status: ", u.Status )
// }

// func (u *User) setName(name string){ // pointer reciever
// 	u.Name = name
// 	fmt.Println("Name modified successfully")
// }

// func main()  {
// 	u1 := User{"aprim", "aprim@gmail.com", false}
// 	u1.GetStatus()
// 	u1.setName("user1")
// 	fmt.Println("The original user is ", u1)
// }

// ------ PRACTICE -------

type CPU struct{
	Model string
	Frequency float64
	IsRunning bool
}

func (cpu *CPU) Overclock(multiplier float64){
	cpu.Frequency = cpu.Frequency * multiplier
}

func (cpu *CPU) TogglePower(){
	cpu.IsRunning = !cpu.IsRunning
}

type Buffer []byte

func (buf Buffer) GetSize() int {
	return len(buf)
}

func (buf *Buffer) Clear(){
	*buf = nil
}

func (buf Buffer) EncryptedCopy() []byte {
	newSlice := make([]byte, len(buf))
	 for i, value := range buf{
		newSlice[i] = value+1
	 } 
	 return Buffer(newSlice)
}

func main()  {
	// myCpu := CPU{"Intel i9", 3.5, false}
	// myCpu.TogglePower()
	// myCpu.Overclock(1.2)
	// fmt.Println("My Cpu details is ", myCpu)
	

	b := Buffer{65, 66, 67} 
	fmt.Println(b.GetSize()) 

	encrypted := b.EncryptedCopy()
	fmt.Println(string(encrypted)) 
	fmt.Println(string(b))        

	b.Clear()
	fmt.Println(b)
}