package main

import "fmt"

func getSum(a, b int) int {

	for b != 0 {
		temp := (a & b) << 1
		a = (a ^ b)
		b = temp
	}

	return a
}

func main() {
	a := 2
	b := 3

	result := getSum(a, b)
	fmt.Println("Result = ", result)
}
