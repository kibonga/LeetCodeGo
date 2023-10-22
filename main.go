package main

import "fmt"

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		num &= num - 1
		res += 1
	}
	return res
}

func main() {
	var x uint32 = 5
	res := hammingWeight(x)
	fmt.Println("Number of 1 bits = ", res)
}
