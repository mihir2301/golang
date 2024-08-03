package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		IsAdult bool   `json:"is_adult"`
	}
	person:=Person{Name: "Mahir",Age:22,IsAdult: true}
	jsonData,err:=json.Marshal(person)
	if err!=nil{
		fmt.Println("eRROR OCCURED",err)
		return
	}
	fmt.Println("json data is ",string(jsonData))
	var originaldata Person
	errors:=json.Unmarshal(jsonData,&originaldata)
	if errors!=nil{
		fmt.Println("ERROR OCCURED",errors)
		return
	}
	fmt.Println(originaldata)
}