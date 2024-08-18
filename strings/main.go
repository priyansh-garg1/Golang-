package main

import (
	"fmt"
	"strings"
)

func main() {
	//String
	// data := "Apple,mango,mango"
	// parts := strings.Split(data, ",")
	// fmt.Println(parts)
	// data2 := "one two three four two two"
	// count := strings.Count(data2, "two")
	// fmt.Println(count)
	// data3 := "       apple        apple     "
	// trimmed := strings.TrimSpace(data3)
	// fmt.Println(trimmed)
	// str1 := "hello"
	// str2 := "world"
	// res := strings.Join([]string{str1, str2}, " ")
	// fmt.Println(res)

	//Time & Date
	// curTime := time.Now()
	// fmt.Println("Current Time:", curTime)
	// formatted := curTime.Format("02/01/2006 03:04:05")
	// fmt.Println(formatted)

	// dateStr := "2024-09-25"
	// format := "2006-01-02"
	// formattedStr, _ := time.Parse(format, dateStr)
	// fmt.Println(formattedStr)

	//Defer in Golang
	// fmt.Println("Starting")
	// defer fmt.Println("middle")
	// defer fmt.Println("middle")
	// fmt.Println("end")

	// fmt.Println("Starting")
	// data := add(3, 5)
	// fmt.Println("Data:", data)
	// defer fmt.Println("Defer Data:", data)
	// fmt.Println("middle")
	// fmt.Println("emmmd")

	// fmt.Println(strings.Trim("adbvaa", "a"))

	// fmt.Println(strings.Split("abvcxxxzssss","v")[1])
	fmt.Println(strings.ToUpper("Quantity"))

}

// func add(a, b int) int {
// 	return (a + b)
// }
