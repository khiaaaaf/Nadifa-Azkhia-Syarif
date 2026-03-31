package main

import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d) // input a b c d

	P1 := permutation(a, c)
	C1 := combination(a, c)
	fmt.Println(P1, C1)

	P2 := permutation(b, d)
	C2 := combination(b, d)
	fmt.Println(P2, C2)
}
