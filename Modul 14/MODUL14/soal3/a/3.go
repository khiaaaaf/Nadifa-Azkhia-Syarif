package main

import "fmt"

const NMAX = 1000000

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

	n = 0

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			insertionSort(&A, n)

			if n%2 == 1 {
				fmt.Println(A[n/2])
			} else {
				fmt.Println((A[n/2-1] + A[n/2]) / 2)
			}
		} else {
			A[n] = x
			n++
		}
	}
}
