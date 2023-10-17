package main

import "fmt"

func countBits(num int, result []uint32) {
	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}
}

func main() {
	var n int = 5
	var result []uint32 = make([]uint32, n+1)
	countBits(n, result)

	fmt.Println(result)
}
