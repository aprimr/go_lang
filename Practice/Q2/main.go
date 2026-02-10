// Problem 2: The "Lexer Buffer" (Slice & Memory Management)
// Interpreters use a "Lexer" to break strings into tokens. In systems, we often need to "Peek" at data without consuming it or "Advance" through a buffer.

// The Task:

// Define a type Lexer which is a struct containing a Source (string) and a Cursor (int).

// Write a method NextChar() byte.

// It should return the character at the current Cursor and then move the Cursor forward by 1.

// If the cursor is at the end of the string, return 0.

// Write a method Peek() byte.

// It should return the character at the current Cursor without moving the cursor.

// Write a method Rewind(steps int).

// It should move the Cursor back by steps, but ensure it never goes below 0.

// Test Case:

// Go
// l := Lexer{Source: "if (x > 0)", Cursor: 0}
// fmt.Printf("%c\n", l.NextChar()) // 'i'
// fmt.Printf("%c\n", l.Peek())     // 'f' (not moving)
// fmt.Printf("%c\n", l.NextChar()) // 'f' (now it moves)
// l.Rewind(1)
// fmt.Printf("%c\n", l.Peek())

package main

import "fmt"

type Lexer struct{
	Source string
	Cursor int
}

func (lex *Lexer) NextChar()(char byte){
	if lex.Cursor == len(lex.Source) {
		return 0
	} else{
		char = lex.Source[lex.Cursor]
		lex.Cursor++
		return 
	}
}

func (lex *Lexer) Peek() (char byte) {
	char = lex.Source[lex.Cursor]
	return
}

func (lex *Lexer) Rewind(steps int){
	if steps > lex.Cursor {
		fmt.Println("Cannot perform rewind, bytes excedded by ", steps - lex.Cursor)
	} else{
		lex.Cursor = lex.Cursor - steps
	}
} 

func main()  {
	l := Lexer{Source: "if (x > 0)", Cursor: 0}
	fmt.Printf("%c\n", l.NextChar())
	fmt.Printf("%c\n", l.Peek())
	fmt.Printf("%c\n", l.NextChar())
	l.Rewind(10)
	fmt.Printf("%c\n", l.Peek())
}