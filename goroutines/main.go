package main

import (
	"fmt"
	"time"
)

func hello() {
	
	time.Sleep(900* time.Millisecond)
	fmt.Println("hi boss")
}
func hi(){
	fmt.Println("hi")
}
func main() {
	fmt.Println("Mihir Gupta")
	hello()
	go hi()
	time.Sleep(1000* time.Millisecond)
}