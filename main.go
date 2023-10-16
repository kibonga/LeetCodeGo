package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	result := n

	for i := 0; i < n; i++ {
		result ^= i ^ nums[i]
	}

	return result
}

func missingNumber2(nums []int) int {
	n := len(nums)
	result := 0

	for i := 0; i <= n; i++ {
		if i < n {
			result ^= nums[i]
		}
		result ^= i
	}

	return result
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	result := missingNumber2(nums)
	fmt.Println("Missing number = ", result)
}
