package main

import "fmt"

// cetakDeret mencetak deret mulai dari n sampai 1 sesuai aturan:
// jika genap -> n/2, jika ganjil -> 3n+1
func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1) // suku terakhir selalu 1
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
