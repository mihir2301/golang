package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error is", err)
		return
	}
	defer res.Body.Close()
	data,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("Error is",err)
		return
	}
	fmt.Println("data is ",string(data))
}
