package main

import "fmt"

func main() {
	var bil, total float64
	var jumlah int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)

	for bil != 9999 {
		total += bil
		jumlah++

		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&bil)
	}

	if jumlah > 0 {
		fmt.Println("Rerata =", total/float64(jumlah))
	} else {
		fmt.Println("Tidak ada data")
	}
}
