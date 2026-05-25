package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func selectionSort(A *arrInt, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if A[j] < A[idxMin] {
				idxMin = j
			}
		}

		temp = A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}

func main() {
	var daerah, m int
	var A arrInt

	fmt.Scan(&daerah)

	for i := 0; i < daerah; i++ {
		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&A[j])
		}

		selectionSort(&A, m)

		for j := 0; j < m; j++ {
			fmt.Print(A[j], " ")
		}
		fmt.Println()
	}
}
