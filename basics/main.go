package main

import (
	"fmt"
)

func init() {
	fmt.Println("Hello world Init")
}

// type Person struct {
// 	name string
// 	age  int
// }

// type Contact struct {
// 	email   string
// 	phone   int
// 	address string
// }

// type Employee struct {
// 	person Person
// 	cont   Contact
// 	salary float64
// }

func main() {
	// fmt.Println("I am Learning Go Languages")

	// // Variables
	// var name string  = "Hello"
	// var version = 1.001
	// version = 2.001
	// var dec float32 = 0.001
	// fmt.Println(dec);/

	// age := 12
	// name := "User"
	// person := 123
	// var agee int = 12
	// fmt.Println("person: ", person, "name: ", name, "age", age)
	// fmt.Printf("age is %d", agee)
	// fmt.Printf("name is %T", name)

	// Input User
	// fmt.Println("Enter your name")
	// var name string
	// Only Take first word
	// fmt.Scan(&name)
	// fmt.Printf("Hello %s", name)
	// // to get full line
	// reader := bufio.NewReader(os.Stdin)
	// name, _ = reader.ReadString('\n')
	// fmt.Println(name)

	// simpleFunction()
	// ans, err := div(5, 0)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// ans, _ := div(5, 2)
	// fmt.Println("Division:", ans)

	// Array
	// var name [5]string
	// name[0] = "Hello world"
	// fmt.Println("Index 0 value:", name[0])
	// fmt.Println(len(name[0]))
	// fmt.Printf("Name is %q",name)

	// Slice
	// number := []int{0,1,2,3,4,5}
	// number = append(number, 2, 3, 4, 5)
	// fmt.Println("Slice is", number)

	// If-Else
	// num := 10
	// if num % 2 == 0 {
	//     fmt.Println("Number is even")
	// } else if num % 2 == 1 {
	//     fmt.Println("Number is odd")
	// } else {
	// 	fmt.Println("Number is neither even nor odd")
	// }

	// Switch
	// day := 3
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// default:
	// 	fmt.Println("Unkown number")
	// }

	// month := "January"
	// switch month {
	// case "January", "March", "May", "July", "August", "October", "December":
	// 	fmt.Println("31 days")
	// case "February":
	// 	fmt.Println("28 or 29 days")
	// default:
	// 	fmt.Println("Invalid month")
	// }

	// For Loop
	// for i:= 0; i<10;i++ {
	// 	fmt.Println("index is :",i)
	// }
	// Like While Loop
	// counter := 0
	// for{
	// 	counter++
	//     if counter == 10 {
	//         break
	//     }
	//     fmt.Println("index is :", counter)
	// }

	// number := []int{1, 2, 3, 4, 5}
	// for index, value := range number {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }
	// str := "Hello, world!"
	// for idx,val := range str {
	// 	fmt.Printf("Index: %d, Char Value: %c\n", idx, val)
	// }

	// Maps
	// students := make(map[string]int)
	// students["John"] = 20
	// students["Jane"] = 22
	// fmt.Println("John's age is:", students["John"])
	// delete(students, "John")
	// age,isExits := students["John"]
	// fmt.Println("Is John exists:", isExits, "John's age is:", age)
	// fmt.Println("John's age is:", students["John"])

	// for key,val := range students {
	// 	fmt.Printf("Key: %s, Value: %d\n", key, val)
	// }
	// person := map[string]int{
	// 	"name"   :10,
	//     "age"     : 30,
	//     "address" :20,
	// }
	// fmt.Println("Person's age is:", person["age"])

	// Struct
	// var p1 Person
	// p1.name = "John"
	// p1.age = 30
	// p2 := Person{name: "John", age: 27}
	// var p3 = new(Person)
	// p3.age = 29
	// fmt.Println("P1 is :", p1, p2, p3)

	// var emp Employee
	// emp.person = p2
	// emp.cont = Contact{
	// 	email:   "john@example.com",
	//     phone:   1234567890,
	//     address: "123 Main St",
	// }
	// fmt.Println(emp)

	// Pointers
	// num := 2
	// ptr := &num
	// fmt.Println("Value of num is:", &num,"Value of ptr is:", *ptr)

	// Type Conversion
	// num := 42
	// fmt.Println("Number is: ", num)
	// fmt.Printf("Number is: %T \n", num)
	// floatNum := float32(num)
	// fmt.Println("Float number is: ", floatNum)
	// fmt.Printf("Float Type is: %T \n", floatNum)
	// num := 10
	// str := strconv.Itoa(num)
	// fmt.Printf("Type is: %T \n", str)

	// for i := 1; i < 10; i++ {
	// 	fmt.Println("Hello", i)
	// }

	// i := 1
	// for i<10 {
	// 	fmt.Println("hello",i);
	// 	i++
	// }

	// arr := []string{"hello", "world", "useer"};

	// for i,val := range arr {
	// 	fmt.Printf("Index: %d, Value: %s\n", i, val)
	// }

	// for i,val := range "hello" {
	// 	fmt.Printf("Index: %d, Value: %c\n", i, val)
	// }

	// jsonmap := map[int]string {
	// 	1: "John",
	//     2: "Jane",
	// }
	// for key, val := range jsonmap {
	//     fmt.Printf("Key: %d, Value: %s\n", key, val)
	// }

	// variadicFunc(1,2,3,4,6,5)

	// func(){
	// 	fmt.Println("Hello World!")
	// }()

}

func init() {
	fmt.Println("Init 2")
}

// func variadicFunc(val ...int){
//         fmt.Println("VALUES are ",val);
// }

// Function
// func simpleFunction() {
// 	fmt.Print("Function")
// 	fmt.Println("val is:", add(10, 20))
// }

// func add(a, b int) int {
// 	fmt.Println(a + b)
// 	return a + b
// }

// func div(a, b float32) (float32, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }
