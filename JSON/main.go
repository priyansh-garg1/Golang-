package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("JSON")
	person := Person{Name: "John", Age: 34, IsAdult: true}
	fmt.Println(person)

	//Convert Person to JSON
	josnData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("JSON DATA:", string(josnData))
	var p1 Person
	err = json.Unmarshal(josnData, &p1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Person from JSON:", p1)
}
