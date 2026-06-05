package main

import "fmt"

func main() {
	var x, data string
	var n int

	var ditemukan bool
	var posisi int
	var jumlah int

	fmt.Print("Masukkan string x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Data ke-", i, ": ")
		fmt.Scan(&data)

		if data == x {
			jumlah++

			if !ditemukan {
				ditemukan = true
				posisi = i
			}
		}
	}

	// a
	if ditemukan {
		fmt.Println("a. String ditemukan")
	} else {
		fmt.Println("a. String tidak ditemukan")
	}

	// b
	if ditemukan {
		fmt.Println("b. Posisi pertama ditemukan di indeks", posisi)
	} else {
		fmt.Println("b. String tidak memiliki posisi")
	}

	// c
	fmt.Println("c. Jumlah string x =", jumlah)

	// d
	if jumlah >= 2 {
		fmt.Println("d. Ada sedikitnya dua string x")
	} else {
		fmt.Println("d. Tidak ada sedikitnya dua string x")
	}
}
