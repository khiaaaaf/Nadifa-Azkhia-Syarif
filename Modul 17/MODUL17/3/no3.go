package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	var x, y float64
	var A, B, C, D int

	fmt.Print("Jumlah tetesan: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		x = rand.Float64()
		y = rand.Float64()

		if x < 0.5 && y < 0.5 {
			A++
		} else if x >= 0.5 && y < 0.5 {
			B++
		} else if x >= 0.5 && y >= 0.5 {
			C++
		} else {
			D++
		}
	}

	fmt.Println("Curah hujan daerah A:", float64(A)*0.0001, "milimeter")
	fmt.Println("Curah hujan daerah B:", float64(B)*0.0001, "milimeter")
	fmt.Println("Curah hujan daerah C:", float64(C)*0.0001, "milimeter")
	fmt.Println("Curah hujan daerah D:", float64(D)*0.0001, "milimeter")
}
