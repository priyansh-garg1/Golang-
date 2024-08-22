package main

import "fmt"

func AppendNumber(nums []int) {
	nums = append(nums, 4)
}
func AppendNumberPointer(nums *[]int){
	*nums = append(*nums,4);
}

func main() {
	nums := []int{1, 2, 3}
	AppendNumber(nums)
	fmt.Println(nums)
	AppendNumberPointer(&nums);
	fmt.Println(nums)
}
