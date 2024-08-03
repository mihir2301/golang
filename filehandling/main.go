package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating  a file ", err)
		return
	}
	defer file.Close()
	content := "Hello"
	_, errors := io.WriteString(file, content)
	if errors != nil {
		fmt.Println("Error while writing a file", errors)
		return
	}
	files,errorss:= os.Open("example.txt")
	if errorss !=nil {
		fmt.Println("Error while opening a file", errorss)
	return	}
	for{
	buffer:= make([]byte,1024)
	n,errors:=files.Read(buffer)
	if errors==io.EOF{
		break
	}
	if errors!=nil{
		fmt.Println("error",errors)
		return
	}
	fmt.Println(string(buffer[:n]))
}}

