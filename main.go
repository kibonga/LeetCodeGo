package main

import (
	"fmt"
	"math"
)

func reverse(x uint32) uint32 {
	var num uint32 = 0

	for x > 0 {
		if (math.MaxUint32 / 10) < num {
			return 0
		}
		num = (10 * num) + (x % 10)
		x /= 10
	}

	return num
}

func main() {
	var x uint32 = 123
	result := reverse(x)
	fmt.Println("Result = ", result)
}
