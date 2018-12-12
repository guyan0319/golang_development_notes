package main

import "fmt"

func main() {
	args := []int{1, 2, 3, 4}
	nums := make([]int, 5)
	for _, e := range args {
		nums[e] = e
	}
	fmt.Println(nums)
}
