package main

import (
	"fmt"
	"math"
)

func main() {
	var A [100]int
	var n, i, x, hapus, cari int
	var jumlah int
	var rata, sd float64

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
		jumlah += A[i]
	}

	fmt.Println("Isi array:")
	for i = 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	fmt.Println("Indeks ganjil:")
	for i = 1; i < n; i += 2 {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	fmt.Println("Indeks genap:")
	for i = 0; i < n; i += 2 {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	fmt.Scan(&x)
	fmt.Println("Kelipatan indeks", x)
	for i = 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(A[i], " ")
		}
	}
	fmt.Println()

	fmt.Scan(&hapus)
	for i = hapus; i < n-1; i++ {
		A[i] = A[i+1]
	}
	n--

	fmt.Println("Setelah dihapus:")
	for i = 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()

	rata = float64(jumlah) / float64(n)

	for i = 0; i < n; i++ {
		sd += math.Pow(float64(A[i])-rata, 2)
	}
	sd = math.Sqrt(sd / float64(n))

	fmt.Println("Rata-rata =", rata)
	fmt.Println("Standar deviasi =", sd)

	fmt.Scan(&cari)
	var frek int
	for i = 0; i < n; i++ {
		if A[i] == cari {
			frek++
		}
	}
	fmt.Println("Frekuensi =", frek)
}
