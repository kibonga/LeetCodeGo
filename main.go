package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	curr := make([]int, 0)

	var dfs func(idx int, currSum int, curr []int)

	dfs = func(idx int, currSum int, curr []int) {
		if currSum == target {
			result = append(result, append([]int{}, curr...))
			return
		}
		if currSum > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			curr = append(curr, candidates[i])
			dfs(i, currSum+candidates[i], curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, 0, curr)
	return result
}

func main() {
	nums := []int{2, 3, 6, 7}
	target := 7

	result := combinationSum(nums, target)

	for _, res := range result {
		fmt.Print("[")
		for _, r := range res {
			fmt.Print(r, " ")
		}
		fmt.Println("]")
	}
}
