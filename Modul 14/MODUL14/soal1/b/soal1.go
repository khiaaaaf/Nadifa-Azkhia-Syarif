package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func insertionSort(A *arrInt, n int) {
	var i, j, temp int

	for i = 1; i < n; i++ {
		temp = A[i]
		j = i

		for j > 0 && temp < A[j-1] {
			A[j] = A[j-1]
			j--
		}

		A[j] = temp
	}
}

func main() {
	var A arrInt
	var x, n int
	var tetap bool
	var selisih int

	n = 0

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		A[n] = x
		n++
	}

	insertionSort(&A, n)

	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	tetap = true
	selisih = A[1] - A[0]

	for i := 1; i < n-1; i++ {
		if A[i+1]-A[i] != selisih {
			tetap = false
		}
	}

	if tetap {
		fmt.Println("Data berjarak", selisih)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
