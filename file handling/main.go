package main

import (
	"fmt"
	"os"
)

func main() {
	// Create and write in file
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file")
	// 	return
	// }
	// defer file.Close()

	// content := "hello world from client"
	// bytereturns,err := io.WriteString(file,content+"\n")
	// fmt.Println("Byte return by file",bytereturns)
	// fmt.Println("File created successfully")

	// Read from file
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while Opening the file")
	// 	return
	// }
	// defer file.Close()
	// buffer := make([]byte, 1024)
	// for {
	// 	n, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err!= nil {
	//         fmt.Println("Error while reading the file")
	//         return
	//     }
	// 	fmt.Println(string(buffer[:n]))
	// }

	//ioutil to read the file
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading the file")
		return
	}
	fmt.Println(string(content))
}
