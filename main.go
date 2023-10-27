package main

func permute(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, n)
	curr := make([]int, 0)
	vis := make(map[int]int)
	var dfs func(idx int)

	dfs = func(idx int) {
		if len(curr) == n {
			res = append(res, append([]int{}, curr...))
		}
		for i := 0; i < n; i++ {
			if vis[i] == 0 {
				vis[i]++
				curr = append(curr, nums[i])
				dfs(i + 1)
				curr = curr[:len(curr)-1]
				vis[i]--
			}
		}
	}

	dfs(0)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)

	for _, arr := range res {
		if len(arr) == 0 {
			continue
		}
		print("[ ")
		for _, el := range arr {
			print(el)
		}
		println(" ]")
	}
}
