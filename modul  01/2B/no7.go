package main

import "fmt"

func main() {

	var bunga, pita string
	jumlah := 0
	i := 1

	for {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita += bunga + " - "
		jumlah++
		i++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}
