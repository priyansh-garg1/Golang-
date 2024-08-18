package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// HTTP get method
	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting Get response",err)
		return;
	}
	defer response.Body.Close()
	// fmt.Printf("Type of response %T\n", response)
	// fmt.Println(response)
	data,err := ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println("Error reading response body", err)
        return;
	}
	fmt.Println("Response :",string(data))
}
