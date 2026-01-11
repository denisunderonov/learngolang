package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	nums = append(nums, 40)

	for _, num := range nums {
		fmt.Println(num)
	}
}
