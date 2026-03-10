package main

import "fmt"

func main() {
	var K int
	var f, akar2 float64 = 1, 1

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	f = float64((4*K+2)*(4*K+2)) / float64((4*K+1)*(4*K+3))
	fmt.Printf("Nilai f(K) = %.10f\n", f)

	for k := 0; k <= K; k++ {
		akar2 *= float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
