package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	curr := make([]int, 0, n)
	sort.Ints(nums)
	var dfs func(idx int)

	dfs = func(idx int) {
		ans = append(ans, append([]int{}, curr...))
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			curr = append(curr, nums[i])
			dfs(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0)
	return ans
}

func main() {
	nums := []int{1, 2, 2}
	res := subsets(nums)

	for _, re := range res {
		if len(re) == 0 {
			continue
		}
		fmt.Print("[ ")
		for _, r := range re {
			fmt.Print(r, " ")
		}
		fmt.Println(" ]")
	}
}
