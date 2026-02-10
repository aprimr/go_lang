// Problem 1: The "Instruction Pointer" (State Machine)
// An interpreter works by reading a list of instructions and moving a "Pointer" through memory. You need to manage a virtual CPU state.

// The Task:

// Define a struct VirtualCPU.

// Fields: Registers (a map of string to int), ProgramCounter (int).

// Write a method Execute(instruction string, value int).

// If instruction is "SET", set a register name (key) to the value.

// If instruction is "ADD", add the value to the existing register.

// Every time Execute is called, increment the ProgramCounter by 1.

// Write a method Reset().

// This should clear all registers and set the ProgramCounter back to 0.

// Systems Goal: Decide which methods need Pointer Receivers to ensure the CPU state is actually updated across calls.

// Test Case:

// Go
// cpu := VirtualCPU{Registers: make(map[string]int)}
// cpu.Execute("SET", 10) // Should affect a register like "A"
// cpu.Execute("ADD", 5)
// fmt.Println(cpu.ProgramCounter) // Should be 2
// cpu.Reset()
// fmt.Println(cpu.ProgramCounter) // Should be 0

package main

import "fmt"

type VirtualCPU struct{
	Registers map[string]int
	ProgramCounter int
}

func (vCpu *VirtualCPU) Execute(instruction string, value int){
	// 1. Logic for SET/ADD
	if instruction == "SET"{
		vCpu.Registers["A"] = value 
		vCpu.ProgramCounter++
	}

	if instruction == "ADD"{
		vCpu.Registers["A"] = vCpu.Registers["A"] + value
		vCpu.ProgramCounter++
	}

}

func (vCpu *VirtualCPU) Reset(){
	vCpu.ProgramCounter = 0
	// vCpu.Registers = nil
	vCpu.Registers = make(map[string]int)
}

func main()  {
	cpu := VirtualCPU{Registers: make(map[string]int)}
	cpu.Execute("SET", 10) // Should affect a register like "A"
	cpu.Execute("ADD", 5)
	fmt.Println(cpu) 
	fmt.Println(cpu.ProgramCounter) // Should be 2
	cpu.Reset()
	fmt.Println(cpu.ProgramCounter) // Should be 0
}