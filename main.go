package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var curr []int

	var dfs func(start int)

	dfs = func(start int) {
		res = append(res, append([]int{}, curr...))

		if start == len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])

			dfs(i + 1)

			curr = curr[:len(curr)-1]
		}
	}

	dfs(0)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)

	for _, re := range res {
		fmt.Print("[ ")
		for _, r := range re {
			fmt.Print(r)
		}
		fmt.Println(" ]")
	}
}
