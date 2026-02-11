package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	//// Opening file
	// file, err := os.Open("example.txt")
	// if err != nil{
	// 	fmt.Println("Error opening file", err)
	// 	return
	// }
	// defer file.Close()
	// fmt.Println("File opened successfully")

	//// Reading small file (Whole file at once)
	// data, err := os.ReadFile("example.txt")
	// if err != nil{
	// 	fmt.Println("Error Reading file: ", err)
	// 	return
	// }

	// fmt.Println("The content is: ", string(data))

	//// reading large file (Chunk by Chunk)
	// file, _ := os.Open("example.txt")
	// defer file.Close() 

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// 	fmt.Println("This is a line")
	// }

	//// Writting to a file
	// content := []byte("Hello I am writting to a file")
	// err := os.WriteFile("example.txt", content, 0644)
	// if err != nil {
	// 	fmt.Println("Error writting to the file")
	// }

	//// Appending to a file
	// f, err := os.OpenFile("example.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY , 0644)
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer f.Close()
	
	// _, e := f.WriteString("I am appending to file \n")
	// if e != nil{
	// 	fmt.Println("Error appending to file:", e)
	// }

	// Task: 1. Create a file named input.txt with some lines of text (some containing the word "ERROR"). 2. Write a function FilterErrors(inputPath string, outputPath string). 3. It should read input.txt line by line. 4. If a line contains "ERROR", write that line into outputPath. 5. Use defer to ensure both files are closed properly.
	FilterErrors("input.txt", "errors.txt")

}


func FilterErrors(inputPath string, outputPath string){
	// Open input file for reading
	f,err := os.OpenFile(inputPath, os.O_RDONLY, 0644)
	if err != nil{
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// open file for writting errros
	ef, er := os.OpenFile(outputPath,os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0466)
	if er != nil{
		fmt.Println("Error opening file:", er)
		return
	}
	defer ef.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			ef.WriteString(line)
			ef.WriteString("\n")
		}
	}
	defer ef.Close()

}