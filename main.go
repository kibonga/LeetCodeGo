package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var result uint32 = 0

	for i := 0; i < 32; i++ {
		n := (num >> i) & 1
		result |= n << (32 - i - 1)
	}

	return result
}

func main() {
	var num uint32 = 43261596

	result := reverseBits(num)
	fmt.Println(result)
}
