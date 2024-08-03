package main

import(
	"fmt"
	"bufio"
	"os"
)
func main(){
	fmt.Println("What is your name")
	reader :=bufio.NewReader(os.Stdin)
	name, _ :=reader.ReadString('\n')
	fmt.Println("Hello", name)
}