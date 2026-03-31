package main

import (
	"fmt"
	"strings"
)

// hitungSkor menghitung jumlah soal yang berhasil diselesaikan dan total skor (waktu)
func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0
	*skor = 0
	for _, t := range waktu {
		if t <= 300 { // jika <= 5 jam (300 menit), soal berhasil
			*soal++
			*skor += t
		}
	}
}

func main() {
	var nama string
	var waktu [8]int

	pemenang := ""
	maxSoal := -1
	minSkor := 0

	for {
		fmt.Scan(&nama)
		if strings.ToLower(nama) == "selesai" {
			break
		}

		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, skor int
		hitungSkor(waktu[:], &soal, &skor)

		// Tentukan pemenang
		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			pemenang = nama
			maxSoal = soal
			minSkor = skor
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}
