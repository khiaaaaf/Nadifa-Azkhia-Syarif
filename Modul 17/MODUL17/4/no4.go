package main

import (
	"fmt"
)

func main() {
	var jumlah float64
	var suku, sukuNext float64
	var i int = 1

	for {
		if i%2 == 1 {
			suku = 1.0 / float64(2*i-1)
		} else {
			suku = -1.0 / float64(2*i-1)
		}

		if (i+1)%2 == 1 {
			sukuNext = 1.0 / float64(2*(i+1)-1)
		} else {
			sukuNext = -1.0 / float64(2*(i+1)-1)
		}

		jumlah += suku

		selisih := suku - sukuNext
		if selisih < 0 {
			selisih = -selisih
		}

		if selisih <= 0.00001 {
			break
		}

		i++
	}

	fmt.Println("Hasil PI:", 4*jumlah)
	fmt.Println("Pada i ke:", i)
}
