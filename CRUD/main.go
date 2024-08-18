package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// func PerformGetMethod() {
// 	res, err := http.Get("http://jsonplaceholder.typicode.com/todo/1")
// 	if err != nil {
// 		fmt.Println("Error getting Get response", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		fmt.Printf("Error: %d - %s\n", res.StatusCode, http.StatusText(res.StatusCode))
// 		return
// 	}

// 	data,err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body", err)
// 		return
// 	}
// 	fmt.Println("Data received!", string(data))

// 	var todo Todo
// 	err = json.NewDecoder(res.body).Decode(&todo)
// 	if err != nil {
// 		fmt.Println("Error decoding JSON", err)
// 		return
// 	}
// 	fmt.Println(todo)
// }

func PerformPostMethod() {
	todo := Todo{
		UserID:    1,
		Title:     "Test Todo",
		Completed: false,
	}

	// Convert the Struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// convert the josn data to string
	jsonString := string(jsonData)

	//JSON to Reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "http://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Status Code:", res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response:", string(body))

}

func PerformUpdateRequest() {
	todo := Todo{
		UserID:    1,
		Title:     "Test Todo",
		Completed: false,
	}

	// Convert the Struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	myURL := "http://jsonplaceholder.typicode.com/todos/1"

	//Make a reuqest
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	//Request by client
	client := http.Client{}
	response, err := client.Do(req)
	if err!= nil {
        fmt.Println("Error:", err)
        return
    }
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)

}

func PerformDeleteMethod() {
	myURL := "http://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err!= nil {
        fmt.Println("Error:", err)
        return
    }
	client := http.Client{}
	res,err := client.Do(req)
	if err!= nil {
        fmt.Println("Error:", err)
        return
    }
	defer res.Body.Close()
	fmt.Println("Status Code:", res.Status)
}

func main() {
	fmt.Println("CRUD!")
	// PerformGetMethod()
	// PerformPostMethod()
	PerformDeleteMethod()

}
