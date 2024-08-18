package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "http://jsonplaceholder.typicode.com/todos/1?key1=1"
	// fmt.Printf("Type of myURL %T \n", myURL)

	parseURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error while parsing the URL: ", err)
		return
	}
	// fmt.Printf("Parsed URL type %T \n", parseURL)
	fmt.Println("Scheme: ", parseURL.Scheme)
	fmt.Println("Host: ", parseURL.Host)
	fmt.Println("Path: ", parseURL.Path)
	fmt.Println("Raw Querey: ", parseURL.RawQuery)

	parseURL.Path = "/newPath"
	parseURL.RawQuery = "key2=val2"
	fmt.Println(parseURL.String())
	fmt.Println("Scheme: ", parseURL.Scheme)
	fmt.Println("Host: ", parseURL.Host)
	fmt.Println("Path: ", parseURL.Path)
	fmt.Println("Raw Querey: ", parseURL.RawQuery)

}
